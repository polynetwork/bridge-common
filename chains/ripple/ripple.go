package ripple

import (
	"github.com/polynetwork/bridge-common/chains"
	"github.com/polynetwork/bridge-common/util"
	ripple_sdk "github.com/polynetwork/ripple-sdk"
	"time"
)

type Rpc = ripple_sdk.RippleSdk

type Client struct {
	*Rpc
	address string
}

func New(url string) *Client {
	sdk := ripple_sdk.NewRippleSdk()
	sdk.NewRpcClient().SetAddress(url)

	return &Client{
		Rpc:     sdk,
		address: url,
	}
}

func (c *Client) Address() string {
	return c.address
}

func (c *Client) GetLatestHeight() (uint64, error) {
	h, err := c.GetRpcClient().GetCurrentHeight()
	return uint64(h), err
}

func (c *Client) GetFee() (float64, error) {
	fee, err := c.GetRpcClient().GetFee()
	if err != nil {
		return 0, err
	}
	rippleFee, err := fee.Drops.OpenLedgerFee.NonNative()
	if err != nil {
		return 0, err
	}
	return rippleFee.Float(), nil
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
