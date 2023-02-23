package aptos

import (
	"context"
	"github.com/polynetwork/bridge-common/chains"
	"github.com/polynetwork/bridge-common/util"
	"github.com/portto/aptos-go-sdk/client"
	"github.com/portto/aptos-go-sdk/models"
	"strconv"
	"time"
)

type Rpc = client.AptosClient

type Client struct {
	Rpc
	address string
}

func New(url string) *Client {
	c := client.NewAptosClient(url)
	return &Client{
		Rpc:     c,
		address: url,
	}
}

func (c *Client) Address() string {
	return c.address
}

func (c *Client) GetLatestHeight() (uint64, error) {
	ledgerInfo, err := c.LedgerInformation(context.Background())
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(ledgerInfo.LedgerVersion, 10, 64)
}

func (c *Client) GetEvents(filter *EventFilter) ([]models.Event, error) {
	return c.GetEventsByCreationNumber(context.Background(), filter.Address, filter.CreationNumber, filter.Query)
}

func (c *Client) GetTxByVersion(version uint64) (*client.TransactionResp, error) {
	return c.GetTransactionByVersion(context.Background(), version)
}

func (c *Client) GetTxByHash(hash string) (*client.TransactionResp, error) {
	return c.GetTransactionByHash(context.Background(), hash)
}

func (c *Client) GetVersionByTx(hash string) (uint64, error) {
	tx, err := c.GetTransactionByHash(context.Background(), hash)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(tx.Version, 10, 64)
}

type SDK struct {
	*chains.ChainSDK
	nodes   []*Client
	options *chains.Options
}

func NewSDK(chainID uint64, urls []string, interval time.Duration, maxGap uint64) (*SDK, error) {

	clients := make([]*Client, len(urls))
	nodes := make([]chains.SDK, len(urls))
	for i, url := range urls {
		c := New(url)
		nodes[i] = c
		clients[i] = c
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

func (s *SDK) Node() *Client {
	return s.nodes[s.ChainSDK.Index()]
}

func (s *SDK) Select() *Client {
	return s.nodes[s.ChainSDK.Select()]
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
