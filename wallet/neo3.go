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
	"fmt"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/tx"
	"github.com/joeqian10/neo3-gogogo/wallet"
	"github.com/polynetwork/bridge-common/chains/neo3"
	"github.com/polynetwork/bridge-common/log"
)

const EMPTY = ""

type Neo3Wallet struct {
	sdk *neo3.SDK
	*wallet.NEP6Wallet
	config *Config
}

func NewNeo3Wallet(config *Config, sdk *neo3.SDK) (*Neo3Wallet, error) {
	ps := helper.ProtocolSettings{
		Magic:          config.Neo3Magic,
		AddressVersion: helper.DefaultAddressVersion,
	}
	w, err := wallet.NewNEP6Wallet(config.Neo3Path, &ps, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("neo3.NewNEP6Wallet error: %v", err)
	}
	err = w.Unlock(config.Neo3Pwd)
	if err != nil {
		return nil, fmt.Errorf("neo3.NEP6Wallet.Unlock error: %v", err)
	}
	if len(w.Accounts) == 0 {
		return nil, fmt.Errorf("no available account in neo3 wallet")
	}
	return &Neo3Wallet{sdk: sdk, NEP6Wallet: w, config: config}, nil
}

func (w *Neo3Wallet) SendTransaction(script []byte) (string, error) {
	return w.SendTransactionWithAccount(script, w.Account())
}

func (w *Neo3Wallet) SendTransactionWithAccount(script []byte, account *wallet.NEP6Account) (string, error) {
	wh := wallet.NewWalletHelperFromWallet(w.sdk.Node(), w.NEP6Wallet)
	a := account.GetScriptHash()
	gasAmount, err := wh.GetBalanceFromAccount(tx.GasToken, a)
	if err != nil {
		return EMPTY, fmt.Errorf("WalletHelper.GetBalanceFromAccount error: %v", err)
	}
	pair := &wallet.AccountAndBalance{
		Account: a,
		Value:   gasAmount,
	}
	trx, err := wh.MakeTransaction(script, nil, []tx.ITransactionAttribute{}, []*wallet.AccountAndBalance{pair})
	if err != nil {
		return EMPTY, fmt.Errorf("WalletHelper.MakeTransaction: %v", err)
	}
	trx, err = wh.SignTransaction(trx, w.config.Neo3Magic)
	if err != nil {
		return EMPTY, fmt.Errorf("WalletHelper.SignTransaction: %v", err)
	}
	res := w.sdk.Node().SendRawTransaction(crypto.Base64Encode(trx.ToByteArray()))
	if res.HasError() {
		return EMPTY, fmt.Errorf("neo3.SendRawTransaction error: %s", res.GetErrorInfo())
	}
	hash := trx.GetHash().String()
	log.Info("send neo3 transaction", "hash", hash)
	return hash, nil
}

func (w *Neo3Wallet) Account() *wallet.NEP6Account {
	return &w.Accounts[0]
}
