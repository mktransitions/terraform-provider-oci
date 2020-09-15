// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v25/common"
	oci_kms "github.com/oracle/oci-go-sdk/v25/keymanagement"
)

func init() {
	RegisterResource("oci_kms_vault", KmsVaultResource())
}

func KmsVaultResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createKmsVault,
		Read:     readKmsVault,
		Update:   updateKmsVault,
		Delete:   deleteKmsVault,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vault_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_of_deletion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restore_from_object_store": {
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				MinItems:      1,
				ConflictsWith: []string{"restore_from_file"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"destination": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"BUCKET",
								"PRE_AUTHENTICATED_REQUEST_URI",
							}, true),
						},

						// Optional
						"bucket": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"object": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},
			"restore_from_file": {
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				MinItems:      1,
				ConflictsWith: []string{"restore_from_object_store"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"restore_vault_from_file_details": {
							Type:     schema.TypeString,
							Required: true,
						},
						"content_length": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateFunc:     validateInt64TypeString,
							DiffSuppressFunc: int64StringDiffSuppressFunction,
						},

						// Optional
						"content_md5": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},
			"restore_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			// Computed
			"crypto_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"restored_from_vault_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createKmsVault(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).kmsVaultClient()

	return CreateResource(d, sync)
}

func readKmsVault(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).kmsVaultClient()

	return ReadResource(sync)
}

func updateKmsVault(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).kmsVaultClient()

	return UpdateResource(d, sync)
}

func deleteKmsVault(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).kmsVaultClient()

	return DeleteResource(d, sync)
}

type KmsVaultResourceCrud struct {
	BaseCrud
	Client                 *oci_kms.KmsVaultClient
	Res                    *oci_kms.Vault
	DisableNotFoundRetries bool
}

func (s *KmsVaultResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *KmsVaultResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_kms.VaultLifecycleStateCreating),
		string(oci_kms.VaultLifecycleStateRestoring),
	}
}

func (s *KmsVaultResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_kms.VaultLifecycleStateActive),
	}
}

func (s *KmsVaultResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_kms.VaultLifecycleStateDeleting),
		string(oci_kms.VaultLifecycleStateSchedulingDeletion),
	}
}

func (s *KmsVaultResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_kms.VaultLifecycleStateDeleted),
		string(oci_kms.VaultLifecycleStatePendingDeletion),
	}
}

func (s *KmsVaultResourceCrud) Create() error {
	if _, ok := s.D.GetOkExists("restore_from_file"); ok {
		err := s.RestoreVaultFromFile()
		if err != nil {
			return err
		}
		s.D.SetId(s.ID())
		return s.UpdateVaultDetails()
	}
	if _, ok := s.D.GetOkExists("restore_from_object_store"); ok {
		err := s.RestoreVaultFromObjectStore()
		if err != nil {
			return err
		}
		s.D.SetId(s.ID())
		return s.UpdateVaultDetails()
	}

	request := oci_kms.CreateVaultRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if vaultType, ok := s.D.GetOkExists("vault_type"); ok {
		request.VaultType = oci_kms.CreateVaultDetailsVaultTypeEnum(vaultType.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.CreateVault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vault
	return nil
}

func (s *KmsVaultResourceCrud) Get() error {
	request := oci_kms.GetVaultRequest{}

	tmp := s.D.Id()
	request.VaultId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.GetVault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vault
	return nil
}

func (s *KmsVaultResourceCrud) Update() error {
	if _, ok := s.D.GetOk("restore_from_file"); ok && s.D.HasChange("restore_trigger") {
		err := s.RestoreVaultFromFile()
		if err != nil {
			return err
		}
		s.D.SetId(s.ID())
	}
	if _, ok := s.D.GetOk("restore_from_object_store"); ok && s.D.HasChange("restore_trigger") {
		err := s.RestoreVaultFromObjectStore()
		if err != nil {
			return err
		}
		s.D.SetId(s.ID())
	}
	return s.UpdateVaultDetails()
}

func (s *KmsVaultResourceCrud) UpdateVaultDetails() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_kms.UpdateVaultRequest{}
	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}
	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}
	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}
	tmp := s.D.Id()
	request.VaultId = &tmp
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.UpdateVault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vault
	return nil
}

func (s *KmsVaultResourceCrud) Delete() error {
	request := oci_kms.ScheduleVaultDeletionRequest{}

	tmp := s.D.Id()
	request.VaultId = &tmp

	if timeOfDeletion, ok := s.D.GetOkExists("time_of_deletion"); ok {
		tmpTime, err := time.Parse(time.RFC3339Nano, timeOfDeletion.(string))
		if err != nil {
			return err
		}
		request.TimeOfDeletion = &oci_common.SDKTime{Time: tmpTime}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	_, err := s.Client.ScheduleVaultDeletion(context.Background(), request)
	return err
}

func (s *KmsVaultResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CryptoEndpoint != nil {
		s.D.Set("crypto_endpoint", *s.Res.CryptoEndpoint)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ManagementEndpoint != nil {
		s.D.Set("management_endpoint", *s.Res.ManagementEndpoint)
	}

	if s.Res.RestoredFromVaultId != nil {
		s.D.Set("restored_from_vault_id", *s.Res.RestoredFromVaultId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	s.D.Set("vault_type", s.Res.VaultType)

	return nil
}

func (s *KmsVaultResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_kms.ChangeVaultCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VaultId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	_, err := s.Client.ChangeVaultCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *KmsVaultResourceCrud) RestoreVaultFromObjectStore() error {
	request := oci_kms.RestoreVaultFromObjectStoreRequest{}

	if backupLocation, ok := s.D.GetOkExists("restore_from_object_store"); ok {
		if tmpList := backupLocation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "restore_from_object_store", 0)
			tmp, err := s.mapToBackupLocation(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BackupLocation = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.RestoreVaultFromObjectStore(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vault
	return nil
}

func (s *KmsVaultResourceCrud) RestoreVaultFromFile() error {
	request := oci_kms.RestoreVaultFromFileRequest{}

	if compartmentId, ok := s.D.GetOk("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if restoreVaultFromFileDetails, ok := s.D.GetOk("restore_from_file.0.restore_vault_from_file_details"); ok {
		decodedFileContent, _ := base64.StdEncoding.DecodeString(restoreVaultFromFileDetails.(string))
		request.RestoreVaultFromFileDetails = ioutil.NopCloser(bytes.NewBuffer(decodedFileContent))
	} else {
		request.RestoreVaultFromFileDetails = ioutil.NopCloser(bytes.NewBuffer([]byte{}))
	}

	if contentLength, ok := s.D.GetOk("restore_from_file.0.content_length"); ok {
		tmp := contentLength.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert content-length string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ContentLength = &tmpInt64
	}

	if contentMd5, ok := s.D.GetOk("restore_from_file.0.content_md5"); ok {
		tmp := contentMd5.(string)
		request.ContentMd5 = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.RestoreVaultFromFile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vault
	return nil
}

func (s *KmsVaultResourceCrud) mapToBackupLocation(fieldKeyFormat string) (oci_kms.BackupLocation, error) {
	var baseObject oci_kms.BackupLocation
	//discriminator
	destinationRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination"))
	var destination string
	if ok {
		destination = destinationRaw.(string)
	} else {
		destination = "" // default value
	}
	switch strings.ToLower(destination) {
	case strings.ToLower("BUCKET"):
		details := oci_kms.BackupLocationBucket{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
			tmp := object.(string)
			details.ObjectName = &tmp
		}
		baseObject = details
	case strings.ToLower("PRE_AUTHENTICATED_REQUEST_URI"):
		details := oci_kms.BackupLocationUri{}
		if uri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
			tmp := uri.(string)
			details.Uri = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown destination '%v' was specified", destination)
	}
	return baseObject, nil
}

func VaultBackupLocationToMap(obj *oci_kms.BackupLocation) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_kms.BackupLocationBucket:
		result["destination"] = "BUCKET"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.ObjectName != nil {
			result["object"] = string(*v.ObjectName)
		}
	case oci_kms.BackupLocationUri:
		result["destination"] = "PRE_AUTHENTICATED_REQUEST_URI"

		if v.Uri != nil {
			result["uri"] = string(*v.Uri)
		}
	default:
		log.Printf("[WARN] Received 'destination' of unknown type %v", *obj)
		return nil
	}

	return result
}
