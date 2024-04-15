package helper

import (
	"fmt"
	"ibc/account"
	"ibc/example/config"
	"ibc/utils"
	"math"
	"math/big"
)

func BalanceAt() {
	fmt.Println("call BalanceAt")

	// get client
	cli, err := utils.GetClient(config.Rawurl)
	if err != nil {
		fmt.Println("get client error", err)
		return
	}

	accountHandle := account.Account{}

	// balance at
	balance, err := accountHandle.BalanceAt(cli.Cli(), config.FromAccount)
	if err != nil {
		fmt.Println("BalanceAt error", err)
		return
	}

	fmt.Println("balance at", balance.String(), "Wei")

	fbalance := big.NewFloat(0)
	fbalance.SetString(balance.String())
	fbalance.Quo(fbalance, big.NewFloat(math.Pow(10, 18)))
	fmt.Println("eth", fbalance)
}
