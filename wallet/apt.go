package wallet

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/portto/aptos-go-sdk/client"
	"github.com/portto/aptos-go-sdk/models"

	"github.com/polynetwork/bridge-common/chains/apt"
	"github.com/polynetwork/bridge-common/log"
)

var APT_ADDR0x1 models.AccountAddress
var APT_TOKEN_TAG models.TypeTagStruct

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
	ChainID  uint8
	config   *Config
	accounts map[models.AccountAddress]models.PrivateKey
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

	tx := &models.Transaction{}
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
	if err != nil {
		err = fmt.Errorf("SimulateTransaction error: %v", err)
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

	resp, err := w.sdk.Node().SubmitTransaction(ctx, tx.UserTransaction)
	log.Info("Sent apt tx", "limit", resps[0].MaxGasAmount, "gas_price", resps[0].GasUnitPrice, "seq", info.SequenceNumber, "err", err)
	if err != nil {
		return
	}

	hash = resp.Hash
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
		sdk, info.ChainID, config, map[models.AccountAddress]models.PrivateKey{},
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
