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
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/polynetwork/bridge-common/chains/eth"
)

type Config struct {
	ChainId           uint64
	KeyStoreProviders []*KeyStoreProviderConfig
}

type IWallet interface {
	Init() error
	Send(addr common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, gasPriceX *big.Float, data []byte) (err error)
	SendWithAccount(account accounts.Account, addr common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, gasPriceX *big.Float, data []byte) (err error)
	Accounts() []accounts.Account
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
	w := &Wallet{config: config, sdk: sdk, chainId: config.ChainId}
	for _, c := range config.KeyStoreProviders {
		w.AddProvider(NewKeyStoreProvider(c))
	}
	return w
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
	return
}

func (w *Wallet) GetAccount(account accounts.Account) (provider Provider, nonces NonceProvider) {
	w.RLock()
	defer w.RUnlock()
	return w.providers[account], w.nonces[account]
}

func (w *Wallet) Send(addr common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, gasPriceX *big.Float, data []byte) (err error) {
	account, _, _ := w.Select()
	return w.SendWithAccount(account, addr, amount, gasLimit, gasPrice, gasPriceX, data)
}

func (w *Wallet) SendWithAccount(account accounts.Account, addr common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, gasPriceX *big.Float, data []byte) (err error) {
	if gasPrice == nil || gasPrice.Sign() <= 0 {
		gasPrice, err = w.GasPrice()
		if err != nil {
			return fmt.Errorf("Get gas price error %v", err)
		}
		if gasPriceX != nil {
			gasPrice, _ = new(big.Float).Mul(new(big.Float).SetInt(gasPrice), gasPriceX).Int(nil)
		}
	}

	provider, nonces := w.GetAccount(account)
	nonce, err := nonces.Acquire()
	if err != nil {
		return err
	}
	if gasLimit == 0 {
		msg := ethereum.CallMsg{From: account.Address, To: &addr, GasPrice: gasPrice, Value: big.NewInt(0), Data: data}
		gasLimit, err = w.sdk.Node().EstimateGas(context.Background(), msg)
		if err != nil {
			nonces.Update(false)
			return fmt.Errorf("Estimate gas limit error %v", err)
		}
	}

	limit := GetChainGasLimit(w.chainId, gasLimit)
	if limit < gasLimit {
		nonces.Update(false)
		return fmt.Errorf("Send tx estimated gas limit(%v) higher than max %v", gasLimit, limit)
	}
	tx := types.NewTransaction(nonce, addr, amount, limit, gasPrice, data)
	tx, err = provider.SignTx(account, tx, big.NewInt(int64(w.chainId)))
	if err != nil {
		nonces.Update(false)
		return fmt.Errorf("Sign tx error %v", err)
	}
	err = w.sdk.Node().SendTransaction(context.Background(), tx)
	//TODO: Check err here before update nonces
	nonces.Update(true)
	return err
}

func (w *Wallet) Account() (accounts.Account, Provider, NonceProvider) {
	w.RLock()
	defer w.RUnlock()
	return w.account, w.provider, w.nonces[w.account]
}

func (w *Wallet) GasPrice() (price *big.Int, err error) {
	return w.sdk.Node().SuggestGasPrice(context.Background())
}

func (w *Wallet) updateAccounts() {
	w.Lock()
	defer w.Unlock()
	accounts := []accounts.Account{}
	for a, _ := range w.providers {
		accounts = append(accounts, a)
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
