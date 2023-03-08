package apt

import (
	"context"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/portto/aptos-go-sdk/client"
	"github.com/portto/aptos-go-sdk/models"

	"github.com/polynetwork/bridge-common/chains"
	"github.com/polynetwork/bridge-common/util"
)


type Rpc = client.AptosClient

type Client struct {
	Rpc
	address string
	index int
}

func New(url string) *Client {
	client := client.NewAptosClient(url)
	return &Client{
		Rpc: client,
		address: url,
	}
}

func (c *Client) Address() string {
	return c.address
}

func (c *Client) Balance(account models.AccountAddress, coin models.TypeTagStruct) (balance uint64, err error) {
	resp := new(struct { Data client.CoinStoreResource })
	err = c.GetResource(context.Background(), hex.EncodeToString(account[:]), fmt.Sprintf("0x1::coin::CoinStore<%s>", coin.ToString()), resp)
	if err == nil {
		balance, err = strconv.ParseUint(resp.Data.Coin.Value, 10, 0)
	}
	return
}

var (
	COIN_STORE_RE = regexp.MustCompile(`0x1::coin::CoinStore<(\w*)::(\w*)::(\w*)\>`)
)

func (c *Client) Balances(account models.AccountAddress) (balances map[string]uint64, err error) {
	resp, err := c.GetAccountResources(context.Background(), hex.EncodeToString(account[:]), nil)
	if err == nil {
		balances = make(map[string]uint64)
		for _, res := range resp {
			ret := COIN_STORE_RE.FindStringSubmatch(res.Type)
			if len(ret) != 4 {
				continue
			}
			balance, err := strconv.ParseUint(res.Data.CoinStoreResource.Coin.Value, 10, 0)
			if err != nil { return nil, err }
			addr, err := models.HexToAccountAddress(ret[1])
			if err != nil { return nil, err }
			balances[models.TypeTagStruct{Address: addr, Module: ret[2], Name: ret[3]}.ToString()] = balance
		}
	}
	return
}

func (c *Client) GetLatestHeight() (uint64, error) {
	info, err := c.LedgerInformation(context.Background())
	if err != nil { return 0, err  }
	return strconv.ParseUint(info.LedgerVersion, 10, 0)
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

func (s *SDK) Select() *Client {
	return s.nodes[s.ChainSDK.Select()]
}

func (s *SDK) Broadcast(ctx context.Context, tx models.UserTransaction, opts ...interface{}) (resp *client.TransactionResp, err error) {
	nodes := s.Nodes()
	for _, idx := range nodes[1:] {
		go func(id int) {
			_, _ = s.nodes[id].SubmitTransaction(ctx, tx)
		} (idx)
	}
	return s.nodes[nodes[0]].SubmitTransaction(ctx, tx, opts...)
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
