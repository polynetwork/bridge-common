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
	"io/ioutil"
	"path/filepath"

	"github.com/gagliardetto/solana-go"
	sol "github.com/polynetwork/bridge-common/chains/solana"
)

type SolanaWallet struct {
	accounts []*solana.PrivateKey
	sdk      *sol.SDK
}

func NewSolanaWallet(config *Config, sdk *sol.SDK) (w *SolanaWallet, err error) {
	files, err := ioutil.ReadDir(config.Path)
	if err != nil {
		return
	}

	w = &SolanaWallet{[]*solana.PrivateKey{}, sdk}
	for _, file := range files {
		if !file.IsDir() {
			key, err := solana.PrivateKeyFromSolanaKeygenFile(filepath.Join(config.Path, file.Name()))
			if err != nil {
				return nil, err
			}
			w.accounts = append(w.accounts, &key)
		}
	}
	return
}

func (w *SolanaWallet) Invoke(instructions []solana.Instruction) (sig solana.Signature, err error) {
	if len(w.accounts) < 1 {
		return sig, fmt.Errorf("No accounts available")
	}
	return w.InvokeWithAccount(w.accounts[0], instructions)

}

func (w *SolanaWallet) InvokeWithAccount(account *solana.PrivateKey, instructions []solana.Instruction) (sig solana.Signature, err error) {
	node := w.sdk.Node()
	recent, err := node.GetRecentBlockhash(context.TODO(), node.Commitment)
	if err != nil {
		return
	}
	tx, err := solana.NewTransaction(instructions, recent.Value.Blockhash,
		solana.TransactionPayer(account.PublicKey()))
	if err != nil {
		return
	}
	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if account.PublicKey().Equals(key) {
				return account
			}
			return nil
		},
	)
	if err != nil {
		return
	}
	return node.SendAndConfirm(tx)
}
