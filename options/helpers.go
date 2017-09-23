// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package options

import "github.com/oracle/bmcs-go-sdk"

type resourceProvider interface {
	GetOk(string) (interface{}, bool)
}

func SetLimitOptions(resource resourceProvider, opts *baremetal.LimitListOptions) {
	if val, ok := resource.GetOk("limit"); ok {
		opts.Limit = uint64(val.(int))
	}
	return
}

func SetPageOptions(resource resourceProvider, opts *baremetal.PageListOptions) {
	if val, ok := resource.GetOk("page"); ok {
		opts.Page = val.(string)
	}
	return
}

func SetListOptions(resource resourceProvider, opts *baremetal.ListOptions) {
	SetLimitOptions(resource, &opts.LimitListOptions)
	SetPageOptions(resource, &opts.PageListOptions)
	return
}

func SetNextPageOption(nextPage string, opts *baremetal.PageListOptions) (hasNextPage bool) {
	if nextPage == "" {
		hasNextPage = false
	} else {
		hasNextPage = true
		opts.Page = nextPage
	}

	return
}
