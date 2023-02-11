package bridge

import (
	"fmt"

	"github.com/polynetwork/bridge-common/tools"
)

type ListToken struct {
	Name string
	Symbol string
	Decimals int
	Address string
	ChainId string
	TokenType string
	DestChains map[string]map[string]struct {
		Name string
		Symbol string
		Decimals int
		Address string
		ChainId string
		TokenType string
		Router string
	}
}

type ListResponse map[string]*ListToken

func (c *Client) List(chainID uint64) (resp ListResponse, err error) {
	resp = make(ListResponse)
	err = tools.GetJsonFor(fmt.Sprintf("%s/tokenlistv4/%v", c.address, chainID), &resp)
	return
}
