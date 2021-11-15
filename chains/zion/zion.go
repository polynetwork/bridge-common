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
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	ccom "github.com/devfans/zion-sdk/contracts/native/cross_chain_manager/common"
	"github.com/devfans/zion-sdk/contracts/native/governance/node_manager"
	hcom "github.com/devfans/zion-sdk/contracts/native/header_sync/common"
	"github.com/devfans/zion-sdk/contracts/native/utils"
	"github.com/devfans/zion-sdk/core/state"

	// "github.com/polynetwork/bridge-common/base"
	"github.com/polynetwork/bridge-common/chains"
	"github.com/polynetwork/bridge-common/chains/eth"
	"github.com/polynetwork/bridge-common/log"
	"github.com/polynetwork/bridge-common/util"
	cstates "github.com/polynetwork/poly/core/states"
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

func New(url string) *Client {
	node_manager.InitABI()
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
	if err == nil {
		data, err = cstates.GetValueFromRawStorageItem(result)
	}
	return
}

func (c *Client) GetDoneTx(chainId uint64, ccId []byte) (data []byte, err error) {
	return c.GetStorage(utils.CrossChainManagerContractAddress,
		append(append([]byte(ccom.DONE_TX), utils.GetUint64Bytes(chainId)...), ccId...),
	)
}

func (c *Client) GetSideChainHeader(chainId uint64, height uint64) (hash []byte, err error) {
	return c.GetStorage(utils.HeaderSyncContractAddress,
		append(append([]byte(hcom.MAIN_CHAIN), utils.GetUint64Bytes(chainId)...), utils.GetUint64Bytes(height)...),
	)
}

func (c *Client) GetSideChainHeaderIndex(chainId uint64, height uint64) (hash []byte, err error) {
	return c.GetStorage(utils.HeaderSyncContractAddress,
		append(append([]byte(hcom.HEADER_INDEX), utils.GetUint64Bytes(chainId)...), utils.GetUint32Bytes(uint32(height))...),
	)
}

func (c *Client) GetSideChainConsensusBlockHeight(chainId uint64) (height uint64, err error) {
	key := util.Concat([]byte(hcom.CONSENSUS_PEER_BLOCK_HEIGHT), utils.GetUint64Bytes(chainId))
	res, err := c.GetStorage(utils.HeaderSyncContractAddress, key)
	height = utils.GetBytesUint64(res)
	return
}

func (c *Client) GetSideChainConsensusPeer(chainId uint64) (data []byte, err error) {
	key := util.Concat([]byte(hcom.CONSENSUS_PEER), utils.GetUint64Bytes(chainId))
	return c.GetStorage(utils.HeaderSyncContractAddress, key)
}

/*
func (c *Client) GetSideChainConsensusHeight(chainId uint64) (height uint64, err error) {
	var id [8]byte
	binary.LittleEndian.PutUint64(id[:], chainId)
	res, err := c.GetStorage(utils.HeaderSyncContractAddress, append([]byte(hcom.CONSENSUS_PEER), id[:]...))
	if err != nil {
		return
	}
	peer := new(neo.NeoConsensus)
	err = peer.Deserialization(common.NewZeroCopySource(res))
	if err != nil {
		return
	}
	height = uint64(peer.Height)
	return
}
*/

func (c *Client) GetSideChainMsg(chainId, height uint64) (msg []byte, err error) {
	return c.GetStorage(
		utils.HeaderSyncContractAddress,
		util.Concat([]byte(hcom.CROSS_CHAIN_MSG), utils.GetUint64Bytes(chainId), utils.GetUint32Bytes(uint32(height))),
	)
}

func (c *Client) GetSideChainMsgHeight(chainId uint64) (height uint64, err error) {
	res, err := c.GetStorage(utils.HeaderSyncContractAddress, append([]byte(hcom.CURRENT_MSG_HEIGHT), utils.GetUint64Bytes(chainId)...))
	if err != nil {
		return 0, err
	}
	if res != nil && len(res) > 0 {
		height = uint64(utils.GetBytesUint32(res))
	}
	return
}

func (c *Client) GetSideChainHeight(chainId uint64) (height uint64, err error) {
	/*
		if chainId == base.NEO {
			return c.GetSideChainConsensusHeight(chainId)
		}
	*/
	var id [8]byte
	binary.LittleEndian.PutUint64(id[:], chainId)
	heightBytes, err := c.GetStorage(utils.HeaderSyncContractAddress, append([]byte(hcom.CURRENT_HEADER_HEIGHT), id[:]...))
	if err != nil {
		return 0, err
	}
	if heightBytes == nil {
		return 0, fmt.Errorf("getPrevHeaderHeight, heightStore is nil")
	}
	/*
		heightBytes, err := cstates.GetValueFromRawStorageItem(res)
		if err != nil {
			return 0, fmt.Errorf("GetHeaderByHeight, deserialize headerBytes from raw storage item err:%v", err)
		}
	*/
	if heightBytes != nil {
		if len(heightBytes) > 7 {
			height = binary.LittleEndian.Uint64(heightBytes)
		} else if len(heightBytes) > 3 {
			height = uint64(binary.LittleEndian.Uint32(heightBytes))
		} else {
			err = fmt.Errorf("Failed to decode heightBytes, %v", heightBytes)
		}
	}
	return
}

func (c *Client) GetSideChainEpoch(chainId uint64) (data []byte, err error) {
	return c.GetStorage(utils.HeaderSyncContractAddress,
		append([]byte(hcom.EPOCH_SWITCH), utils.GetUint64Bytes(chainId)...))
}

func (c *Client) GetSideChainEpochWithHeight(chainId, height uint64) (data []byte, err error) {
	return c.GetStorage(utils.CrossChainManagerContractAddress,
		util.Concat([]byte(hcom.EPOCH_SWITCH), utils.GetUint64Bytes(chainId), utils.GetUint64Bytes(height)),
	)
}

func EpochProofKey(epochId uint64) common.Hash {
	enc := EpochProofDigest.Bytes()
	enc = append(enc, utils.GetUint64Bytes(epochId)...)
	proofHashKey := crypto.Keccak256Hash(enc)
	key := utils.ConcatKey(NODE_MANAGER_ADDRESS, []byte(SKP_PROOF), proofHashKey.Bytes())
	return crypto.Keccak256Hash(key[common.AddressLength:])
}

func (c *Client) GetEpochInfo(height uint64) (epochInfo *node_manager.EpochInfo, err error) {
	payload, err := new(node_manager.MethodEpochInput).Encode()
	if err != nil {
		return
	}
	arg := ethereum.CallMsg{
		From: common.Address{},
		To:   &NODE_MANAGER_ADDRESS,
		Data: payload,
	}
	res, err := c.CallContract(context.Background(), arg, big.NewInt(int64(height)))
	if err != nil {
		return
	}
	output := new(node_manager.MethodEpochOutput)
	if err = output.Decode(res); err != nil {
		return
	}
	epochInfo = output.Epoch
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
