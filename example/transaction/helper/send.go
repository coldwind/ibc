package helper

import (
	"context"
	"fmt"
	"ibc/example/config"
	"ibc/transaction"
	"ibc/utils"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func SendEth() {
	cli, err := utils.GetClient(config.Rawurl)
	if err != nil {
		fmt.Println("get client error", err)
		return
	}

	tx, err := transaction.SendEth(cli.Cli(), config.PrivateKey, config.ToAccount, big.NewInt(5000000000000000000))
	if err != nil {
		fmt.Println("send error", err)
		return
	}
	fmt.Println("sent", tx.Hash().Hex())

	from := common.HexToAddress(config.FromAccount)
	val, _ := cli.Cli().BalanceAt(context.Background(), from, nil)

	fmt.Println("balance", val.String())
}
