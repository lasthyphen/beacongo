// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package resource

// NoUsage implements Usage() by always returning 0.
var NoUsage User = noUsage{}

type noUsage struct{}

func (noUsage) CPUUsage() float64 { return 0 }

func (noUsage) DiskUsage() (float64, float64) { return 0, 0 }