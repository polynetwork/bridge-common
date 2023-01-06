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
	"fmt"
	"sort"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ontio/ontology/smartcontract/service/native/cross_chain/cross_chain_manager"

	"github.com/polynetwork/bridge-common/base"
	"github.com/polynetwork/bridge-common/chains"
	"github.com/polynetwork/bridge-common/log"
	"github.com/polynetwork/bridge-common/util"
	psdk "github.com/polynetwork/poly-go-sdk"
	"github.com/polynetwork/poly/common"
	hcom "github.com/polynetwork/poly/native/service/header_sync/common"
	"github.com/polynetwork/poly/native/service/header_sync/neo"
	"github.com/polynetwork/poly/native/service/utils"
)

var (
	CCM_ADDRESS                        = utils.CrossChainManagerContractAddress.ToHexString()
	signatureManagerContractAddress, _ = common.AddressParseFromBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08})
	SM_ADDRESS                         = signatureManagerContractAddress.ToHexString()
	_POLY_ID                           uint64
)

type Rpc = psdk.PolySdk
type Client struct {
	*Rpc
	address string
}

func ReadChainID() uint64 {
	return atomic.LoadUint64(&_POLY_ID)
}

func New(url string) *Client {
	s := psdk.NewPolySdk()
	s.NewRpcClient().SetAddress(url)
	hdr, err := s.GetHeaderByHeight(0)
	if err != nil {
		log.Error("Failed to initialize poly sdk", "err", err)
		return nil
	}
	atomic.StoreUint64(&_POLY_ID, hdr.ChainID)
	s.SetChainId(hdr.ChainID)
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
		if err != nil && !strings.Contains(err.Error(), "INVALID PARAMS") {
			log.Info("Wait poly tx confirmation error", "count", count, "hash", hash, "err", err)
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

func (c *Client) GetSideChainHeaderIndex(chainId uint64, height uint64) (hash []byte, err error) {
	return c.GetStorage(utils.HeaderSyncContractAddress.ToHexString(),
		append(append([]byte(hcom.HEADER_INDEX), utils.GetUint64Bytes(chainId)...), utils.GetUint32Bytes(uint32(height))...),
	)
}

func (c *Client) GetSideChainConsensusBlockHeight(chainId uint64) (height uint64, err error) {
	key := util.Concat([]byte(hcom.CONSENSUS_PEER_BLOCK_HEIGHT), utils.GetUint64Bytes(chainId))
	res, err := c.GetStorage(utils.HeaderSyncContractAddress.ToHexString(), key)
	height = utils.GetBytesUint64(res)
	return
}

func (c *Client) GetSideChainConsensusPeer(chainId uint64) (data []byte, err error) {
	key := util.Concat([]byte(hcom.CONSENSUS_PEER), utils.GetUint64Bytes(chainId))
	return c.GetStorage(utils.HeaderSyncContractAddress.ToHexString(), key)
}

func (c *Client) GetSideChainConsensusHeight(chainId uint64) (height uint64, err error) {
	var id [8]byte
	binary.LittleEndian.PutUint64(id[:], chainId)
	res, err := c.GetStorage(utils.HeaderSyncContractAddress.ToHexString(), append([]byte(hcom.CONSENSUS_PEER), id[:]...))
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

func (c *Client) GetSideChainMsg(chainId, height uint64) (msg []byte, err error) {
	return c.GetStorage(
		utils.HeaderSyncContractAddress.ToHexString(),
		util.Concat([]byte(hcom.CROSS_CHAIN_MSG), utils.GetUint64Bytes(chainId), utils.GetUint32Bytes(uint32(height))),
	)
}

func (c *Client) GetSideChainMsgHeight(chainId uint64) (height uint64, err error) {
	res, err := c.GetStorage(utils.HeaderSyncContractAddress.ToHexString(), append([]byte(hcom.CURRENT_MSG_HEIGHT), utils.GetUint64Bytes(chainId)...))
	if err != nil {
		return 0, err
	}
	if res != nil && len(res) > 0 {
		height = uint64(utils.GetBytesUint32(res))
	}
	return
}

func (c *Client) GetSideChainHeight(chainId uint64) (height uint64, err error) {
	if chainId == base.NEO {
		return c.GetSideChainConsensusHeight(chainId)
	}
	var id [8]byte
	binary.LittleEndian.PutUint64(id[:], chainId)
	res, err := c.GetStorage(utils.HeaderSyncContractAddress.ToHexString(), append([]byte(hcom.CURRENT_HEADER_HEIGHT), id[:]...))
	if err != nil {
		return 0, err
	}
	if res != nil {
		if len(res) > 7 {
			height = binary.LittleEndian.Uint64(res)
		} else {
			height = uint64(binary.LittleEndian.Uint32(res))
		}
	}
	return
}

func (c *Client) GetSideChainEpoch(chainId uint64) (data []byte, err error) {
	return c.GetStorage(utils.HeaderSyncContractAddress.ToHexString(),
		append([]byte(hcom.EPOCH_SWITCH), utils.GetUint64Bytes(chainId)...))
}

func (c *Client) GetSideChainEpochWithHeight(chainId, height uint64) (data []byte, err error) {
	return c.GetStorage(utils.CrossChainManagerContractAddress.ToHexString(),
		util.Concat([]byte(hcom.EPOCH_SWITCH), utils.GetUint64Bytes(chainId), utils.GetUint64Bytes(height)),
	)
}

type SDK struct {
	*chains.ChainSDK
	nodes   []*Client
	options *chains.Options
}

func (s *SDK) Node() *Client {
	return s.nodes[s.ChainSDK.Index()]
}

func (s *SDK) AllNodes() []*Client {
	return s.nodes
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

type SigInfo struct {
	Status  bool
	SigInfo map[string][]byte
}

func (s *SigInfo) Serialization(sink *common.ZeroCopySink) {
	sink.WriteBool(s.Status)
	sink.WriteUint64(uint64(len(s.SigInfo)))
	sigInfoList := make([]string, 0, len(s.SigInfo))
	for k := range s.SigInfo {
		sigInfoList = append(sigInfoList, k)
	}
	sort.SliceStable(sigInfoList, func(i, j int) bool {
		return sigInfoList[i] > sigInfoList[j]
	})
	for _, k := range sigInfoList {
		sink.WriteString(k)
		v := s.SigInfo[k]
		sink.WriteVarBytes(v)
	}
}

func (s *SigInfo) Deserialization(source *common.ZeroCopySource) error {
	status, eof := source.NextBool()
	if eof {
		return fmt.Errorf("SigInfo deserialize status length error")
	}
	n, eof := source.NextUint64()
	if eof {
		return fmt.Errorf("SigInfo deserialize SigInfo length error")
	}
	sigInfo := make(map[string][]byte)
	for i := 0; uint64(i) < n; i++ {
		k, e := source.NextString()
		if e {
			return fmt.Errorf("SigInfo deserialize key error")
		}
		v, e := source.NextVarBytes()
		if e {
			return fmt.Errorf("SigInfo deserialize value error")
		}
		sigInfo[k] = v
	}
	s.Status = status
	s.SigInfo = sigInfo
	return nil
}
