package utils

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

type Block struct {
	Number       *big.Int
	Hash         string
	Time         uint64
	Transactions types.Transactions
}

func (h *Client) GetBlockByNumber(number *big.Int) (*Block, error) {
	block, err := h.cli.BlockByNumber(context.Background(), number)
	if err != nil {
		return nil, err
	}

	res := &Block{}
	res.Number = block.Number()
	res.Hash = block.Hash().String()
	res.Time = block.Time()
	res.Transactions = block.Transactions()

	return res, nil
}
