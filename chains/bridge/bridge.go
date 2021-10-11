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

package bridge

import (
	"time"

	"github.com/polynetwork/bridge-common/chains"
)

type Client struct {
	address string
}

func New(url string) *Client {
	return &Client{
		address: url,
	}
}

func (c *Client) Address() string {
	return c.address
}

func (c *Client) GetLatestHeight() (uint64, error) {
	return 1, nil
}

type SDK struct {
	*chains.ChainSDK
	nodes []*Client
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
