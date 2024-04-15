package account

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Key struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

func (h *Account) Generate() (*Key, error) {
	priKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	res := &Key{}

	priKeyBytes := crypto.FromECDSA(priKey)
	res.PrivateKey = hexutil.Encode(priKeyBytes)[2:]

	pubKey := priKey.Public()
	pubKeyECDSA := pubKey.(*ecdsa.PublicKey)
	pubKeyBytes := crypto.FromECDSAPub(pubKeyECDSA)
	res.PublicKey = hexutil.Encode(pubKeyBytes)[4:]
	res.Address = crypto.PubkeyToAddress(*pubKeyECDSA).Hex()

	return res, nil
}
