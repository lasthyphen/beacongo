// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"errors"
	"fmt"

	"github.com/lasthyphen/beacongo/ids"
	"github.com/lasthyphen/beacongo/utils/crypto"
	"github.com/lasthyphen/beacongo/utils/hashing"
	"github.com/lasthyphen/beacongo/utils/math"
	"github.com/lasthyphen/beacongo/vms/components/djtx"
	"github.com/lasthyphen/beacongo/vms/components/verify"
	"github.com/lasthyphen/beacongo/vms/platformvm/fx"
	"github.com/lasthyphen/beacongo/vms/platformvm/stakeable"
	"github.com/lasthyphen/beacongo/vms/secp256k1fx"
)

var (
	errLockedFundsNotMarkedAsLocked = errors.New("locked funds not marked as locked")
	errWrongLocktime                = errors.New("wrong locktime reported")
	errUnknownOwners                = errors.New("unknown owners")
	errCantSign                     = errors.New("can't sign")
)

// stake the provided amount while deducting the provided fee.
// Arguments:
// - [keys] are the owners of the funds
// - [amount] is the amount of funds that are trying to be staked
// - [fee] is the amount of DJTX that should be burned
// - [changeAddr] is the address that change, if there is any, is sent to
// Returns:
// - [inputs] the inputs that should be consumed to fund the outputs
// - [returnedOutputs] the outputs that should be immediately returned to the
//                     UTXO set
// - [stakedOutputs] the outputs that should be locked for the duration of the
//                   staking period
// - [signers] the proof of ownership of the funds being moved
func (vm *VM) stake(
	keys []*crypto.PrivateKeySECP256K1R,
	amount uint64,
	fee uint64,
	changeAddr ids.ShortID,
) (
	[]*djtx.TransferableInput, // inputs
	[]*djtx.TransferableOutput, // returnedOutputs
	[]*djtx.TransferableOutput, // stakedOutputs
	[][]*crypto.PrivateKeySECP256K1R, // signers
	error,
) {
	addrs := ids.NewShortSet(len(keys)) // The addresses controlled by [keys]
	for _, key := range keys {
		addrs.Add(key.PublicKey().Address())
	}
	utxos, err := djtx.GetAllUTXOs(vm.internalState, addrs) // The UTXOs controlled by [keys]
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("couldn't get UTXOs: %w", err)
	}

	kc := secp256k1fx.NewKeychain(keys...) // Keychain consumes UTXOs and creates new ones

	// Minimum time this transaction will be issued at
	now := uint64(vm.clock.Time().Unix())

	ins := []*djtx.TransferableInput{}
	returnedOuts := []*djtx.TransferableOutput{}
	stakedOuts := []*djtx.TransferableOutput{}
	signers := [][]*crypto.PrivateKeySECP256K1R{}

	// Amount of DJTX that has been staked
	amountStaked := uint64(0)

	// Consume locked UTXOs
	for _, utxo := range utxos {
		// If we have consumed more DJTX than we are trying to stake, then we
		// have no need to consume more locked DJTX
		if amountStaked >= amount {
			break
		}

		if assetID := utxo.AssetID(); assetID != vm.ctx.DJTXAssetID {
			continue // We only care about staking DJTX, so ignore other assets
		}

		out, ok := utxo.Out.(*stakeable.LockOut)
		if !ok {
			// This output isn't locked, so it will be handled during the next
			// iteration of the UTXO set
			continue
		}
		if out.Locktime <= now {
			// This output is no longer locked, so it will be handled during the
			// next iteration of the UTXO set
			continue
		}

		inner, ok := out.TransferableOut.(*secp256k1fx.TransferOutput)
		if !ok {
			// We only know how to clone secp256k1 outputs for now
			continue
		}

		inIntf, inSigners, err := kc.Spend(out.TransferableOut, now)
		if err != nil {
			// We couldn't spend the output, so move on to the next one
			continue
		}
		in, ok := inIntf.(djtx.TransferableIn)
		if !ok { // should never happen
			vm.ctx.Log.Warn("expected input to be djtx.TransferableIn but is %T", inIntf)
			continue
		}

		// The remaining value is initially the full value of the input
		remainingValue := in.Amount()

		// Stake any value that should be staked
		amountToStake := math.Min64(
			amount-amountStaked, // Amount we still need to stake
			remainingValue,      // Amount available to stake
		)
		amountStaked += amountToStake
		remainingValue -= amountToStake

		// Add the input to the consumed inputs
		ins = append(ins, &djtx.TransferableInput{
			UTXOID: utxo.UTXOID,
			Asset:  djtx.Asset{ID: vm.ctx.DJTXAssetID},
			In: &stakeable.LockIn{
				Locktime:       out.Locktime,
				TransferableIn: in,
			},
		})

		// Add the output to the staked outputs
		stakedOuts = append(stakedOuts, &djtx.TransferableOutput{
			Asset: djtx.Asset{ID: vm.ctx.DJTXAssetID},
			Out: &stakeable.LockOut{
				Locktime: out.Locktime,
				TransferableOut: &secp256k1fx.TransferOutput{
					Amt:          amountToStake,
					OutputOwners: inner.OutputOwners,
				},
			},
		})

		if remainingValue > 0 {
			// This input provided more value than was needed to be locked.
			// Some of it must be returned
			returnedOuts = append(returnedOuts, &djtx.TransferableOutput{
				Asset: djtx.Asset{ID: vm.ctx.DJTXAssetID},
				Out: &stakeable.LockOut{
					Locktime: out.Locktime,
					TransferableOut: &secp256k1fx.TransferOutput{
						Amt:          remainingValue,
						OutputOwners: inner.OutputOwners,
					},
				},
			})
		}

		// Add the signers needed for this input to the set of signers
		signers = append(signers, inSigners)
	}

	// Amount of DJTX that has been burned
	amountBurned := uint64(0)

	for _, utxo := range utxos {
		// If we have consumed more DJTX than we are trying to stake, and we
		// have burned more DJTX then we need to, then we have no need to
		// consume more DJTX
		if amountBurned >= fee && amountStaked >= amount {
			break
		}

		if assetID := utxo.AssetID(); assetID != vm.ctx.DJTXAssetID {
			continue // We only care about burning DJTX, so ignore other assets
		}

		out := utxo.Out
		inner, ok := out.(*stakeable.LockOut)
		if ok {
			if inner.Locktime > now {
				// This output is currently locked, so this output can't be
				// burned. Additionally, it may have already been consumed
				// above. Regardless, we skip to the next UTXO
				continue
			}
			out = inner.TransferableOut
		}

		inIntf, inSigners, err := kc.Spend(out, now)
		if err != nil {
			// We couldn't spend this UTXO, so we skip to the next one
			continue
		}
		in, ok := inIntf.(djtx.TransferableIn)
		if !ok {
			// Because we only use the secp Fx right now, this should never
			// happen
			continue
		}

		// The remaining value is initially the full value of the input
		remainingValue := in.Amount()

		// Burn any value that should be burned
		amountToBurn := math.Min64(
			fee-amountBurned, // Amount we still need to burn
			remainingValue,   // Amount available to burn
		)
		amountBurned += amountToBurn
		remainingValue -= amountToBurn

		// Stake any value that should be staked
		amountToStake := math.Min64(
			amount-amountStaked, // Amount we still need to stake
			remainingValue,      // Amount available to stake
		)
		amountStaked += amountToStake
		remainingValue -= amountToStake

		// Add the input to the consumed inputs
		ins = append(ins, &djtx.TransferableInput{
			UTXOID: utxo.UTXOID,
			Asset:  djtx.Asset{ID: vm.ctx.DJTXAssetID},
			In:     in,
		})

		if amountToStake > 0 {
			// Some of this input was put for staking
			stakedOuts = append(stakedOuts, &djtx.TransferableOutput{
				Asset: djtx.Asset{ID: vm.ctx.DJTXAssetID},
				Out: &secp256k1fx.TransferOutput{
					Amt: amountToStake,
					OutputOwners: secp256k1fx.OutputOwners{
						Locktime:  0,
						Threshold: 1,
						Addrs:     []ids.ShortID{changeAddr},
					},
				},
			})
		}

		if remainingValue > 0 {
			// This input had extra value, so some of it must be returned
			returnedOuts = append(returnedOuts, &djtx.TransferableOutput{
				Asset: djtx.Asset{ID: vm.ctx.DJTXAssetID},
				Out: &secp256k1fx.TransferOutput{
					Amt: remainingValue,
					OutputOwners: secp256k1fx.OutputOwners{
						Locktime:  0,
						Threshold: 1,
						Addrs:     []ids.ShortID{changeAddr},
					},
				},
			})
		}

		// Add the signers needed for this input to the set of signers
		signers = append(signers, inSigners)
	}

	if amountBurned < fee || amountStaked < amount {
		return nil, nil, nil, nil, fmt.Errorf(
			"provided keys have balance (unlocked, locked) (%d, %d) but need (%d, %d)",
			amountBurned, amountStaked, fee, amount)
	}

	djtx.SortTransferableInputsWithSigners(ins, signers) // sort inputs and keys
	djtx.SortTransferableOutputs(returnedOuts, Codec)    // sort outputs
	djtx.SortTransferableOutputs(stakedOuts, Codec)      // sort outputs

	return ins, returnedOuts, stakedOuts, signers, nil
}

// authorize an operation on behalf of the named subnet with the provided keys.
func (vm *VM) authorize(
	vs MutableState,
	subnetID ids.ID,
	keys []*crypto.PrivateKeySECP256K1R,
) (
	verify.Verifiable, // Input that names owners
	[]*crypto.PrivateKeySECP256K1R, // Keys that prove ownership
	error,
) {
	subnetTx, _, err := vs.GetTx(subnetID)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"failed to fetch subnet %s: %w",
			subnetID,
			err,
		)
	}
	subnet, ok := subnetTx.UnsignedTx.(*UnsignedCreateSubnetTx)
	if !ok {
		return nil, nil, errWrongTxType
	}

	// Make sure the owners of the subnet match the provided keys
	owner, ok := subnet.Owner.(*secp256k1fx.OutputOwners)
	if !ok {
		return nil, nil, errUnknownOwners
	}

	// Add the keys to a keychain
	kc := secp256k1fx.NewKeychain(keys...)

	// Make sure that the operation is valid after a minimum time
	now := uint64(vm.clock.Time().Unix())

	// Attempt to prove ownership of the subnet
	indices, signers, matches := kc.Match(owner, now)
	if !matches {
		return nil, nil, errCantSign
	}

	return &secp256k1fx.Input{SigIndices: indices}, signers, nil
}

// Verify that [tx] is semantically valid.
// [db] should not be committed if an error is returned
// [ins] and [outs] are the inputs and outputs of [tx].
// [creds] are the credentials of [tx], which allow [ins] to be spent.
// Precondition: [tx] has already been syntactically verified
func (vm *VM) semanticVerifySpend(
	utxoDB UTXOGetter,
	tx UnsignedTx,
	ins []*djtx.TransferableInput,
	outs []*djtx.TransferableOutput,
	creds []verify.Verifiable,
	feeAmount uint64,
	feeAssetID ids.ID,
) error {
	utxos := make([]*djtx.UTXO, len(ins))
	for index, input := range ins {
		utxo, err := utxoDB.GetUTXO(input.InputID())
		if err != nil {
			return fmt.Errorf(
				"failed to read consumed UTXO %s due to: %w",
				&input.UTXOID,
				err,
			)
		}
		utxos[index] = utxo
	}

	return vm.semanticVerifySpendUTXOs(tx, utxos, ins, outs, creds, feeAmount, feeAssetID)
}

// Verify that [tx] is semantically valid.
// [db] should not be committed if an error is returned
// [ins] and [outs] are the inputs and outputs of [tx].
// [creds] are the credentials of [tx], which allow [ins] to be spent.
// [utxos[i]] is the UTXO being consumed by [ins[i]]
// Precondition: [tx] has already been syntactically verified
func (vm *VM) semanticVerifySpendUTXOs(
	tx UnsignedTx,
	utxos []*djtx.UTXO,
	ins []*djtx.TransferableInput,
	outs []*djtx.TransferableOutput,
	creds []verify.Verifiable,
	feeAmount uint64,
	feeAssetID ids.ID,
) error {
	if len(ins) != len(creds) {
		return fmt.Errorf(
			"there are %d inputs but %d credentials. Should be same number",
			len(ins),
			len(creds),
		)
	}
	if len(ins) != len(utxos) {
		return fmt.Errorf(
			"there are %d inputs but %d utxos. Should be same number",
			len(ins),
			len(utxos),
		)
	}
	for _, cred := range creds { // Verify credentials are well-formed.
		if err := cred.Verify(); err != nil {
			return err
		}
	}

	// Time this transaction is being verified
	now := uint64(vm.clock.Time().Unix())

	// Track the amount of unlocked transfers
	unlockedProduced := feeAmount
	unlockedConsumed := uint64(0)

	// Track the amount of locked transfers and their owners
	// locktime -> ownerID -> amount
	lockedProduced := make(map[uint64]map[ids.ID]uint64)
	lockedConsumed := make(map[uint64]map[ids.ID]uint64)

	for index, input := range ins {
		utxo := utxos[index] // The UTXO consumed by [input]

		if assetID := utxo.AssetID(); assetID != feeAssetID {
			return errAssetIDMismatch
		}
		if assetID := input.AssetID(); assetID != feeAssetID {
			return errAssetIDMismatch
		}

		out := utxo.Out
		locktime := uint64(0)
		// Set [locktime] to this UTXO's locktime, if applicable
		if inner, ok := out.(*stakeable.LockOut); ok {
			out = inner.TransferableOut
			locktime = inner.Locktime
		}

		in := input.In
		// The UTXO says it's locked until [locktime], but this input, which
		// consumes it, is not locked even though [locktime] hasn't passed. This
		// is invalid.
		if inner, ok := in.(*stakeable.LockIn); now < locktime && !ok {
			return errLockedFundsNotMarkedAsLocked
		} else if ok {
			if inner.Locktime != locktime {
				// This input is locked, but its locktime is wrong
				return errWrongLocktime
			}
			in = inner.TransferableIn
		}

		// Verify that this tx's credentials allow [in] to be spent
		if err := vm.fx.VerifyTransfer(tx, in, creds[index], out); err != nil {
			return fmt.Errorf("failed to verify transfer: %w", err)
		}

		amount := in.Amount()

		if now >= locktime {
			newUnlockedConsumed, err := math.Add64(unlockedConsumed, amount)
			if err != nil {
				return err
			}
			unlockedConsumed = newUnlockedConsumed
			continue
		}

		owned, ok := out.(fx.Owned)
		if !ok {
			return errUnknownOwners
		}
		owner := owned.Owners()
		ownerBytes, err := Codec.Marshal(CodecVersion, owner)
		if err != nil {
			return fmt.Errorf("couldn't marshal owner: %w", err)
		}
		ownerID := hashing.ComputeHash256Array(ownerBytes)
		owners, ok := lockedConsumed[locktime]
		if !ok {
			owners = make(map[ids.ID]uint64)
			lockedConsumed[locktime] = owners
		}
		newAmount, err := math.Add64(owners[ownerID], amount)
		if err != nil {
			return err
		}
		owners[ownerID] = newAmount
	}

	for _, out := range outs {
		if assetID := out.AssetID(); assetID != feeAssetID {
			return errAssetIDMismatch
		}

		output := out.Output()
		locktime := uint64(0)
		// Set [locktime] to this output's locktime, if applicable
		if inner, ok := output.(*stakeable.LockOut); ok {
			output = inner.TransferableOut
			locktime = inner.Locktime
		}

		amount := output.Amount()

		if locktime == 0 {
			newUnlockedProduced, err := math.Add64(unlockedProduced, amount)
			if err != nil {
				return err
			}
			unlockedProduced = newUnlockedProduced
			continue
		}

		owned, ok := output.(fx.Owned)
		if !ok {
			return errUnknownOwners
		}
		owner := owned.Owners()
		ownerBytes, err := Codec.Marshal(CodecVersion, owner)
		if err != nil {
			return fmt.Errorf("couldn't marshal owner: %w", err)
		}
		ownerID := hashing.ComputeHash256Array(ownerBytes)
		owners, ok := lockedProduced[locktime]
		if !ok {
			owners = make(map[ids.ID]uint64)
			lockedProduced[locktime] = owners
		}
		newAmount, err := math.Add64(owners[ownerID], amount)
		if err != nil {
			return err
		}
		owners[ownerID] = newAmount
	}

	// Make sure that for each locktime, tokens produced <= tokens consumed
	for locktime, producedAmounts := range lockedProduced {
		consumedAmounts := lockedConsumed[locktime]
		for ownerID, producedAmount := range producedAmounts {
			consumedAmount := consumedAmounts[ownerID]

			if producedAmount > consumedAmount {
				increase := producedAmount - consumedAmount
				if increase > unlockedConsumed {
					return fmt.Errorf(
						"address %s produces %d unlocked and consumes %d unlocked for locktime %d",
						ownerID,
						increase,
						unlockedConsumed,
						locktime,
					)
				}
				unlockedConsumed -= increase
			}
		}
	}

	// More unlocked tokens produced than consumed. Invalid.
	if unlockedProduced > unlockedConsumed {
		return fmt.Errorf(
			"tx produces more unlocked (%d) than it consumes (%d)",
			unlockedProduced,
			unlockedConsumed,
		)
	}
	return nil
}

// Removes the UTXOs consumed by [ins] from the UTXO set
func consumeInputs(
	utxoDB UTXODeleter,
	ins []*djtx.TransferableInput,
) {
	for _, input := range ins {
		utxoDB.DeleteUTXO(input.InputID())
	}
}

// Adds the UTXOs created by [outs] to the UTXO set.
// [txID] is the ID of the tx that created [outs].
func produceOutputs(
	utxoDB UTXOAdder,
	txID ids.ID,
	assetID ids.ID,
	outs []*djtx.TransferableOutput,
) {
	for index, out := range outs {
		utxoDB.AddUTXO(&djtx.UTXO{
			UTXOID: djtx.UTXOID{
				TxID:        txID,
				OutputIndex: uint32(index),
			},
			Asset: djtx.Asset{ID: assetID},
			Out:   out.Output(),
		})
	}
}
