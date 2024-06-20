package main

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

}
