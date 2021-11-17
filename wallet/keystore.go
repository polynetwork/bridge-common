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
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/polynetwork/bridge-common/log"
)

type KeyStoreProviderConfig struct {
	Passwords map[string]string
	Path      string
}

type KeyStoreProvider struct {
	*keystore.KeyStore
	config *KeyStoreProviderConfig
}

func NewKeyStoreProvider(config *KeyStoreProviderConfig) *KeyStoreProvider {
	for addr, pass := range config.Passwords {
		config.Passwords[strings.ToLower(addr)] = pass
	}
	store := keystore.NewKeyStore(config.Path, keystore.StandardScryptN, keystore.StandardScryptP)
	return &KeyStoreProvider{config: config, KeyStore: store}
}

func (p *KeyStoreProvider) Accounts() []accounts.Account {
	list := []accounts.Account{}
	accounts := p.KeyStore.Accounts()
	for _, account := range accounts {
		_, ok := p.config.Passwords[strings.ToLower(account.Address.String())]
		if ok {
			list = append(list, account)
		}
	}
	return list
}

func (p *KeyStoreProvider) Init(account accounts.Account) error {
	pass, ok := p.config.Passwords[strings.ToLower(account.Address.String())]
	if !ok {
		log.Warn("Skipping account unlock", "account", account)
		return nil
	}
	return p.Unlock(account, pass)
}
