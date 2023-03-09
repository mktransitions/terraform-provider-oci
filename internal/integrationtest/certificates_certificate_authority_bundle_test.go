// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CertificatesCertificateAuthorityBundleSingularDataSourceRepresentation = map[string]interface{}{
		"certificate_authority_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_certificates_management_certificate_authority.test_certificate_authority.id}`},
		"stage":                    acctest.Representation{RepType: acctest.Optional, Create: `CURRENT`},
	}

	CertificatesCertificateAuthorityBundleResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority", acctest.Required, acctest.Create, certificateAuthorityRepresentation)
)

// issue-routing-tag: certificates/default
func TestCertificatesCertificateAuthorityBundleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCertificatesCertificateAuthorityBundleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_certificates_certificate_authority_bundle.test_certificate_authority_bundle"

	acctest.SaveConfigContent("", "", "", t)
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_certificate_authority_bundle", "test_certificate_authority_bundle", acctest.Optional, acctest.Create, CertificatesCertificateAuthorityBundleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CertificatesCertificateAuthorityBundleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cert_chain_pem"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_authority_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_authority_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_pem"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "serial_number"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stages.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "validity.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "validity.0.time_of_validity_not_before"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "validity.0.time_of_validity_not_after"),
				resource.TestCheckResourceAttr(singularDatasourceName, "version_number", "1"),
			),
		},
	})
}
