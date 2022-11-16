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
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	erc20 "github.com/polynetwork/bridge-common/abi/mintable_erc20_abi"
	nftmapping "github.com/polynetwork/bridge-common/abi/nft_mapping_abi"
	"github.com/polynetwork/bridge-common/chains"
	"github.com/polynetwork/bridge-common/log"
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

func (c *Client) GetHeader(height uint64) (header Header, err error) {
	head := make(json.RawMessage, 0)
	err = c.Rpc.CallContext(context.Background(), &head, "eth_getBlockByNumber", util.BlockNumArg(height), false)
	if err == nil && len(head) == 0 {
		err = ethereum.NotFound
	}
	if err != nil {
		return
	}
	header = Header(head)
	if height == 0 {
		_, err = header.GetHeight()
	} else {
		err = header.Verify(height)
	}
	return
}

func (c *Client) GetProof(addr string, key string, height uint64) (proof *ETHProof, err error) {
	heightHex := "latest"
	if height > 0 {
		heightHex = hexutil.EncodeBig(big.NewInt(int64(height)))
	}
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
		if err != nil {
			log.Info("Wait poly tx confirmation error", "count", count, "hash", hash, "err", err)
		}
		time.Sleep(time.Second * 3)
	}
	return
}

func (c *Client) GetBalance(token, owner string) (balance *big.Int, err error) {
	tokenAddress := common.HexToAddress(token)
	ownerAddr := common.HexToAddress(owner)
	if token == "0000000000000000000000000000000000000000" {
		var result hexutil.Big
		err = c.Rpc.CallContext(context.Background(), &result, "eth_getBalance", "0x"+owner, "latest")
		balance = (*big.Int)(&result)
	} else {
		var contract *erc20.ERC20Extended
		contract, err = erc20.NewERC20Mintable(tokenAddress, c)
		if err == nil {
			balance, err = contract.BalanceOf(nil, ownerAddr)
		}
	}
	return
}

func (c *Client) GetNFTOwner(asset string, tokenId int64) (owner common.Address, err error) {
	cm, err := nftmapping.NewCrossChainNFTMapping(common.HexToAddress(asset), c)
	if err != nil {
		return common.Address{}, err
	}
	return cm.OwnerOf(nil, big.NewInt(tokenId))
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
