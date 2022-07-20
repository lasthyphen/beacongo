// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avalanche

import (
	"github.com/lasthyphen/beacongo/snow"
	"github.com/lasthyphen/beacongo/snow/consensus/avalanche"
	"github.com/lasthyphen/beacongo/snow/engine/avalanche/vertex"
	"github.com/lasthyphen/beacongo/snow/engine/common"
	"github.com/lasthyphen/beacongo/snow/validators"
)

// Config wraps all the parameters needed for an avalanche engine
type Config struct {
	Ctx *snow.ConsensusContext
	common.AllGetsServer
	VM         vertex.DAGVM
	Manager    vertex.Manager
	Sender     common.Sender
	Validators validators.Set

	Params    avalanche.Parameters
	Consensus avalanche.Consensus
}
