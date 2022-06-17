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

package zion

import (
	"context"
	"encoding/hex"
	"math/big"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/devfans/zion-sdk/contracts/native/governance/node_manager"
	"github.com/devfans/zion-sdk/contracts/native/utils"
	"github.com/devfans/zion-sdk/core/state"
	"github.com/devfans/zion-sdk/contracts/native/go_abi/node_manager_abi"

	"github.com/polynetwork/bridge-common/chains"
	"github.com/polynetwork/bridge-common/chains/eth"
	"github.com/polynetwork/bridge-common/log"
	"github.com/polynetwork/bridge-common/util"
)

var (
	CCM_ADDRESS          = utils.CrossChainManagerContractAddress
	NODE_MANAGER_ADDRESS = utils.NodeManagerContractAddress
	SKP_PROOF            = "st_proof"
	EpochProofDigest     = common.HexToHash("e4bf3526f07c80af3a5de1411dd34471c71bdd5d04eedbfa1040da2c96802041")

	_ZION_ID uint64
)

type Client struct {
	eth.Client
}

func ReadChainID() uint64 {
	return atomic.LoadUint64(&_ZION_ID)
}

func init() {
	node_manager.InitABI()
}

func New(url string) *Client {
	c := eth.New(url)
	if c == nil {
		return nil
	}
	id, err := c.NetworkID(context.Background())
	if err != nil {
		log.Error("Failed to get network id", "url", url, "err", err)
		return nil
	}
	atomic.StoreUint64(&_ZION_ID, id.Uint64())
	return &Client{*c}
}

func (c *Client) GetLatestHeight() (uint64, error) {
	var result hexutil.Big
	err := c.Rpc.CallContext(context.Background(), &result, "eth_blockNumber")
	for err != nil {
		return 0, err
	}
	return (*big.Int)(&result).Uint64(), err
}

func (c *Client) GetBlockHeightByTxHash(hash common.Hash) (uint64, error) {
	receipt, err := c.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return 0, err
	}
	return receipt.BlockNumber.Uint64(), nil
}

func (c *Client) ConfirmWait(hash common.Hash, blocks uint64, count int) (height uint64, err error) {
	var h, current uint64
	for count > 0 {
		count--
		h, err = c.GetBlockHeightByTxHash(hash)
		if err == nil {
			if blocks == 0 {
				return h, nil
			}
			current, err = c.GetLatestHeight()
			if err == nil && current >= h+blocks {
				return h, nil
			}
		}
		// TODO: Check confirm here
		if err != nil && !strings.Contains(err.Error(), "not found") {
			log.Info("Wait poly tx confirmation error", "count", count, "hash", hash, "err", err)
		}
		time.Sleep(time.Second)
	}
	return
}

func (c *Client) GetStorageAt(contract common.Address, key []byte) (data []byte, err error) {
	return c.StorageAt(context.Background(), contract, state.Key2Slot(key), nil)
}

func (c *Client) GetStorage(account common.Address, key []byte) (data []byte, err error) {
	var result hexutil.Bytes
	keyHex := hex.EncodeToString(key)
	err = c.Rpc.CallContext(context.Background(), &result, "eth_getStorageAtCacheDB", account, keyHex, "latest")
	return result, err
}

func EpochProofKey(epochId uint64) common.Hash {
	enc := EpochProofDigest.Bytes()
	enc = append(enc, utils.GetUint64Bytes(epochId)...)
	proofHashKey := crypto.Keccak256Hash(enc)
	key := utils.ConcatKey(NODE_MANAGER_ADDRESS, []byte(SKP_PROOF), proofHashKey.Bytes())
	return crypto.Keccak256Hash(key[common.AddressLength:])
}

func (c *Client) GetEpochInfo(height uint64) (epochInfo *node_manager.EpochInfo, err error) {
	abi, err := node_manager_abi.NewINodeManager(utils.NodeManagerContractAddress, c)
	if err != nil { return }
	var options *bind.CallOpts
	if height > 0 {
		options = &bind.CallOpts{BlockNumber: big.NewInt(int64(height))}
	}
	data, err := abi.Epoch(options)
	if err != nil { return }
	info := new(node_manager.EpochInfo)
	err = rlp.DecodeBytes(data, info)
	if err != nil { return }
	return epochInfo, err
}

func (c *Client) EpochById(epochId uint64) (epochInfo *node_manager.EpochInfo, err error) {
	abi, err := node_manager_abi.NewINodeManager(utils.NodeManagerContractAddress, c)
	if err != nil { return }
	data, err := abi.GetEpochByID(nil, epochId)
	if err != nil { return }
	if len(data) == 0 {
		return
	}
	info := new(node_manager.EpochInfo)
	err = rlp.DecodeBytes(data, info)
	if err != nil { return }
	return epochInfo, err
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
	return &SDK{ChainSDK: sdk, nodes: clients}, err
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
