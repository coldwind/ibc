package utils

import (
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	cli *ethclient.Client
}

var clientPool sync.Map

func GetClient(rawurl string) (*Client, error) {
	if val, ok := clientPool.Load(rawurl); ok {
		return val.(*Client), nil
	} else {
		if client, err := ethclient.Dial(rawurl); err == nil {
			res := &Client{
				cli: client,
			}
			clientPool.Store(rawurl, res)
			return res, nil
		} else {
			return nil, err
		}
	}
}

func (h *Client) Cli() *ethclient.Client {
	return h.cli
}
