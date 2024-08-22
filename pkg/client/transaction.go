package client

import (
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
)

// GetTransactionSignWeight queries transaction sign weight
func (g *GrpcClient) GetTransactionSignWeight(tx *core.Transaction) (*api.TransactionSignWeight, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	result, err := g.Client.GetTransactionSignWeight(ctx, tx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTransactionListFromPending Query transaction information in the pending pool
func (g *GrpcClient) GetTransactionListFromPending() (*api.TransactionIdList, error) {
	ctx, cancel := g.getContext()
	defer cancel()
	txs, err := g.Client.GetTransactionListFromPending(ctx, &api.EmptyMessage{})
	return txs, err
}

func (g *GrpcClient) GetTransactionFromPending(txHash []byte) (*core.Transaction, error) {
	ctx, cancel := g.getContext()
	defer cancel()
	return g.Client.GetTransactionFromPending(ctx, &api.BytesMessage{Value: txHash})
}
