/*
 * Copyright (C) 2022 The poly network Authors
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
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/polynetwork/bridge-common/chains/star"
	"github.com/polynetwork/bridge-common/log"
	"github.com/starcoinorg/starcoin-go/client"
	"github.com/starcoinorg/starcoin-go/types"
)

type StarWallet struct {
	sdk *star.SDK
	config *Config
	accounts map[types.AccountAddress]types.Ed25519PrivateKey
}

func (w *StarWallet) Accounts() (list []types.AccountAddress) {
	for key := range w.accounts {
		list = append(list, key)
	}
	return
}

func (w *StarWallet) Send(ctx context.Context, account *types.AccountAddress, payload types.TransactionPayload) (hash string, err error) {
	if account == nil {
		if len(w.accounts) == 0 {
			err = fmt.Errorf("No star account available")
			return
		}
		for addr := range w.accounts {
			account = &addr
			break
		}
	}
	nonce, err := w.sdk.Node().GetAccountSequenceNumber(ctx, client.BytesToHexString(account[:]))
	if err != nil { return }
	gasPrice, err := w.sdk.Node().GetGasUnitPrice(ctx)
	if err != nil { return }
	tx, err := w.sdk.Node().BuildRawUserTransaction(ctx, *account, payload, gasPrice, w.config.GasLimit, nonce)
	if err != nil { return }
	return w.sdk.Node().SubmitTransaction(ctx, w.accounts[*account], tx)
}

func NewStarWallet(config *Config, sdk *star.SDK) (w *StarWallet) {
	data, err := ioutil.ReadFile(config.Path)
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

	w = &StarWallet{
		sdk, config, map[types.AccountAddress]types.Ed25519PrivateKey{},
	}

	for addrHex, keyHex := range payload {
		addr, err := types.ToAccountAddress(addrHex)
		if err != nil {
			log.Error("Invalid star wallet address hex", "hex", config.Path)
			continue
		}
		bytes, err := client.HexStringToBytes(keyHex)
		if err != nil {
			log.Error("Invalid star wallet private key hex")
			continue
		}
		w.accounts[*addr] = bytes
	}
	return
}