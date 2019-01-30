// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	clusterKubeConfigSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_id":    Representation{repType: Required, create: `${oci_containerengine_cluster.test_cluster.id}`},
		"expiration":    Representation{repType: Optional, create: `2592000`},
		"token_version": Representation{repType: Optional, create: `1.0.0`},
	}

	ClusterKubeConfigResourceConfig = ClusterResourceConfig
)

func TestContainerengineClusterKubeConfigResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_containerengine_cluster_kube_config.test_cluster_kube_config"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_containerengine_cluster_kube_config", "test_cluster_kube_config", Optional, Create, clusterKubeConfigSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ClusterKubeConfigResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "expiration", "2592000"),
					resource.TestCheckResourceAttr(singularDatasourceName, "token_version", "1.0.0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				),
			},
		},
	})
}
