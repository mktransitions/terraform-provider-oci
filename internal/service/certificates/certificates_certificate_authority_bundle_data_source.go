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

func CertificatesCertificateAuthorityBundleDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCertificatesCertificateAuthorityBundle,
		Schema: map[string]*schema.Schema{
			"cert_chain_pem": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_authority_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"certificate_authority_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_pem": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_authority_version_name": {
				Type:     schema.TypeString,
				Optional: true,
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
			"version_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version_number": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func readCertificatesCertificateAuthorityBundle(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesCertificateAuthorityBundleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesClient()

	return tfresource.ReadResource(sync)
}

type CertificatesCertificateAuthorityBundleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_certificates.CertificatesClient
	Res    *oci_certificates.GetCertificateAuthorityBundleResponse
}

func (s *CertificatesCertificateAuthorityBundleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CertificatesCertificateAuthorityBundleDataSourceCrud) Get() error {
	request := oci_certificates.GetCertificateAuthorityBundleRequest{}

	if certificateAuthorityId, ok := s.D.GetOkExists("certificate_authority_id"); ok {
		tmp := certificateAuthorityId.(string)
		request.CertificateAuthorityId = &tmp
	}

	if versionNumber, ok := s.D.GetOkExists("version_number"); ok {
		tmp := versionNumber.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert versionNumber string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.VersionNumber = &tmpInt64
	}

	if certificateAuthorityVersionName, ok := s.D.GetOkExists("certificate_authority_version_name"); ok {
		tmp := certificateAuthorityVersionName.(string)
		request.CertificateAuthorityVersionName = &tmp
	}

	if stage, ok := s.D.GetOkExists("stage"); ok {
		request.Stage = oci_certificates.GetCertificateAuthorityBundleStageEnum(stage.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "certificates")

	response, err := s.Client.GetCertificateAuthorityBundle(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CertificatesCertificateAuthorityBundleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.CertificateAuthorityId)

	if s.Res.CertChainPem != nil {
		s.D.Set("cert_chain_pem", *s.Res.CertChainPem)
	}

	if s.Res.CertificateAuthorityName != nil {
		s.D.Set("certificate_authority_name", *s.Res.CertificateAuthorityName)
	}

	if s.Res.CertificatePem != nil {
		s.D.Set("certificate_pem", *s.Res.CertificatePem)
	}

	if s.Res.RevocationStatus != nil {
		s.D.Set("revocation_status", []interface{}{revocationStatusToMap(s.Res.RevocationStatus)})
	} else {
		s.D.Set("revocation_status", nil)
	}

	if s.Res.SerialNumber != nil {
		s.D.Set("serial_number", *s.Res.SerialNumber)
	}

	stages := []interface{}{}
	for _, item := range s.Res.Stages {
		stages = append(stages, item)
	}
	s.D.Set("stages", stages)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Validity != nil {
		s.D.Set("validity", []interface{}{validityToMap(s.Res.Validity)})
	} else {
		s.D.Set("validity", nil)
	}

	if s.Res.VersionName != nil {
		s.D.Set("version_name", *s.Res.VersionName)
	}

	if s.Res.VersionNumber != nil {
		s.D.Set("version_number", strconv.FormatInt(*s.Res.VersionNumber, 10))
	}

	return nil
}
