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
	CertificatesCertificateBundleSingularDataSourceRepresentation = map[string]interface{}{
		"certificate_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_certificates_management_certificate.test_certificate.id}`},
		"certificate_bundle_type": acctest.Representation{RepType: acctest.Optional, Create: `CERTIFICATE_CONTENT_WITH_PRIVATE_KEY`},
		"stage":                   acctest.Representation{RepType: acctest.Optional, Create: `CURRENT`},
	}

	CertificatesCertificateBundleResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Required, acctest.Create, certificatesManagementCertificateRepresentation)
)

// issue-routing-tag: certificates/default
func TestCertificatesCertificateBundleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCertificatesCertificateBundleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_certificates_certificate_bundle.test_certificate_bundle"

	acctest.SaveConfigContent("", "", "", t)
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_certificate_bundle", "test_certificate_bundle", acctest.Optional, acctest.Create, CertificatesCertificateBundleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CertificatesCertificateBundleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cert_chain_pem"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_bundle_type", "CERTIFICATE_CONTENT_WITH_PRIVATE_KEY"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_pem"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_key_pem"),
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
