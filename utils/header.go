package utils

import "context"

func (h *Client) GetHeader() {
	header, err := h.cli.HeaderByNumber(context.Background(), nil)
}
