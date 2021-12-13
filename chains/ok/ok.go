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

package ok

import (
	"math"
	"sync/atomic"
	"time"

	oksdk "github.com/okex/exchain-go-sdk"
	"github.com/polynetwork/bridge-common/chains"
	"github.com/polynetwork/bridge-common/util"
	"github.com/tendermint/tendermint/types"
)

type Rpc = oksdk.Client

type Client struct {
	*Rpc
	address string
	weight  *uint64
}

func New(url string) *Client {
	config, _ := oksdk.NewClientConfig(url, "okexchain-65", oksdk.BroadcastBlock, "0.01okt", 200000, 0, "")
	client := oksdk.NewClient(config)
	var max uint64 = math.MaxUint64
	return &Client{
		Rpc:     &client,
		address: url,
		weight:  &max,
	}
}

func (c *Client) Address() string {
	return c.address
}

func (c *Client) GetLatestHeight() (uint64, error) {
	return atomic.LoadUint64(c.weight), nil
}

func (c *Client) GetValidators(height uint64) (validators []*types.Validator, err error) {
	vr, err := c.Tendermint().QueryValidatorsResult(int64(height))
	if err != nil {
		if w := atomic.LoadUint64(c.weight); w > 0 {
			atomic.StoreUint64(c.weight, w-1)
		}
		return nil, err
	}
	return vr.Validators, nil
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
