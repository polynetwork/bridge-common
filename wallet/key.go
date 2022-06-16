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
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
)

type KeyProvider struct {
	Address common.Address
	priv *ecdsa.PrivateKey
}

func NewKeyProvider(key string) (p *KeyProvider, err error) {
	p = new(KeyProvider)
	p.priv, err = crypto.HexToECDSA(key)
	if err != nil {
		return nil, fmt.Errorf("%w invalid key hex", err)
	}
	p.Address = crypto.PubkeyToAddress(p.priv.PublicKey)
	return
}

func (p *KeyProvider) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	if !bytes.Equal(account.Address.Bytes(), p.Address.Bytes()) {
		return nil, fmt.Errorf("unexpected account unmatch %s, %s", account.Address.String(), p.Address.String())
	}
	signer := types.LatestSignerForChainID(chainID)
	return types.SignTx(tx, signer, p.priv)
}

func (p *KeyProvider) Init(accounts.Account) (err error) {
	return
}

func (p *KeyProvider) Accounts() []accounts.Account {
	return []accounts.Account{{Address: p.Address}}
}

func (p *KeyProvider) Save(path, passphrase string) (error) {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	_, err := ks.ImportECDSA(p.priv, passphrase)
	return err
}
