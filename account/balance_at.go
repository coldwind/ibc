package account

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Account struct {
}

func (h *Account) BalanceAt(cli *ethclient.Client, hexAddress string) (*big.Int, error) {
	address := common.HexToAddress(hexAddress)
	return cli.BalanceAt(context.Background(), address, nil)
}
