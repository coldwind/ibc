package helper

import (
	"fmt"
	"ibc/example/config"
	"ibc/utils"
)

func IsContract() {
	fmt.Println("call IsContract")

	// get client
	cli, err := utils.GetClient(config.Rawurl)
	if err != nil {
		fmt.Println("get client error", err)
		return
	}

	is := cli.IsContract(config.FromAccount)

	fmt.Println("is contract:", is)
}
