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

package poly

import (
	"encoding/binary"
	"time"

	"github.com/ontio/ontology/smartcontract/service/native/cross_chain/cross_chain_manager"
	"github.com/polynetwork/bridge-common/chains"
	"github.com/polynetwork/bridge-common/util"
	hcom "github.com/polynetwork/poly/native/service/header_sync/common"
	"github.com/polynetwork/poly/native/service/utils"

	"github.com/beego/beego/v2/core/logs"
	psdk "github.com/polynetwork/poly-go-sdk"
)

var (
	CCM_ADDRESS = utils.CrossChainManagerContractAddress.ToHexString()
)

type Rpc = psdk.PolySdk
type Client struct {
	*Rpc
	address string
}

func New(url string) *Client {
	s := psdk.NewPolySdk()
	s.NewRpcClient().SetAddress(url)
	return &Client{
		Rpc:     s,
		address: url,
	}
}

func (c *Client) Address() string {
	return c.address
}

func (c *Client) GetLatestHeight() (uint64, error) {
	h, err := c.GetCurrentBlockHeight()
	return uint64(h), err
}

func (c *Client) Confirm(hash string, blocks uint64, count int) (height uint64, err error) {
	var h, current uint32
	for count > 0 {
		count--
		h, err = c.GetBlockHeightByTxHash(hash)
		if err == nil {
			if blocks == 0 {
				return uint64(h), nil
			}
			current, err = c.GetCurrentBlockHeight()
			if err == nil && current >= h+uint32(blocks) {
				return uint64(h), nil
			}
		}
		if err != nil {
			logs.Error("Wait poly tx %s confirmation error %v", hash, err)
		}
		time.Sleep(time.Second)
	}
	return
}

func (c *Client) GetDoneTx(chainId uint64, ccId []byte) (data []byte, err error) {
	return c.GetStorage(utils.CrossChainManagerContractAddress.ToHexString(),
		append(append([]byte(cross_chain_manager.DONE_TX), utils.GetUint64Bytes(chainId)...), ccId...),
	)
}

func (c *Client) GetSideChainHeader(chainId uint64, height uint64) (hash []byte, err error) {
	return c.GetStorage(utils.HeaderSyncContractAddress.ToHexString(),
		append(append([]byte(hcom.MAIN_CHAIN), utils.GetUint64Bytes(chainId)...), utils.GetUint64Bytes(height)...),
	)
}

func (c *Client) GetSideChainHeight(chainId uint64) (height uint64, err error) {
	var id [8]byte
	binary.LittleEndian.PutUint64(id[:], chainId)
	res, err := c.GetStorage(utils.HeaderSyncContractAddress.ToHexString(), append([]byte(hcom.CURRENT_HEADER_HEIGHT), id[:]...))
	if err != nil {
		return 0, err
	}
	if res != nil && len(res) > 0 {
		height = binary.LittleEndian.Uint64(res)
	}
	return
}

func (c *Client) GetSideChainEpoch(chainId uint64) (data []byte, err error) {
	return c.GetStorage(utils.HeaderSyncContractAddress.ToHexString(),
		append([]byte(hcom.EPOCH_SWITCH), utils.GetUint64Bytes(chainId)...))
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
