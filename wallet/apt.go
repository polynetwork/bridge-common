package wallet

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/portto/aptos-go-sdk/client"
	"github.com/portto/aptos-go-sdk/models"

	"github.com/polynetwork/bridge-common/chains/apt"
	"github.com/polynetwork/bridge-common/log"
)

var APT_ADDR0x1 models.AccountAddress
var APT_TOKEN_TAG models.TypeTagStruct

type (
	AptAccountAddress = models.AccountAddress
	AptTypeTagStruct = models.TypeTagStruct
	AptTransactionPayload = models.TransactionPayload
)

const (
	APT_GAS_PRICE = "100"
	APT_GAS_LIMIT = "2000000"
)

func init () {
	APT_ADDR0x1, _ = models.HexToAccountAddress("0x1")
	APT_TOKEN_TAG = models.TypeTagStruct{
		Address: APT_ADDR0x1,
		Module:  "aptos_coin",
		Name:    "AptosCoin",
	}
}

type AptWallet struct {
	sdk      *apt.SDK
	Broadcast bool
	ChainID  uint8
	config   *Config
	accounts map[models.AccountAddress]models.PrivateKey
	nonces   map[models.AccountAddress]string
	sync.Mutex
}

func (w *AptWallet) GetNonce(account models.AccountAddress) (nonce string, err error) {
	var ok bool
	w.Lock()
	nonce, ok = w.nonces[account]
	if ok {
		delete(w.nonces, account)
	}
	w.Unlock()
	if !ok {
		info, err := w.GetAccountInfo(account)
		if err != nil {
			return "", err
		}
		log.Info("Apt wallet fetching account seq", "account", account.PrefixZeroTrimmedHex(), "seq", info.SequenceNumber)
		return info.SequenceNumber, nil
	} 
	return
}

func (w *AptWallet) UpdateNonce(account models.AccountAddress, _success bool) (err error) {
	info, err := w.GetAccountInfo(account)
	if err != nil {
		return
	}

	log.Info("Apt wallet updating account", "account", account.PrefixZeroTrimmedHex(), "seq", info.SequenceNumber)
	w.Lock()
	w.nonces[account] = info.SequenceNumber
	w.Unlock()
	return
}

func (w *AptWallet) Accounts() (list []models.AccountAddress) {
	for key := range w.accounts {
		list = append(list, key)
	}
	return
}

func (w *AptWallet) GetAccountInfo(account models.AccountAddress) (*client.AccountInfo, error) {
	return w.sdk.Node().GetAccount(context.Background(), hex.EncodeToString(account[:]))
}

func(w *AptWallet) Transfer(ctx context.Context, from, to models.AccountAddress, amount uint64, ttl time.Duration) (hash string, err error) {
	payload := models.EntryFunctionPayload{
		Module: models.Module{
			Address: APT_ADDR0x1,
			Name:    "coin",
		},
		Function:      "transfer",
		TypeArguments: []models.TypeTag{APT_TOKEN_TAG},
		Arguments:     []interface{}{to, amount},
	}
	return w.Send(ctx, &from, payload, ttl)
}

func (w *AptWallet) CreateAccount(ctx context.Context, account *models.AccountAddress, to models.AccountAddress, ttl time.Duration) (hash string, err error) {
	payload := models.EntryFunctionPayload{
		Module: models.Module{
			Address: APT_ADDR0x1,
			Name:    "aptos_account",
		},
		Function:  "create_account",
		Arguments: []interface{}{to},
	}
	return w.Send(ctx, account, payload, ttl)
}

func (w *AptWallet) SendWithOptions(ctx context.Context, account *models.AccountAddress, payload models.TransactionPayload, ttl time.Duration, seq, limit, price string, priceX float64) (hash string, err error) {
	if account == nil {
		if len(w.accounts) == 0 {
			err = fmt.Errorf("no apt account available")
			return
		}
		for addr := range w.accounts {
			account = &addr
			break
		}
		// clear seq
		seq = ""
	}

	key, ok := w.accounts[*account]
	if !ok {
		err = fmt.Errorf("account not found in wallet, %x", (*account)[:])
		return
	}

	if seq == "" {
		info, err := w.GetAccountInfo(*account)
		if err != nil {
			return "", err
		}
		seq = info.SequenceNumber
	}

	tx := new(models.Transaction)
	err = tx.SetChainID(w.ChainID).
		SetSender(hex.EncodeToString((*account)[:])).
		SetPayload(payload).
		SetExpirationTimestampSecs(uint64(time.Now().Add(ttl).Unix())).
		SetSequenceNumber(seq).Error()
	if err != nil {
		return
	}

	siger := models.NewSingleSigner(key)

	tx.Authenticator = models.TransactionAuthenticatorEd25519{
		PublicKey: siger.PublicKey,
	}

	if limit == "" {
		resps, err := w.sdk.Node().SimulateTransaction(ctx, tx.UserTransaction, true, true)
		if err != nil || len(resps) == 0 || !resps[0].Success {
			if len(resps) > 0 {
				err = fmt.Errorf("SimulateTransaction failure, vm_status: %s", resps[0].VmStatus)
			} else {
				err = fmt.Errorf("SimulateTransaction error: %v", err)
			}
			return "", err
		}
		limit = resps[0].MaxGasAmount
		if price == "" {
			price = resps[0].GasUnitPrice
		}
	} else if price == "" {
		price = APT_GAS_PRICE
	}
	
    gasPrice, err := strconv.ParseUint(price, 10, 0)
	if err != nil { 
		return
	}
	if priceX > 0 {
		gasPrice = uint64(float64(gasPrice) * priceX)
	}

	err = tx.SetMaxGasAmount(limit).SetGasUnitPrice(gasPrice).Error()
	if err != nil {
		return
	}

	err = siger.Sign(tx).Error()
	if err != nil {
		return
	}

	var resp *client.TransactionResp
	if w.Broadcast {
		resp, err = w.sdk.Broadcast(ctx, tx.UserTransaction)
	} else {
		resp, err = w.sdk.Node().SubmitTransaction(ctx, tx.UserTransaction)
	}
	if err != nil {
		return
	}
	hash = resp.Hash
	log.Info("Sent apt tx", "limit", resp.MaxGasAmount, "gas_price", resp.GasUnitPrice, "seq", resp.SequenceNumber, "hash", hash, "err", err)
	return
}

func (w *AptWallet) Send(ctx context.Context, account *models.AccountAddress, payload models.TransactionPayload, ttl time.Duration) (hash string, err error) {
	if account == nil {
		if len(w.accounts) == 0 {
			err = fmt.Errorf("no apt account available")
			return
		}
		for addr := range w.accounts {
			account = &addr
			break
		}
	}

	key, ok := w.accounts[*account]
	if !ok {
		err = fmt.Errorf("account not found in wallet, %x", (*account)[:])
		return
	}
	
	info, err := w.GetAccountInfo(*account)
	if err != nil {
		return
	}

	tx := new(models.Transaction)
	err = tx.SetChainID(w.ChainID).
		SetSender(hex.EncodeToString((*account)[:])).
		SetPayload(payload).
		SetExpirationTimestampSecs(uint64(time.Now().Add(ttl).Unix())).
		SetSequenceNumber(info.SequenceNumber).Error()
	if err != nil {
		return
	}

	siger := models.NewSingleSigner(key)

	tx.Authenticator = models.TransactionAuthenticatorEd25519{
		PublicKey: siger.PublicKey,
	}

	resps, err := w.sdk.Node().SimulateTransaction(ctx, tx.UserTransaction, true, true)
	if err != nil || len(resps) == 0 || !resps[0].Success {
		if len(resps) > 0 {
			err = fmt.Errorf("SimulateTransaction failure, vm_status: %s", resps[0].VmStatus)
		} else {
			err = fmt.Errorf("SimulateTransaction error: %v", err)
		}
		return
	}

	err = tx.SetMaxGasAmount(resps[0].MaxGasAmount).SetGasUnitPrice(resps[0].GasUnitPrice).Error()
	if err != nil {
		return
	}

	err = siger.Sign(tx).Error()
	if err != nil {
		return
	}

	var resp *client.TransactionResp
	if w.Broadcast {
		resp, err = w.sdk.Broadcast(ctx, tx.UserTransaction)
	} else {
		resp, err = w.sdk.Node().SubmitTransaction(ctx, tx.UserTransaction)
	}
	if err != nil {
		return
	}
	hash = resp.Hash
	log.Info("Sent apt tx", "limit", resps[0].MaxGasAmount, "gas_price", resps[0].GasUnitPrice, "seq", info.SequenceNumber, "hash", hash, "err", err)
	return
}

func NewAptWallet(config *Config, sdk *apt.SDK) (w *AptWallet) {
	if config.ReadFile == nil {
		config.ReadFile = ioutil.ReadFile
	}
	data, err := config.ReadFile(config.Path)
	if err != nil {
		log.Error("Failed to load wallet file", "path", config.Path)
		return
	}
	payload := map[string]string{}
	err = json.Unmarshal(data, &payload)
	if err != nil {
		log.Error("Failed to decode wallet file", "path", config.Path)
		return
	}

	info, err := sdk.Node().LedgerInformation(context.Background())
	if err != nil {
		log.Fatal("Failed to setup apt wallet", "err", err)
	}

	w = &AptWallet{
		sdk: sdk, ChainID: info.ChainID, config: config, accounts: make(map[models.AccountAddress]models.PrivateKey),
		nonces: make(map[models.AccountAddress]string),
	}

	for addrHex, keyHex := range payload {
		addr, err := models.HexToAccountAddress(addrHex)
		if err != nil {
			log.Error("Invalid star wallet address hex", "hex", config.Path)
			continue
		}
		bytes, err := hex.DecodeString(strings.Replace(keyHex, "0x", "", 1))
		if err != nil {
			log.Error("Invalid star wallet private key hex")
			continue
		}
		w.accounts[addr] = ed25519.NewKeyFromSeed(bytes)
	}
	return
}
