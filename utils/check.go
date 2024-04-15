package utils

import (
	"context"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

func (h *Client) IsValid(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

func (h *Client) IsContract(hexAddress string) bool {
	if !h.IsValid(hexAddress) {
		return false
	}

	address := common.HexToAddress(hexAddress)
	code, err := h.Cli().CodeAt(context.Background(), address, nil)
	if err != nil {
		fmt.Println("is contract error", err)
		return false
	}

	return len(code) > 0
}
