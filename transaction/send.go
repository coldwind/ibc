package transaction

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SendEth(cli *ethclient.Client, privateKey string, to string, eth *big.Int) (*types.Transaction, error) {
	chainId, err := cli.NetworkID(context.Background())
	if err != nil {
		fmt.Println("get chainId error", err)
		return nil, err
	}

	priKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println("load private key error", err)
		return nil, err
	}

	pubKey := priKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("load public key error", err)
		return nil, errors.New("load public key error")
	}

	fromAddress := crypto.PubkeyToAddress(*pubKeyECDSA)

	// nonce
	nonce, err := cli.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("get nonce error", err)
		return nil, err
	}

	// gas
	gasLimit := uint64(21000)
	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("get gas price error", err)
		return nil, err
	}

	// to
	toAddress := common.HexToAddress(to)

	// 生成未签名事务
	tx := types.NewTransaction(nonce, toAddress, eth, gasLimit, gasPrice, nil)

	// 签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), priKey)
	if err != nil {
		fmt.Println("signed error", err)
		return nil, err
	}

	err = cli.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("send transaction", err)
		return nil, err
	}

	return signedTx, nil
}
