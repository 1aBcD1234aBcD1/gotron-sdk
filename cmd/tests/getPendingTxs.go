package main

import (
	"encoding/hex"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/grpc"
	"sync"
)

func main() {
	grpClient := client.NewGrpcClient("localhost:50051")
	err := grpClient.Start(grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	var locker sync.RWMutex
	m := make(map[string]bool)

	for {
		txsList, err := grpClient.GetTransactionListFromPending()
		if err != nil {
			panic(err)
		}
		for _, t := range txsList.TxId {
			go func(t string) {
				locker.RLock()
				if m[t] {
					locker.RUnlock()
					return
				}
				locker.RUnlock()

				locker.Lock()
				m[t] = true
				locker.Unlock()

				data, err := hex.DecodeString(t)
				if err != nil {
					panic(err)
				}
				tx, err := grpClient.GetTransactionFromPending(data)
				if err != nil {
					panic(err)
				}
				if tx.RawData != nil {
					trigger := core.TriggerSmartContract{}
					err = tx.RawData.Contract[0].Parameter.UnmarshalTo(&trigger)
					if err != nil {
						fmt.Println(fmt.Sprintf("%s no hace nada", t))
					} else {
						fmt.Println(fmt.Sprintf("%s %x", t, trigger.ContractAddress))
					}
				}
			}(t)
		}
	}

	grpClient.Stop()
}
