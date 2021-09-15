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

package matic

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	abcitypes "github.com/christianxiao/tendermint/abci/types"
	amino "github.com/christianxiao/tendermint/crypto/encoding/amino"
	"github.com/christianxiao/tendermint/crypto/merkle"
	"github.com/christianxiao/tendermint/rpc/client"
	tmtypes "github.com/christianxiao/tendermint/types"
	"github.com/polynetwork/bridge-common/chains"
	"github.com/polynetwork/bridge-common/chains/matic/cosmos"
	htypes "github.com/polynetwork/bridge-common/chains/matic/heimdall/types"
	"github.com/polynetwork/bridge-common/log"
	"github.com/polynetwork/bridge-common/util"
)

const (
	SPRINT_SIZE uint64 = 64
)

type Rpc = client.HTTP

type Client struct {
	*Rpc
	address string
	codec   *cosmos.Codec
}

// GetSpanKey appends prefix to start block
func GetSpanKey(id uint64) []byte {
	return append([]byte{0x36}, []byte(strconv.FormatUint(id, 10))...)
}

func New(url string) *Client {
	codec := cosmos.New()
	amino.RegisterAmino(codec)
	// heimdall client
	var tclient *client.HTTP
	if strings.HasPrefix(url, "http://") {
		tclient = client.NewHTTP(url, "/websocket")
	} else {
		cc := &http.Client{ // https
			Transport: &http.Transport{
				// Set to true to prevent GZIP-bomb DoS attacks
				DisableCompression: true,
				TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			},
		}
		tclient = client.NewHTTPWithClient(url, "/websocket", cc)
	}
	err := tclient.Start()
	if err != nil {
		log.Error("Failed to start heimdall client", "err", err)
		return nil
	}
	return &Client{
		address: url,
		codec:   codec,
		Rpc:     tclient,
	}
}

func (c *Client) Address() string {
	return c.address
}

func (c *Client) MarshalHeader(header cosmos.CosmosHeader) (data []byte, err error) {
	data, err = c.codec.MarshalBinaryBare(header)
	return
}

func (c *Client) ComposeHeaderProof(height, hmHeight, spanId uint64, hp *cosmos.HeaderWithOptionalProof) (err error) {
	_, err = c.GetLatestSpan(hmHeight)
	if err != nil {
		return err
	}

	spanRes, _, err := c.GetSpanRes(spanId, hmHeight-1)
	if err != nil {
		return err
	}

	cosmosHeader, err := c.GetCosmosHeader(hmHeight)
	if err != nil {
		return err
	}

	kp := merkle.KeyPath{}
	kp = kp.AppendKey([]byte("bor"), merkle.KeyEncodingURL)
	kp = kp.AppendKey(spanRes.Key, merkle.KeyEncodingURL)
	value := cosmos.CosmosProofValue{
		Kp:    kp.String(),
		Value: spanRes.GetValue(),
	}
	proof := &cosmos.CosmosProof{
		Value:  value,
		Proof:  *spanRes.Proof,
		Header: *cosmosHeader,
	}
	hp.Proof, err = c.codec.MarshalBinaryBare(proof)
	return
}

func (c *Client) GetSpanRes(id uint64, height uint64) (*abcitypes.ResponseQuery, *htypes.Span, error) {
	res, err := c.ABCIQueryWithOptions(
		"/store/bor/key",
		GetSpanKey(id),
		client.ABCIQueryOptions{Prove: true, Height: int64(height)})
	if err != nil {
		return nil, nil, err
	}
	if len(res.Response.Value) == 0 || len(res.Response.Proof.GetOps()) == 0 || len(res.Response.Key) == 0 {
		return nil, nil, fmt.Errorf("Heimdall height (%d), Span(%d) too new?", height, id)
	}
	span := new(htypes.Span)
	err = c.codec.UnmarshalBinaryBare(res.Response.Value[:], span)
	return &res.Response, span, err
}

func (c *Client) GetLatestSpan(height uint64) (*htypes.Span, error) {
	res, err := c.ABCIQueryWithOptions(
		"custom/bor/latest-span",
		nil,
		client.ABCIQueryOptions{Prove: true, Height: int64(height)})
	if err != nil {
		log.Error("tendermint_client.GetSpanRes failed", "block", height, "err", err)
		return nil, err
	}

	var span = new(htypes.Span)
	err = json.Unmarshal(res.Response.Value, &span)
	if err != nil {
		log.Error("tendermint_client.GetLatestSpan - unmarshal failed", "height", height, "err", err)
		return nil, err
	}

	return span, nil
}

func (c *Client) GetSpan(id uint64) (*htypes.Span, error) {
	res, err := c.ABCIQueryWithOptions(
		"custom/bor/span",
		GetSpanKey(id), client.ABCIQueryOptions{},
	)
	if err != nil {
		return nil, err
	}
	var span = new(htypes.Span)
	err = json.Unmarshal(res.Response.Value, &span)
	if err != nil {
		log.Error("tendermint_client.GetSpan - unmarshal failed", "spanId", id, "err", err)
		return nil, err
	}
	return span, nil
}

func (c *Client) GetLatestHeight() (uint64, error) {
	res, err := c.Status()
	if err != nil {
		return 0, err
	}
	return uint64(res.SyncInfo.LatestBlockHeight), nil
}

func (c *Client) GetValidators(height uint64) (validators []*tmtypes.Validator, err error) {
	h := int64(height)
	vr, err := c.Validators(&h)
	if err != nil {
		return nil, err
	}
	return vr.Validators, nil
}

func (c *Client) GetCosmosHeader(height uint64) (*cosmos.CosmosHeader, error) {
	h := int64(height)
	rc, err := c.Commit(&h)
	if err != nil {
		return nil, err
	}
	vs, err := c.GetValidators(height)
	if err != nil {
		return nil, err
	}
	return &cosmos.CosmosHeader{
		Header:  *rc.Header,
		Commit:  rc.Commit,
		Valsets: vs,
	}, nil
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
