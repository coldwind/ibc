package helper

import (
	"fmt"
	"ibc/account"
)

func GenKey() {
	fmt.Println("call GenKey")

	accountHandle := account.Account{}

	res, err := accountHandle.Generate()
	if err != nil {
		fmt.Println("gen key error", err)
		return
	}

	fmt.Println("PrivateKey:", res.PrivateKey)
	fmt.Println("PublicKey:", res.PublicKey)
	fmt.Println("Address:", res.Address)
}
