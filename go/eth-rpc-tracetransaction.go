package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Test() {
	// ruleid: eth-rpc-tracetransaction
	data, err := client.TraceTransaction(ctx, "hash", nil)
	// ruleid: eth-rpc-tracetransaction
	data, err := client.TraceBlockByNumber(ctx, 5, nil)
	// ruleid: eth-rpc-tracetransaction
	data, err := client.TraceBlockByHash(ctx, []byte{0x05}, nil)
	// ruleid: eth-rpc-tracetransaction
	data, err := client.TraceBlock(ctx, []byte{0x05}, nil)
	// ruleid: eth-rpc-tracetransaction
	data, err := client.TraceChain(ctx, 5, nil)

	url := "https://eth-mainnet.g.alchemy.com/v2/docs-demo"
	// ruleid: eth-rpc-tracetransaction
	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"trace_transaction\",\"params\":[\"0x8fc90a6c3ee3001cdcbbb685b4fbe67b1fa2bec575b15b0395fea5540d0901ae\"]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

	// ok: eth-rpc-tracetransaction
	data, err := client.TraceSomething(ctx, 5, nil)
	// ok: eth-rpc-tracetransaction
	data, err := client.TraceTransaction(ctx, "hash")
	// ok: eth-rpc-tracetransaction
	data, err := client.TraceBlockByNumber(ctx, 5)
	// ok: eth-rpc-tracetransaction
	data, err := client.TraceBlockByHash(ctx, []byte{0x05})
	// ok: eth-rpc-tracetransaction
	data, err := client.TraceBlock(ctx, []byte{0x05})
}
