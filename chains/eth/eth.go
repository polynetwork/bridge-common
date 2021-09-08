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

package eth

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/polynetwork/bridge-common/chains"
	"github.com/polynetwork/bridge-common/util"
)

type Client struct {
	Rpc *rpc.Client
	*ethclient.Client
	address string
}

func New(url string) *Client {
	rpcClient, _ := rpc.Dial(url)
	rawClient, _ := ethclient.Dial(url)
	return &Client{
		Rpc:     rpcClient,
		Client:  rawClient,
		address: url,
	}
}

func (c *Client) Address() string {
	return c.address
}

func (c *Client) GetProof(addr string, key string, height uint64) (proof *ETHProof, err error) {
	heightHex := hexutil.EncodeBig(big.NewInt(int64(height)))
	proof = &ETHProof{}
	err = c.Rpc.CallContext(context.Background(), &proof, "eth_getProof", addr, []string{key}, heightHex)
	return
}

func (c *Client) GetLatestHeight() (uint64, error) {
	var result hexutil.Big
	err := c.Rpc.CallContext(context.Background(), &result, "eth_blockNumber")
	for err != nil {
		return 0, err
	}
	return (*big.Int)(&result).Uint64(), err
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

func WithOptions(chainID uint64, urls []string, interval time.Duration, maxGap uint64) *SDK {
	return util.Single(&SDK{
		options: &chains.Options{
			ChainID:  chainID,
			Nodes:    urls,
			Interval: interval,
			MaxGap:   maxGap,
		},
	}).(*SDK)
}

func (s *SDK) Create() interface{} {
	sdk, _ := NewSDK(s.options.ChainID, s.options.Nodes, s.options.Interval, s.options.MaxGap)
	return sdk
}

func (s *SDK) Key() string {
	if s.ChainSDK == nil {
		return s.ChainSDK.Key()
	} else if s.options != nil {
		return s.options.Key()
	} else {
		panic("Unable to identify the sdk")
	}
}
