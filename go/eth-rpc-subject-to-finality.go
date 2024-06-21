package main




func negativeMain() {
	client, err := ethclient.Dial(WssRpcNode)
	if err != nil {
		log.Fatal(err)
	}

	// ok: eth-rpc-subject-to-finality
	block, err := client.BalanceAt(context.Background(), "0x00000d", "0xaccount")

	// ok: eth-rpc-subject-to-finality
	block, err := client.BlockByNumber(context.Background(),100000)

	// ok: eth-rpc-subject-to-finality
	block, err := client.BlockByHash(context.Background(), header.Hash())

	// ok: eth-rpc-subject-to-finality
	block, err := client.BlockReceipts(context.Background(), header.Hash())

	// ok: eth-rpc-subject-to-finality
	block, err := client.FilterLogs(context.Background(), "filtering criteria")

	// ok: eth-rpc-subject-to-finality
	block, err := client.SubscribeFilterLogs(context.Background(), "subscription criteria", "channel")

	// ok: eth-rpc-subject-to-finality
	block, err := client.TransactionByHash(context.Background(), header.Hash())

	// ok: eth-rpc-subject-to-finality
	block, err := client.TransactionReceipt(context.Background(), header.Hash())
}


import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
)

func main() {
	client, err := ethclient.Dial(WssRpcNode)
	if err != nil {
		log.Fatal(err)
	}

	// ruleid: eth-rpc-subject-to-finality
	block, err := client.BalanceAt(context.Background(), "0x00000d", "0xaccount")

	// ruleid: eth-rpc-subject-to-finality
	block, err := client.BlockByNumber(context.Background(),100000)

	// ruleid: eth-rpc-subject-to-finality
	block, err := client.BlockByHash(context.Background(), header.Hash())

	// ruleid: eth-rpc-subject-to-finality
	block, err := client.BlockReceipts(context.Background(), header.Hash())

	// ruleid: eth-rpc-subject-to-finality
	block, err := client.FilterLogs(context.Background(), "filtering criteria")

	// ruleid: eth-rpc-subject-to-finality
	block, err := client.SubscribeFilterLogs(context.Background(), "subscription criteria", "channel")

	// ruleid: eth-rpc-subject-to-finality
	block, err := client.TransactionByHash(context.Background(), header.Hash())

	// ruleid: eth-rpc-subject-to-finality
	block, err := client.TransactionReceipt(context.Background(), header.Hash())
}

func handRolledClient() {

	// ruleid: eth-rpc-subject-to-finality
	reply, err := executeEthRPC(client.JsonRpcEndpoint,"eth_getBalance","[]",)

	// ruleid: eth-rpc-subject-to-finality
	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"eth_getBalance\"}")

	// ruleid: eth-rpc-subject-to-finality
	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByHash\"}")

	// ruleid: eth-rpc-subject-to-finality
	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockReceipts\"}")

	// ruleid: eth-rpc-subject-to-finality
	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\"}")

	// ruleid: eth-rpc-subject-to-finality
	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"eth_blockNumber\"}")

	// ruleid: eth-rpc-subject-to-finality
	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"eth_getLogs\"}")

	// ruleid: eth-rpc-subject-to-finality
	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"eth_getTransactionByHash\"}")

	// ruleid: eth-rpc-subject-to-finality
	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"eth_getTransactionReceipt\"}")

	// ruleid: eth-rpc-subject-to-finality
	sub, err := ec.c.EthSubscribe(ctx, ch, "logs", arg)

	// ruleid: eth-rpc-subject-to-finality
	sub, err := ec.c.EthSubscribe(ctx, ch, "logs")

	// ruleid: eth-rpc-subject-to-finality
	sub, err := ec.c.EthSubscribe(ctx, ch, "newHeads")

	// ok: eth-rpc-subject-to-finality
	sub, err := ec.c.EthSubscribe(ctx, ch, "othersub")

	// ok: eth-rpc-subject-to-finality
	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"blockNumber\"}")
}


