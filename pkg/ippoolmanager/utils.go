// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ippoolmanager

import (
	"github.com/spidernet-io/spiderpool/pkg/constant"
	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
)

// TODO(Icarus9913): Deprecated.
func IsAutoCreatedIPPool(pool *spiderpoolv2beta1.SpiderIPPool) bool {
	// only the auto-created IPPool owns the annotation "ipam.spidernet.io/application"
	poolAnno := pool.GetAnnotations()
	_, ok := poolAnno[constant.AnnoSpiderSubnetPoolApp]
	return ok
}
