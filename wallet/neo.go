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

	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/tx"
	"github.com/joeqian10/neo-gogogo/wallet"
	"github.com/polynetwork/bridge-common/chains/neo"
	"github.com/polynetwork/bridge-common/log"
)

type NeoWallet struct {
	sdk *neo.SDK
	*wallet.Wallet
	config *Config
}

func NewNeoWallet(config *Config, sdk *neo.SDK) *NeoWallet {
	w, err := wallet.NewWalletFromFile(config.Path)
	if err != nil {
		log.Error("Failed to load neo wallet file", "err", err)
		return nil
	}
	err = w.DecryptAll(config.Password)
	if err != nil {
		log.Error("Failed to decrypt neo wallet with password", "err", err)
		return nil
	}
	return &NeoWallet{sdk: sdk, Wallet: w, config: config}
}

func (w *NeoWallet) Invoke(script []byte, attributes []*tx.TransactionAttribute) (hash string, err error) {
	return w.SendInvocation(script, attributes, w.config.SysFee, w.config.NetFee)
}

func (w *NeoWallet) SendInvocation(
	script []byte, attributes []*tx.TransactionAttribute, sysFee, netFee float64,
) (hash string, err error) {
	return w.SendInvocationWithAccount(w.Account(), script, attributes, sysFee, netFee)
}

func (w *NeoWallet) SendInvocationWithAccount(
	account *wallet.Account, script []byte, attributes []*tx.TransactionAttribute, sysFee, netFee float64,
) (hash string, err error) {
	address, err := helper.AddressToScriptHash(account.Address)
	if err != nil {
		return
	}
	client := w.sdk.Node()
	tb := tx.NewTransactionBuilder(client.Address())
	sysfee := helper.Fixed8FromFloat64(sysFee)
	netfee := helper.Fixed8FromFloat64(netFee)
	t, err := tb.MakeInvocationTransaction(script, address, attributes, address, sysfee, netfee)
	if err != nil {
		return
	}
	err = tx.AddSignature(t, account.KeyPair)
	if err != nil {
		return
	}
	res := client.SendRawTransaction(t.RawTransactionString())
	if res.HasError() {
		err = fmt.Errorf("Send neo raw transaction err %s", res.ErrorResponse.Error.Message)
		return
	}
	hash = t.HashString()
	log.Info("Send neo transaction", "hash", hash)
	return
}

func (w *NeoWallet) Account() *wallet.Account {
	return w.Accounts[0]
}

func (w *NeoWallet) Init() (err error) {
	return
}
