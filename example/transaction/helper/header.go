package helper

import (
	"fmt"
	"ibc/utils"
)

func HeaderByNumber() {
	// get client
	cli, err := utils.GetClient("https://cloudflare-eth.com")
	if err != nil {
		fmt.Println("get client error", err)
		return
	}

	head, err := cli.GetHeaderByNumber(nil)
	if err != nil {
		fmt.Println("get header error", err)
		return
	}
	fmt.Println("number", head.Number)
	fmt.Println("difficulty", head.Difficulty)
	fmt.Println("parentHash", head.ParentHash)
	fmt.Println("hash", head.Hash)
	fmt.Println("time", head.Time)
}
