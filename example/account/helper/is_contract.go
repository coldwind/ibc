package helper

import (
	"fmt"
	"ibc/utils"
)

func IsContract() {
	fmt.Println("call IsContract")

	// get client
	cli, err := utils.GetClient("https://cloudflare-eth.com")
	if err != nil {
		fmt.Println("get client error", err)
		return
	}

	is := cli.IsContract("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	fmt.Println("is contract:", is)
}
