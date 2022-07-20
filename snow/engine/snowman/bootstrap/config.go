// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrap

import (
	"github.com/lasthyphen/beacongo/snow/engine/common"
	"github.com/lasthyphen/beacongo/snow/engine/common/queue"
	"github.com/lasthyphen/beacongo/snow/engine/common/tracker"
	"github.com/lasthyphen/beacongo/snow/engine/snowman/block"
)

type Config struct {
	common.Config
	common.AllGetsServer

	// Blocked tracks operations that are blocked on blocks
	Blocked *queue.JobsWithMissing

	VM            block.ChainVM
	WeightTracker tracker.WeightTracker

	Bootstrapped func()
}
