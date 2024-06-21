package main

import "github.com/ethereum/go-ethereum/core/types"

type Thing struct {
	Id     Int
	Status bool
}

func Test() {
	var debug Receipt
	// ruleid: eth-txreceipt-status
	a := debug.Status

	var debug2 Thing
	// ok: eth-txreceipt-status
	b := debug2.Status
}
