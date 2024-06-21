package main

import alias "github.com/ethereum/go-ethereum/core"
import "github.com/ethereum/go-ethereum/core"
import ethtypes "github.com/ethereum/go-ethereum/core/types"

type TransferFunc func(StateDB, common.Address, common.Address, *uint256.Int)
type BlockContext struct {
	// CanTransfer returns whether the account contains
	// sufficient ether to transfer the value
	CanTransfer CanTransferFunc
	// Transfer transfers ether from one account to the other
	Transfer TransferFunc
}
type EVM struct {
		// Context provides auxiliary blockchain related information
		Context BlockContext
}


func (evm *EVM) Call(caller ContractRef, addr common.Address, input []byte, gas uint64, value *uint256.Int) (ret []byte, leftOverGas uint64, err error) {

	// from geth
	// ruleid: eth-geth-internal-value-transfer
	evm.Context.Transfer(evm.StateDB, caller.Address(), addr, value)

	// from sei: https://github.com/sei-protocol/sei-chain/blob/c5610ba293d141cc037cd543ef5a3140a68b27a9/x/evm/keeper/keeper.go#L201-L207
	txfer := func(db vm.StateDB, sender, recipient common.Address, amount *big.Int) {
		if IsPayablePrecompile(&recipient) {
			state.TransferWithoutEvents(db, sender, recipient, amount)
		} else {
			// ruleid: eth-geth-internal-value-transfer
			alias.Transfer(db, sender, recipient, amount)

			// ruleid: eth-geth-internal-value-transfer
			core.Transfer(db, sender, recipient, amount)
		}
	}

	// ok: eth-geth-internal-value-transfer
	evm.Context.Transfer(evm.StateDB, caller.Address(), addr, value, extra)

	// ok: eth-geth-internal-value-transfer
	ethtypes.Transfer(db, sender, recipient, amount)

	return nil, 0, nil
}
