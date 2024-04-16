package transaction

import (
	"context"
	"fmt"
	"ibc/example/config"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"golang.org/x/crypto/sha3"
)

func TransferERC20(cli *ethclient.Client, hexContractAddress string, hexFromAddress string, hexToAddress string, amount *big.Int) {
	value := big.NewInt(0)

	contractAddress := common.HexToAddress(hexContractAddress)
	fromAddress := common.HexToAddress(hexFromAddress)
	toAddress := common.HexToAddress(hexToAddress)

	// transfer sign
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println("methodID", hexutil.Encode(methodID))

	// 填充字节
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println("paddedAddress", hexutil.Encode(paddedAddress))

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	// 合并成data
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// 获取gas相关数据
	gasLimit, err := cli.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})

	if err != nil {
		fmt.Println("gas limit get error", err)
		return
	}
	fmt.Println("gas limit", gasLimit)

	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("gas price get error", err)
		return
	}
	fmt.Println("gas price", gasPrice)

	nonce, _ := cli.NonceAt(context.Background(), fromAddress, nil)

	// 生成交易事务
	tx := types.NewTransaction(nonce, contractAddress, value, gasLimit, gasPrice, data)

	// 签名
	chainId, err := cli.NetworkID(context.Background())
	if err != nil {
		fmt.Println("chainId error", err)
		return
	}

	// 私钥
	priKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		fmt.Println("PrivateKey error", err)
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), priKey)
	if err != nil {
		fmt.Println("SignTx error", err)
		return
	}

	err = cli.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("SendTransaction error", err)
		return
	}
}
