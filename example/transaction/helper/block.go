package helper

import (
	"context"
	"fmt"
	"ibc/example/config"
	"ibc/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func BlockByNumber() {
	fmt.Println("call BlockByNumber")
	// get client
	cli, err := utils.GetClient(config.Rawurl)
	if err != nil {
		fmt.Println("get client error", err)
		return
	}

	// chainId
	chainId, err := cli.Cli().NetworkID(context.Background())
	fmt.Println("chainId", chainId)
	if err != nil {
		fmt.Println("get chain Id error", err)
		return
	}

	block, err := cli.GetBlockByNumber(nil)
	if err != nil {
		fmt.Println("get header error", err)
		return
	}
	fmt.Println("number", block.Number)
	fmt.Println("hash", block.Hash)
	fmt.Println("time", block.Time)

	for _, tx := range block.Transactions {
		fmt.Println("----------------------------------------------")
		fmt.Println("transaction hash hex", tx.Hash().Hex())
		fmt.Println("transaction eth amount", tx.Value().String())
		fmt.Println("transaction Gas", tx.Gas())
		fmt.Println("transaction GasPrice", tx.GasPrice().Uint64())
		fmt.Println("transaction Nonce", tx.Nonce())
		fmt.Println("transaction data", tx.Data())
		from, err := types.Sender(types.NewLondonSigner(chainId), tx)
		if err != nil {
			fmt.Println("no from", err)
		} else {
			fmt.Println("transaction from", from.Hex())
		}
		fmt.Println("transaction To", tx.To().Hex())
		if _, isPending, err := cli.Cli().TransactionByHash(context.Background(), tx.Hash()); err == nil {
			fmt.Println("transaction isPending", isPending)
		}

		// status
		recept, err := cli.Cli().TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			fmt.Println("no status", err)
		} else {
			fmt.Println("transaction status", recept.Status)
		}
	}

	transactionCount, _ := cli.Cli().TransactionCount(context.Background(), common.HexToHash(block.Hash))
	fmt.Println("transaction count", transactionCount)
	fmt.Println("transaction len", block.Transactions.Len())
}
