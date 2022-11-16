/*
 * Copyright (C) 2021 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package neo

import (
	"fmt"
	"github.com/joeqian10/neo-gogogo/nep5"
	"math/big"
	"time"

	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/rpc/models"

	"github.com/polynetwork/bridge-common/chains"
	"github.com/polynetwork/bridge-common/util"
)

const (
	GET_CURRENT_HEIGHT = "currentSyncHeight"
)

type Rpc = rpc.RpcClient

type Client struct {
	*Rpc
	address string
}

func New(url string) *Client {
	client := rpc.NewClient(url)
	return &Client{
		Rpc:     client,
		address: url,
	}
}

func (c *Client) Address() string {
	return c.address
}

func (c *Client) GetLatestHeight() (uint64, error) {
	res := c.GetBlockCount()
	if res.ErrorResponse.Error.Message != "" {
		return 0, fmt.Errorf("%s", res.ErrorResponse.Error.Message)
	}
	return uint64(res.Result), nil
}

func (c *Client) GetPolyEpochHeight(ccm string, chainId uint64) (height uint64, err error) {
	arg := models.InvokeStack{Type: "Integer", Value: chainId}
	response := c.InvokeFunction("0x"+helper.ReverseString(ccm), GET_CURRENT_HEIGHT, helper.ZeroScriptHashString, arg)
	if response.HasError() || response.Result.State == "FAULT" {
		return 0, fmt.Errorf("[GetCurrentNeoChainSyncHeight] GetCurrentHeight error: %s", "Engine faulted! "+response.Error.Message)
	}

	var b []byte
	s := response.Result.Stack
	if s != nil && len(s) != 0 {
		s[0].Convert()
		b = helper.HexToBytes(s[0].Value.(string))
	}
	if len(b) > 0 {
		height = helper.BytesToUInt64(b)
		height++ // means the next block header needs to be synced
	}
	return
}

func (c *Client) GetBalance(token, owner string) (balance *big.Int, err error) {
	scriptHash, err := helper.UInt160FromString(token)
	if err != nil {
		return nil, err
	}
	nep5Helper := nep5.NewNep5Helper(scriptHash, c.address)
	addrHash, err := helper.UInt160FromString(owner)
	if err != nil {
		return nil, err
	}
	amount, err := nep5Helper.BalanceOf(addrHash)
	return new(big.Int).SetUint64(amount), nil
}

type SDK struct {
	*chains.ChainSDK
	nodes   []*Client
	options *chains.Options
}

func (s *SDK) Node() *Client {
	return s.nodes[s.ChainSDK.Index()]
}

func (s *SDK) Select() *Client {
	return s.nodes[s.ChainSDK.Select()]
}

func NewSDK(chainID uint64, urls []string, interval time.Duration, maxGap uint64) (*SDK, error) {

	clients := make([]*Client, len(urls))
	nodes := make([]chains.SDK, len(urls))
	for i, url := range urls {
		client := New(url)
		nodes[i] = client
		clients[i] = client
	}
	sdk, err := chains.NewChainSDK(chainID, nodes, interval, maxGap)
	if err != nil {
		return nil, err
	}
	return &SDK{ChainSDK: sdk, nodes: clients}, nil
}

func WithOptions(chainID uint64, urls []string, interval time.Duration, maxGap uint64) (*SDK, error) {
	sdk, err := util.Single(&SDK{
		options: &chains.Options{
			ChainID:  chainID,
			Nodes:    urls,
			Interval: interval,
			MaxGap:   maxGap,
		},
	})
	if err != nil {
		return nil, err
	}
	return sdk.(*SDK), nil
}

func (s *SDK) Create() (interface{}, error) {
	return NewSDK(s.options.ChainID, s.options.Nodes, s.options.Interval, s.options.MaxGap)
}

func (s *SDK) Key() string {
	if s.ChainSDK != nil {
		return s.ChainSDK.Key()
	} else if s.options != nil {
		return s.options.Key()
	} else {
		panic("Unable to identify the sdk")
	}
}
