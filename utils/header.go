package utils

import (
	"context"
	"math/big"
)

type Head struct {
	Number     *big.Int // 区块号
	Hash       string   // 当前区块hash值
	ParentHash string   // 上一个区块的hash值
	Difficulty *big.Int // 创建区块的难易程度
	Time       uint64
}

func (h *Client) GetHeaderByNumber(blockNum *big.Int) (*Head, error) {
	header, err := h.cli.HeaderByNumber(context.Background(), blockNum)
	if err != nil {
		return nil, err
	}
	res := &Head{}

	res.Difficulty = header.Difficulty
	res.Number = header.Number
	res.Hash = header.Hash().Hex()
	res.ParentHash = header.ParentHash.Hex()
	res.Time = header.Time

	return res, nil
}

func (h *Client) GetBlock() {

}
