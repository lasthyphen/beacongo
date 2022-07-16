// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snow

import (
	"github.com/lasthyphen/beacongo/ids"
)

// Acceptor is implemented when a struct is monitoring if a message is accepted

// Rejector is implemented when a struct is monitoring if a message is rejected
type Rejector interface {
	Reject(ctx *ConsensusContext, containerID ids.ID, container []byte) error
}

// Issuer is implemented when a struct is monitoring if a message is issued
type Issuer interface {
	Issue(ctx *ConsensusContext, containerID ids.ID, container []byte) error
}
