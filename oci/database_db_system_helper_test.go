// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/oci-go-sdk/v25/common"
	oci_database "github.com/oracle/oci-go-sdk/v25/database"
)

func sweepDatabaseDbSystemResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	dbSystemIds, err := getDbSystemIds(compartment)
	if err != nil {
		return err
	}
	for _, dbSystemId := range dbSystemIds {
		if ok := SweeperDefaultResourceId[dbSystemId]; !ok {
			terminateDbSystemRequest := oci_database.TerminateDbSystemRequest{}

			terminateDbSystemRequest.DbSystemId = &dbSystemId

			terminateDbSystemRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.TerminateDbSystem(context.Background(), terminateDbSystemRequest)
			if error != nil {
				fmt.Printf("Error deleting DbSystem %s %s, It is possible that the resource is already deleted. Please verify manually \n", dbSystemId, error)
				continue
			}
			waitTillCondition(testAccProvider, &dbSystemId, dbSystemSweepWaitCondition, time.Duration(3*time.Minute),
				dbSystemSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDbSystemIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "DbSystemId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listDbSystemsRequest := oci_database.ListDbSystemsRequest{}
	listDbSystemsRequest.CompartmentId = &compartmentId
	listDbSystemsRequest.LifecycleState = oci_database.DbSystemSummaryLifecycleStateAvailable

	// Terminate the newest dbSystem first to make sure any standby databases created by Data Guard Assocuations are deleted first
	listDbSystemsRequest.SortBy = oci_database.ListDbSystemsSortByTimecreated
	listDbSystemsRequest.SortOrder = oci_database.ListDbSystemsSortOrderDesc

	listDbSystemsResponse, err := databaseClient.ListDbSystems(context.Background(), listDbSystemsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DbSystem list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dbSystem := range listDbSystemsResponse.Items {
		id := *dbSystem.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "DbSystemId", id)
	}
	return resourceIds, nil
}

func dbSystemSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dbSystemResponse, ok := response.Response.(oci_database.GetDbSystemResponse); ok {
		return dbSystemResponse.LifecycleState != oci_database.DbSystemLifecycleStateTerminated
	}
	return false
}

func dbSystemSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetDbSystem(context.Background(), oci_database.GetDbSystemRequest{
		DbSystemId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func getDbNodeIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "DbNodeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	dbSystemIds, err := getDbSystemIds(compartment)
	if err != nil {
		return resourceIds, fmt.Errorf("Error getting dbSystemId required for DbNode resource requests \n")
	}
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	for _, dbSystemId := range dbSystemIds {
		listDbNodesRequest := oci_database.ListDbNodesRequest{}
		listDbNodesRequest.CompartmentId = &compartmentId
		listDbNodesRequest.DbSystemId = &dbSystemId
		listDbNodesResponse, err := databaseClient.ListDbNodes(context.Background(), listDbNodesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DbSystem list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, dbNode := range listDbNodesResponse.Items {
			id := *dbNode.Id
			resourceIds = append(resourceIds, id)
			addResourceIdToSweeperResourceIdMap(compartmentId, "DbNodeId", id)
		}
	}
	return resourceIds, nil
}
