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

package switcheo
/*
import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	"github.com/polynetwork/bridge-common/chains"
	"github.com/polynetwork/bridge-common/log"
	"github.com/polynetwork/bridge-common/util"
	"github.com/polynetwork/cosmos-poly-module/btcx"
	"github.com/polynetwork/cosmos-poly-module/ccm"
	"github.com/polynetwork/cosmos-poly-module/ft"
	"github.com/polynetwork/cosmos-poly-module/headersync"
	"github.com/polynetwork/cosmos-poly-module/lockproxy"
	tmclient "github.com/tendermint/tendermint/rpc/client/http"
)

type Rpc = tmclient.HTTP

type Client struct {
	*Rpc
	cdc     *codec.Codec
	address string
}

func New(url string) *Client {
	sdk := ontology_go_sdk.NewOntologySdk()
	sdk.NewRpcClient().SetAddress(url)

	config := types.GetConfig()
	config.SetBech32PrefixForAccount("swth", "swthpub")
	config.SetBech32PrefixForValidator("swthvaloper", "swthvaloperpub")
	config.SetBech32PrefixForConsensusNode("swthvalcons", "swthvalconspub")
	rawClient, err := tmclient.New(url, "/websocket")
	if err != nil {
		log.Error("Failed to connecto to cosmos ws", "err", err)
		return nil
	}
	cdc := codec.New()
	{
		bank.RegisterCodec(cdc)
		types.RegisterCodec(cdc)
		codec.RegisterCrypto(cdc)
		auth.RegisterCodec(cdc)
		btcx.RegisterCodec(cdc)
		ccm.RegisterCodec(cdc)
		ft.RegisterCodec(cdc)
		headersync.RegisterCodec(cdc)
		lockproxy.RegisterCodec(cdc)
	}
	return &Client{
		Rpc:     rawClient,
		cdc:     cdc,
		address: url,
	}
}

func (c *Client) Address() string {
	return c.address
}

func (c *Client) GetLatestHeight() (uint64, error) {
	status, err := c.Status()
	if err != nil {
		log.Error("Get cosmos current block status error", "err", err)
		return 0, err
	}
	res := status.SyncInfo.LatestBlockHeight
	return uint64(res), nil
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

func (s *SDK) Key() string {
	if s.ChainSDK != nil {
		return s.ChainSDK.Key()
	} else if s.options != nil {
		return s.options.Key()
	} else {
		panic("Unable to identify the sdk")
	}
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
 */
