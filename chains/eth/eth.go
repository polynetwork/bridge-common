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
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/polynetwork/bridge-common/chains"
	"github.com/polynetwork/bridge-common/log"
	"github.com/polynetwork/bridge-common/util"
)

type Client struct {
	Rpc *rpc.Client
	*ethclient.Client
	address string

	ws *ethclient.Client
	sync.RWMutex
	subsOn bool
	subs []chan<-uint64
	url string
	index int
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

func (c *Client) WsClient() (r *ethclient.Client, a string) {
	c.RLock()
	r = c.ws
	a = c.url
	c.RUnlock()
	return
}

// Update ws client
func (c *Client) Connect(url string) (err error) {
	if !strings.HasPrefix(url, "ws") {
		return fmt.Errorf("invalid ws url, %s", url)
	}

	r, a := c.WsClient()
	if a == url {
		log.Info("ws was already connected", "url", url)
		return
	} else if a != "" {
		log.Warn("ws client will be updated", "url", a, "new_url", url)
	}

	ws, err := ethclient.Dial(url)
	if err != nil {
		return fmt.Errorf("ws connect failure, err %v", err)
	}

	if r != nil {
		r.Close()
	}

	c.Lock()
	defer c.Unlock()
	c.url = url
	c.ws = ws
	return
}


func (c *Client) dispatch(ch <-chan *types.Header) {
	for header := range ch {
		height := header.Number.Uint64()
		c.RLock()
		subs := c.subs
		c.RUnlock()
		for _, sink := range subs {
			select {
			case sink <- height:
			default:
			}
		}
	}
}

func (c *Client) listen(ch chan <- *types.Header) {
	for {
		ws, url := c.WsClient()
		if ws == nil {
			log.Warn("ws client is not conneced")
			time.Sleep(time.Second * 10)
			continue
		}
		log.Info("Subscribing eth chain update", "url", url)
		sub, err := ws.SubscribeNewHead(context.Background(), ch)
		if err != nil {
			log.Error("ws subscribe new head failure", "url", url, "err", err)
		} else {
			err = <- sub.Err()
			log.Error("ws subscription closed", "url", url, "err", err)
			sub.Unsubscribe()
		}
		time.Sleep(time.Second * 2)
	}
}

func (c *Client) Listen(url string) (err error) {
	err = c.Connect(url)
	if err != nil { return }
	c.Lock()
	if c.subsOn {
		err = fmt.Errorf("ws client was already listening, url %s new url %s", c.url, url)
	} else {
		c.subsOn = true
	}
	c.Unlock()
	if err != nil { return }
	ch := make(chan *types.Header, 10)
	go c.dispatch(ch)
	go c.listen(ch)
	return
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

func(c *Client) Subscribe(ch chan <-uint64) {
	c.Lock()
	has := false
	for _, sub := range c.subs {
		if sub == ch {
			has = true
			break
		}
	}
	if !has {
		c.subs = append(c.subs, ch)
	}
	c.Unlock()
}

func(c *Client) Unsubscribe(ch chan <-uint64) {
	c.Lock()
	defer c.Unlock()
	index := len(c.subs)
	for i, sub := range c.subs {
		if sub == ch {
			index = i
			break
		}
	}
	if index == len(c.subs) {
		return
	}
	start := 0
	end := len(c.subs) - 1
	if index == start {
		start = 1
	} else if index < end {
		copy(c.subs[index: len(c.subs)-1], c.subs[index+1: len(c.subs)])
	}
	c.subs = c.subs[start: end]
}

func (c *Client) GetLatestHeight() (uint64, error) {
	var result hexutil.Big
	err := c.Rpc.CallContext(context.Background(), &result, "eth_blockNumber")
	for err != nil {
		return 0, err
	}
	return (*big.Int)(&result).Uint64(), err
}

// TransactionByHash returns the transaction with the given hash.
func (c *Client) TransactionWithExtraByHash(ctx context.Context, hash common.Hash) (json *rpcTransaction, err error) {
	err = c.Rpc.CallContext(ctx, &json, "eth_getTransactionByHash", hash)
	return
	/*
		if err != nil {
			return nil, err
		} else if json == nil || json.tx == nil {
			return nil, nil
		} else if _, r, _ := json.tx.RawSignatureValues(); r == nil {
			return nil, fmt.Errorf("server returned transaction without signature")
		}
		if json.From != nil && json.BlockHash != nil {
			setSenderFromServer(json.tx, *json.From, *json.BlockHash)
		}
		return json, nil
	*/
}

func (c *Client) GetTxHeight(ctx context.Context, hash common.Hash) (height uint64, pending bool, err error) {
	tx, err := c.TransactionWithExtraByHash(context.Background(), hash)
	if err != nil || tx == nil {
		return
	}
	pending = tx.BlockNumber == nil
	if !pending {
		v := big.NewInt(0)
		v.SetString((*tx.BlockNumber)[2:], 16)
		height = v.Uint64()
	}
	return
}

func (c *Client) Confirm(hash common.Hash, blocks uint64, count int) (height, confirms uint64, pending bool, err error) {
	var current uint64
	for count > 0 {
		count--
		confirms = 0
		height, pending, err = c.GetTxHeight(context.Background(), hash)
		if height > 0 {
			if blocks == 0 {
				return
			}
			current, err = c.GetLatestHeight()
			if current >= height {
				confirms = current - height
				if confirms >= blocks {
					return
				}
			}
		}
		log.Info("Wait tx confirmation", "count", count, "hash", hash, "height", height, "latest", current, "pending", pending, "err", err)
		time.Sleep(time.Second)
	}
	return
}

func (c *Client) Index() int {
	return c.index
}

type SDK struct {
	*chains.ChainSDK
	nodes   []*Client
	options *chains.Options
}

func (s *SDK) Node() *Client {
	return s.nodes[s.ChainSDK.Index()]
}

func (s *SDK) Broadcast(tx *types.Transaction) (err error) {
	nodes := s.Nodes()
	for _, idx := range nodes[1:] {
		go func(id int) {
			_ = s.nodes[id].SendTransaction(context.Background(), tx)
		} (idx)
	}
	return s.nodes[nodes[0]].SendTransaction(context.Background(), tx)
}

func (s *SDK) Select() *Client {
	return s.nodes[s.ChainSDK.Select()]
}

func NewSDK(chainID uint64, urls []string, interval time.Duration, maxGap uint64) (*SDK, error) {

	clients := make([]*Client, len(urls))
	nodes := make([]chains.SDK, len(urls))
	for i, url := range urls {
		client := New(url)
		client.index = i
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
