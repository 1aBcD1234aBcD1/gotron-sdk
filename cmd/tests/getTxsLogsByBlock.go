package main

import (
	"bytes"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"google.golang.org/grpc"
)

func main() {
	grpClient := client.NewGrpcClient("localhost:50051")
	err := grpClient.Start(grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	blockId := []byte{}
	for {
		nowBlock, err := grpClient.GetNowBlock()
		if err != nil {
			panic(err)
		}

		if !bytes.Equal(nowBlock.Blockid, blockId) {
			blockId = nowBlock.Blockid
			fmt.Println(fmt.Sprintf("New block kek: 0x%x", blockId))
			for _, tx := range nowBlock.Transactions {
				fmt.Println(fmt.Sprintf("%x: %d", tx.Txid, len(tx.Logs)))
			}
		}

	}

	grpClient.Stop()
}
