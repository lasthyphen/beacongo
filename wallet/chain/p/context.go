// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package p

import (
	stdcontext "context"

	"github.com/lasthyphen/beacongo/api/info"
	"github.com/lasthyphen/beacongo/ids"
	"github.com/lasthyphen/beacongo/utils/constants"
	"github.com/lasthyphen/beacongo/vms/avm"
)

var _ Context = &context{}

type Context interface {
	NetworkID() uint32
	HRP() string
	DJTXAssetID() ids.ID
	BaseTxFee() uint64
	CreateSubnetTxFee() uint64
	CreateBlockchainTxFee() uint64
}

type context struct {
	networkID             uint32
	hrp                   string
	djtxAssetID           ids.ID
	baseTxFee             uint64
	createSubnetTxFee     uint64
	createBlockchainTxFee uint64
}

func NewContextFromURI(ctx stdcontext.Context, uri string) (Context, error) {
	infoClient := info.NewClient(uri)
	xChainClient := avm.NewClient(uri, "X")
	return NewContextFromClients(ctx, infoClient, xChainClient)
}

func NewContextFromClients(
	ctx stdcontext.Context,
	infoClient info.Client,
	xChainClient avm.Client,
) (Context, error) {
	networkID, err := infoClient.GetNetworkID(ctx)
	if err != nil {
		return nil, err
	}

	asset, err := xChainClient.GetAssetDescription(ctx, "DJTX")
	if err != nil {
		return nil, err
	}

	txFees, err := infoClient.GetTxFee(ctx)
	if err != nil {
		return nil, err
	}

	return NewContext(
		networkID,
		asset.AssetID,
		uint64(txFees.TxFee),
		uint64(txFees.CreateSubnetTxFee),
		uint64(txFees.CreateBlockchainTxFee),
	), nil
}

func NewContext(
	networkID uint32,
	djtxAssetID ids.ID,
	baseTxFee uint64,
	createSubnetTxFee uint64,
	createBlockchainTxFee uint64,
) Context {
	return &context{
		networkID:             networkID,
		hrp:                   constants.GetHRP(networkID),
		djtxAssetID:           djtxAssetID,
		baseTxFee:             baseTxFee,
		createSubnetTxFee:     createSubnetTxFee,
		createBlockchainTxFee: createBlockchainTxFee,
	}
}

func (c *context) NetworkID() uint32             { return c.networkID }
func (c *context) HRP() string                   { return c.hrp }
func (c *context) DJTXAssetID() ids.ID           { return c.djtxAssetID }
func (c *context) BaseTxFee() uint64             { return c.baseTxFee }
func (c *context) CreateSubnetTxFee() uint64     { return c.createSubnetTxFee }
func (c *context) CreateBlockchainTxFee() uint64 { return c.createBlockchainTxFee }
