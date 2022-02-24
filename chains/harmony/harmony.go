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

package harmony

import (
	"context"
	"math/big"
	"time"

	"github.com/polynetwork/bridge-common/chains"
	"github.com/polynetwork/bridge-common/chains/eth"
	"github.com/polynetwork/bridge-common/util"
	"github.com/polynetwork/bridge-common/chains/custom"
)

type Client struct {
	*eth.Client
	caller *custom.Client
}

func New(url string) *Client {
	client, _ := custom.Dial(url)
	return &Client{
		eth.New(url), client,
	}
}

func(c *Client) HeaderByNumber(height uint64) (data []byte, err error) {
	resp, err := c.caller.CallContextRaw(
		context.Background(), "eth_getBlockByNumber",
		custom.ToBlockNumArg(big.NewInt(int64(height))), false)

	return resp.Result, err
}

func(c *Client) HeaderByNumberRLP(height uint64) (data []byte, err error) {
	resp, err := c.caller.CallContextRaw(
		context.Background(), "hmy_getHeaderByNumberRLPHex",
		custom.ToBlockNumArg(big.NewInt(int64(height))), false)

	if err == nil {
		data = resp.Result
	}
	return
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
