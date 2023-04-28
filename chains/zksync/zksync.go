package zksync

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum"
	"github.com/polynetwork/bridge-common/chains"
	"github.com/polynetwork/bridge-common/chains/eth"
	"github.com/polynetwork/bridge-common/util"
	"time"
)

type Client struct {
	*eth.Client
}

func New(url string) *Client {
	return &Client{
		eth.New(url),
	}
}

func (c *Client) GetZkL1BatchNumber(height uint64) (number uint64, err error) {
	head := make(json.RawMessage, 0)
	err = c.Rpc.CallContext(context.Background(), &head, "eth_getBlockByNumber", util.BlockNumArg(height), false)
	if err == nil && len(head) == 0 {
		err = ethereum.NotFound
	}
	if err != nil {
		return
	}
	header := Header(head)
	return header.L1BatchNumber()
}

type SDK struct {
	*chains.ChainSDK
	nodes   []*Client
	l1Nodes *eth.Client
	options *chains.Options
}

func (s *SDK) Node() *Client {
	return s.nodes[s.ChainSDK.Index()]
}

func (s *SDK) Select() *Client {
	return s.nodes[s.ChainSDK.Select()]
}

func (s *SDK) L1Node() *eth.Client {
	return s.l1Nodes
}

func NewSDK(chainID uint64, urls []string, l1Url string, interval time.Duration, maxGap uint64) (*SDK, error) {
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
	l1Client := eth.New(l1Url)
	return &SDK{ChainSDK: sdk, nodes: clients, l1Nodes: l1Client}, nil
}

func WithOptions(chainID uint64, urls []string, l1Url string, interval time.Duration, maxGap uint64) (*SDK, error) {
	sdk, err := util.Single(&SDK{
		options: &chains.Options{
			ChainID:  chainID,
			Nodes:    urls,
			L1Node:   l1Url,
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
	return NewSDK(s.options.ChainID, s.options.Nodes, s.options.L1Node, s.options.Interval, s.options.MaxGap)
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
