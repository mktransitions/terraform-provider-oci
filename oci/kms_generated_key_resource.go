// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/hashicorp/terraform/helper/hashcode"
	oci_kms "github.com/oracle/oci-go-sdk/v25/keymanagement"
)

func init() {
	RegisterResource("oci_kms_generated_key", KmsGeneratedKeyResource())
}

func KmsGeneratedKeyResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createKmsGeneratedKey,
		Read:     readKmsGeneratedKey,
		Delete:   deleteKmsGeneratedKey,
		Schema: map[string]*schema.Schema{
			// Required
			"crypto_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"include_plaintext_key": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"key_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key_shape": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"algorithm": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"length": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Optional
			"associated_data": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"logging_context": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"ciphertext": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plaintext": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plaintext_checksum": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createKmsGeneratedKey(d *schema.ResourceData, m interface{}) error {
	sync := &KmsGeneratedKeyResourceCrud{}
	sync.D = d
	endpoint, ok := d.GetOkExists("crypto_endpoint")
	if !ok {
		return fmt.Errorf("crypto_endpoint missing")
	}
	client, err := m.(*OracleClients).KmsCryptoClient(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return CreateResource(d, sync)
}

func readKmsGeneratedKey(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteKmsGeneratedKey(d *schema.ResourceData, m interface{}) error {
	return nil
}

type KmsGeneratedKeyResourceCrud struct {
	BaseCrud
	Client                 *oci_kms.KmsCryptoClient
	Res                    *oci_kms.GeneratedKey
	DisableNotFoundRetries bool
}

func (s *KmsGeneratedKeyResourceCrud) ID() string {
	return string(hashcode.String(*s.Res.Ciphertext))
}

func (s *KmsGeneratedKeyResourceCrud) Create() error {
	request := oci_kms.GenerateDataEncryptionKeyRequest{}

	if associatedData, ok := s.D.GetOkExists("associated_data"); ok {
		request.AssociatedData = objectMapToStringMap(associatedData.(map[string]interface{}))
	}

	if includePlaintextKey, ok := s.D.GetOkExists("include_plaintext_key"); ok {
		tmp := includePlaintextKey.(bool)
		request.IncludePlaintextKey = &tmp
	}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	if keyShape, ok := s.D.GetOkExists("key_shape"); ok {
		if tmpList := keyShape.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "key_shape", 0)
			tmp, err := s.mapToKeyShape(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.KeyShape = &tmp
		}
	}

	if loggingContext, ok := s.D.GetOkExists("logging_context"); ok {
		request.LoggingContext = objectMapToStringMap(loggingContext.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.GenerateDataEncryptionKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.GeneratedKey
	return nil
}

func (s *KmsGeneratedKeyResourceCrud) SetData() error {
	if s.Res.Ciphertext != nil {
		s.D.Set("ciphertext", *s.Res.Ciphertext)
	}

	if s.Res.Plaintext != nil {
		s.D.Set("plaintext", *s.Res.Plaintext)
	}

	if s.Res.PlaintextChecksum != nil {
		s.D.Set("plaintext_checksum", *s.Res.PlaintextChecksum)
	}

	return nil
}

func (s *KmsGeneratedKeyResourceCrud) mapToKeyShape(fieldKeyFormat string) (oci_kms.KeyShape, error) {
	result := oci_kms.KeyShape{}

	if algorithm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "algorithm")); ok {
		result.Algorithm = oci_kms.KeyShapeAlgorithmEnum(algorithm.(string))
	}

	if length, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "length")); ok {
		tmp := length.(int)
		result.Length = &tmp
	}

	return result, nil
}
