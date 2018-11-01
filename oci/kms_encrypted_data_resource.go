// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"fmt"

	"github.com/hashicorp/terraform/helper/hashcode"
	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
)

func EncryptedDataResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createEncryptedData,
		Read:     readEncryptedData,
		Delete:   deleteEncryptedData,
		Schema: map[string]*schema.Schema{
			// Required
			"crypto_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"plaintext": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"associated_data": {
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
		},
	}
}

func createEncryptedData(d *schema.ResourceData, m interface{}) error {
	sync := &EncryptedDataResourceCrud{}
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

func readEncryptedData(d *schema.ResourceData, m interface{}) error {
	sync := &EncryptedDataResourceCrud{}
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

	return ReadResource(sync)
}

func deleteEncryptedData(d *schema.ResourceData, m interface{}) error {
	return nil
}

type EncryptedDataResourceCrud struct {
	BaseCrud
	Client                 *oci_kms.KmsCryptoClient
	Res                    *oci_kms.EncryptedData
	DisableNotFoundRetries bool
}

func (s *EncryptedDataResourceCrud) ID() string {
	return string(hashcode.String(*s.Res.Ciphertext))
}

func (s *EncryptedDataResourceCrud) Create() error {
	request := oci_kms.EncryptRequest{}

	if associatedData, ok := s.D.GetOkExists("associated_data"); ok {
		request.AssociatedData = objectMapToStringMap(associatedData.(map[string]interface{}))
	}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	if plaintext, ok := s.D.GetOkExists("plaintext"); ok {
		tmp := plaintext.(string)
		request.Plaintext = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.Encrypt(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EncryptedData
	return nil
}

func (s *EncryptedDataResourceCrud) Get() error {

	if cipherText, ok := s.D.GetOkExists("ciphertext"); ok {
		tmp := cipherText.(string)
		encryptedData := oci_kms.EncryptedData{Ciphertext: &tmp}
		s.Res = &encryptedData
	} else {
		return s.Create()
	}

	return nil
}

func (s *EncryptedDataResourceCrud) SetData() error {
	if s.Res.Ciphertext != nil {
		s.D.Set("ciphertext", *s.Res.Ciphertext)
	}

	return nil
}
