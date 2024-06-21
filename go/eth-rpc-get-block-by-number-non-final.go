package main

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc/types"
)

func Main() {
	rpcURL := "https://mainnet.infura.io/v3/YOUR_INFURA_API"

	// Note these clients are often rolled in a separate function or implemented by a non-geth library, so we can't
	// really scope in on the package imports.
	client, err := ethclient.Dial(rpcURL)

	// ruleid: eth-rpc-get-block-by-number-non-final
	currentBlock, err := client.BlockByNumber(context.Background(), types.LatestBlockNumber)

	// ruleid: eth-rpc-get-block-by-number-non-final
	currentBlock, err := client.BlockByNumber(context.Background(), types.PendingBlockNumber)

	// ruleid: eth-rpc-get-block-by-number-non-final
	currentBlock, err := client.BlockByNumber(context.Background(), types.EarliestBlockNumber)

	// ruleid: eth-rpc-get-block-by-number-non-final
	currentBlock, err := client.BlockByNumber(context.Background(), nil)

	// ruleid: eth-rpc-get-block-by-number-non-final
	currentBlock, err := client.BlockByNumber(context.Background(), 5000000)

	// ok: eth-rpc-get-block-by-number-non-final
	currentBlock, err := client.BlockByNumber(context.Background(), types.FinalizedBlockNumber)

	// ok: eth-rpc-get-block-by-number-non-final
	currentBlock, err := client.BlockByNumber(context.Background(), types.SafeBlockNumber)


}
