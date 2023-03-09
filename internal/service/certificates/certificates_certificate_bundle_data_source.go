// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_certificates "github.com/oracle/oci-go-sdk/v65/certificates"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CertificatesCertificateBundleDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCertificatesCertificateBundle,
		Schema: map[string]*schema.Schema{
			"cert_chain_pem": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_bundle_type": {
				Type:             schema.TypeString,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"CERTIFICATE_CONTENT_PUBLIC_ONLY",
					"CERTIFICATE_CONTENT_WITH_PRIVATE_KEY",
				}, true),
				Optional: true,
				Computed: true,
			},
			"certificate_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"certificate_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_pem": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_version_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"private_key_pem": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_key_pem_passphrase": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"revocation_status": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time_revoked": {
							Type:     schema.TypeString,
							Required: true,
						},
						"revocation_reason": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"stage": {
				Type:             schema.TypeString,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"CURRENT",
					"PENDING",
					"LATEST",
					"PREVIOUS",
					"DEPRECATED",
				}, true),
				Optional: true,
			},
			"stages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"validity": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time_of_validity_not_before": {
							Type:     schema.TypeString,
							Required: true,
						},
						"time_of_validity_not_after": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"version_number": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func readCertificatesCertificateBundle(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesCertificateBundleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesClient()

	return tfresource.ReadResource(sync)
}

type CertificatesCertificateBundleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_certificates.CertificatesClient
	Res    *oci_certificates.GetCertificateBundleResponse
}

func (s *CertificatesCertificateBundleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CertificatesCertificateBundleDataSourceCrud) Get() error {
	request := oci_certificates.GetCertificateBundleRequest{}

	if certificateId, ok := s.D.GetOkExists("certificate_id"); ok {
		tmp := certificateId.(string)
		request.CertificateId = &tmp
	}

	if versionNumber, ok := s.D.GetOkExists("version_number"); ok {
		tmp := versionNumber.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert versionNumber string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.VersionNumber = &tmpInt64
	}

	if certificateVersionName, ok := s.D.GetOkExists("certificate_version_name"); ok {
		tmp := certificateVersionName.(string)
		request.CertificateVersionName = &tmp
	}

	if stage, ok := s.D.GetOkExists("stage"); ok {
		request.Stage = oci_certificates.GetCertificateBundleStageEnum(stage.(string))
	}

	if certificateBundleType, ok := s.D.GetOkExists("certificate_bundle_type"); ok {
		request.CertificateBundleType = oci_certificates.GetCertificateBundleCertificateBundleTypeEnum(certificateBundleType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "certificates")

	response, err := s.Client.GetCertificateBundle(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CertificatesCertificateBundleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetCertificateId())

	if s.Res.GetCertificateName() != nil {
		s.D.Set("certificate_name", *s.Res.GetCertificateName())
	}

	if s.Res.GetVersionNumber() != nil {
		s.D.Set("version_number", strconv.FormatInt(*s.Res.GetVersionNumber(), 10))
	}

	if s.Res.GetSerialNumber() != nil {
		s.D.Set("serial_number", *s.Res.GetSerialNumber())
	}

	if s.Res.GetTimeCreated() != nil {
		s.D.Set("time_created", s.Res.GetTimeCreated().String())
	}

	if s.Res.GetValidity() != nil {
		s.D.Set("validity", []interface{}{validityToMap(s.Res.GetValidity())})
	} else {
		s.D.Set("validity", nil)
	}

	stages := []interface{}{}
	for _, item := range s.Res.GetStages() {
		stages = append(stages, item)
	}
	s.D.Set("stages", stages)

	if s.Res.GetCertificatePem() != nil {
		s.D.Set("certificate_pem", *s.Res.GetCertificatePem())
	}

	if s.Res.GetCertChainPem() != nil {
		s.D.Set("cert_chain_pem", *s.Res.GetCertChainPem())
	}

	if s.Res.GetVersionName() != nil {
		s.D.Set("version_name", *s.Res.GetVersionName())
	}

	if s.Res.GetRevocationStatus() != nil {
		s.D.Set("revocation_status", []interface{}{revocationStatusToMap(s.Res.GetRevocationStatus())})
	} else {
		s.D.Set("revocation_status", nil)
	}

	if bundle, ok := s.Res.CertificateBundle.(oci_certificates.CertificateBundleWithPrivateKey); ok {
		if bundle.PrivateKeyPem != nil {
			s.D.Set("private_key_pem", *bundle.PrivateKeyPem)
		}
		if bundle.PrivateKeyPemPassphrase != nil {
			s.D.Set("private_key_passphrase", *bundle.PrivateKeyPemPassphrase)
		}
		s.D.Set("certificate_bundle_type", "CERTIFICATE_CONTENT_WITH_PRIVATE_KEY")
	} else {
		s.D.Set("certificate_bundle_type", "CERTIFICATE_CONTENT_PUBLIC_ONLY")
	}

	return nil
}

func validityToMap(obj *oci_certificates.Validity) map[string]interface{} {
	result := map[string]interface{}{
		"time_of_validity_not_before": obj.TimeOfValidityNotBefore.String(),
		"time_of_validity_not_after":  obj.TimeOfValidityNotAfter.String(),
	}

	return result
}

func revocationStatusToMap(obj *oci_certificates.RevocationStatus) map[string]interface{} {
	result := map[string]interface{}{
		"time_revoked":      obj.TimeRevoked.String(),
		"revocation_reason": obj.RevocationReason,
	}

	return result
}
