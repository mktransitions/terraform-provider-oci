// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_blockchain "github.com/oracle/oci-go-sdk/v25/blockchain"
)

func init() {
	RegisterDatasource("oci_blockchain_peers", BlockchainPeersDataSource())
}

func BlockchainPeersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBlockchainPeers,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"blockchain_platform_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"peer_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     BlockchainPeerResource(),
						},
					},
				},
			},
		},
	}
}

func readBlockchainPeers(d *schema.ResourceData, m interface{}) error {
	sync := &BlockchainPeersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockchainPlatformClient()

	return ReadResource(sync)
}

type BlockchainPeersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_blockchain.BlockchainPlatformClient
	Res    *oci_blockchain.ListPeersResponse
}

func (s *BlockchainPeersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BlockchainPeersDataSourceCrud) Get() error {
	request := oci_blockchain.ListPeersRequest{}

	if blockchainPlatformId, ok := s.D.GetOkExists("blockchain_platform_id"); ok {
		tmp := blockchainPlatformId.(string)
		request.BlockchainPlatformId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "blockchain")

	response, err := s.Client.ListPeers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPeers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BlockchainPeersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}
	peer := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PeerSummaryToMap(item))
	}
	peer["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, BlockchainPeersDataSource().Schema["peer_collection"].Elem.(*schema.Resource).Schema)
		peer["items"] = items
	}

	resources = append(resources, peer)
	if err := s.D.Set("peer_collection", resources); err != nil {
		return err
	}

	return nil
}
