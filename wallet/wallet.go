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

package wallet

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/polynetwork/bridge-common/chains/eth"
	"github.com/polynetwork/bridge-common/log"
	"github.com/polynetwork/bridge-common/util"
)

type Config struct {
	ChainId           uint64
	KeyStoreProviders []*KeyStoreProviderConfig
	Nodes             []string

	// NEO wallet
	Path     string
	Password string
	SysFee   float64
	NetFee   float64

	// ONT wallet
	GasPrice uint64
	GasLimit uint64
}

type IWallet interface {
	Init() error
	Send(addr common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, gasPriceX *big.Float, data []byte) (hash string, err error)
	SendWithAccount(account accounts.Account, addr common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, gasPriceX *big.Float, data []byte) (hash string, err error)
	EstimateWithAccount(account accounts.Account, addr common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, gasPriceX *big.Float, data []byte) (hash string, err error)
	Accounts() []accounts.Account
	GetBalance(common.Address) (*big.Int, error)
}

type Wallet struct {
	sync.RWMutex
	chainId   uint64
	providers map[accounts.Account]Provider
	provider  Provider                           // active account provider
	account   accounts.Account                   // active account
	nonces    map[accounts.Account]NonceProvider // account nonces
	sdk       *eth.SDK
	accounts  []accounts.Account
	cursor    int
	config    *Config
}

type Provider interface {
	SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error)
	Init(accounts.Account) error
	Accounts() []accounts.Account
}

func New(config *Config, sdk *eth.SDK) *Wallet {
	w := &Wallet{
		config: config, sdk: sdk, chainId: config.ChainId, providers: map[accounts.Account]Provider{},
		nonces:   map[accounts.Account]NonceProvider{},
		accounts: []accounts.Account{},
	}
	for _, c := range config.KeyStoreProviders {
		w.AddProvider(NewKeyStoreProvider(c))
	}
	return w
}

func (w *Wallet) GetBalance(address common.Address) (balance *big.Int, err error) {
	return w.sdk.Node().BalanceAt(context.Background(), address, nil)
}

func (w *Wallet) AddProvider(p Provider) {
	w.Lock()
	defer w.Unlock()

	for _, a := range p.Accounts() {
		w.providers[a] = p
	}
}

func (w *Wallet) Accounts() []accounts.Account {
	w.RLock()
	defer w.RUnlock()
	return w.accounts
}

func (w *Wallet) Init() (err error) {
	w.updateAccounts()
	w.Lock()
	defer w.Unlock()
	for a, p := range w.providers {
		err = p.Init(a)
		if err != nil {
			return
		}
	}
	if len(w.accounts) == 0 {
		return fmt.Errorf("No valid account provided")
	}
	w.account = w.accounts[0]
	w.provider = w.providers[w.account]
	if w.sdk != nil {
		w.VerifyChainId()
	}
	return
}

func (w *Wallet) VerifyChainId() {
	chainId, err := w.sdk.Node().ChainID(context.Background())
	if err != nil {
		util.Fatal("Failed to verfiy chain id %v", err)
	}
	id := uint64(chainId.Int64())
	if w.chainId != 0 && w.chainId != id {
		util.Fatal("ChainID does not match specified %v, on chain: %v", w.chainId, id)
	}
	w.chainId = id
}

func (w *Wallet) GetAccount(account accounts.Account) (provider Provider, nonces NonceProvider) {
	w.RLock()
	defer w.RUnlock()
	return w.providers[account], w.nonces[account]
}

func (w *Wallet) Send(addr common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, gasPriceX *big.Float, data []byte) (hash string, err error) {
	account, _, _ := w.Select()
	return w.sendWithAccount(false, account, addr, amount, gasLimit, gasPrice, gasPriceX, data)
}

func (w *Wallet) SendWithAccount(account accounts.Account, addr common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, gasPriceX *big.Float, data []byte) (hash string, err error) {
	return w.sendWithAccount(false, account, addr, amount, gasLimit, gasPrice, gasPriceX, data)
}

func (w *Wallet) EstimateWithAccount(account accounts.Account, addr common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, gasPriceX *big.Float, data []byte) (hash string, err error) {
	return w.sendWithAccount(true, account, addr, amount, gasLimit, gasPrice, gasPriceX, data)
}

func (w *Wallet) sendWithAccount(dry bool, account accounts.Account, addr common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, gasPriceX *big.Float, data []byte) (hash string, err error) {
	if gasPrice == nil || gasPrice.Sign() <= 0 {
		gasPrice, err = w.GasPrice()
		if err != nil {
			err = fmt.Errorf("Get gas price error %v", err)
			return
		}
		if gasPriceX != nil {
			gasPrice, _ = new(big.Float).Mul(new(big.Float).SetInt(gasPrice), gasPriceX).Int(nil)
		}
	}

	provider, nonces := w.GetAccount(account)
	nonce, err := nonces.Acquire()
	if err != nil {
		return
	}
	if gasLimit == 0 {
		msg := ethereum.CallMsg{From: account.Address, To: &addr, GasPrice: gasPrice, Value: amount, Data: data}
		gasLimit, err = w.sdk.Node().EstimateGas(context.Background(), msg)
		if err != nil {
			nonces.Update(false)
			if strings.Contains(err.Error(), "has been executed") {
				log.Info("Transaction already executed")
				return "", nil
			}

			err = fmt.Errorf("Estimate gas limit error %v, account %s", err, account.Address)
			return
		}
	}

	if dry {
		fmt.Printf(
			"Estimated tx successfully account %s gas_price %s gas_limit %v nonce %v target %s amount %s data %x\n",
			account.Address, gasPrice, gasLimit, nonce, addr, amount, data)
		return
	}

	gasLimit = uint64(1.3 * float32(gasLimit))
	limit := GetChainGasLimit(w.chainId, gasLimit)
	if limit < gasLimit {
		nonces.Update(false)
		err = fmt.Errorf("Send tx estimated gas limit(%v) higher than max %v", gasLimit, limit)
		return
	}
	tx := types.NewTransaction(nonce, addr, amount, limit, gasPrice, data)
	tx, err = provider.SignTx(account, tx, big.NewInt(int64(w.chainId)))
	if err != nil {
		nonces.Update(false)
		err = fmt.Errorf("Sign tx error %v", err)
		return
	}
	log.Info("Compose dst chain tx", "hash", tx.Hash(), "account", account.Address)
	err = w.sdk.Node().SendTransaction(context.Background(), tx)
	//TODO: Check err here before update nonces
	nonces.Update(true)
	return tx.Hash().String(), err
}

func (w *Wallet) Account() (accounts.Account, Provider, NonceProvider) {
	w.RLock()
	defer w.RUnlock()
	return w.account, w.provider, w.nonces[w.account]
}

func (w *Wallet) GasPrice() (price *big.Int, err error) {
	return w.sdk.Node().SuggestGasPrice(context.Background())
}

/*
func (w *Wallet) GasTip() (price *big.Int, err error) {
	return w.sdk.Node().SuggestGasTipCap(context.Background())
}
*/

func (w *Wallet) updateAccounts() {
	w.Lock()
	defer w.Unlock()
	accounts := []accounts.Account{}
	for a, _ := range w.providers {
		accounts = append(accounts, a)
		w.nonces[a] = NewRemoteNonceProvider(w.sdk, a.Address)
	}
	w.accounts = accounts
	w.cursor = 0
}

// Round robin
func (w *Wallet) Select() (accounts.Account, Provider, NonceProvider) {
	w.Lock()
	defer w.Unlock()
	account := w.accounts[w.cursor]
	w.cursor = (w.cursor + 1) % len(w.accounts)
	return account, w.providers[account], w.nonces[account]
}

func (w *Wallet) Upgrade() *EthWallet {
	return w
}
