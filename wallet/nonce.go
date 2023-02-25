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
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/polynetwork/bridge-common/chains/eth"
	"github.com/polynetwork/bridge-common/log"
)

type NonceProvider interface {
	Acquire() (uint64, error)
	Update(bool) error
}

func NewRemoteNonceProvider(sdk *eth.SDK, address common.Address) *RemoteNonceProvider {
	return &RemoteNonceProvider{sdk, address}
}

type RemoteNonceProvider struct {
	sdk     *eth.SDK
	address common.Address
}

func (p *RemoteNonceProvider) Acquire() (uint64, error) {
	return p.sdk.Node().NonceAt(context.Background(), p.address, nil)
}

func (p *RemoteNonceProvider) Update(success bool) error {
	return nil
}

type DummyNonceProvider uint64
func (p DummyNonceProvider) Acquire() (uint64, error) { return uint64(p), nil }
func (p DummyNonceProvider) Update(_success bool) (err error) { return nil }

func NewCacheNonceProvider(sdk *eth.SDK, address common.Address) *CacheNonceProvider {
	p := &CacheNonceProvider{sdk: sdk, address: address}
	go p.Update(true)
	return p
}

type CacheNonceProvider struct {
	sdk     *eth.SDK
	address common.Address
	sync.Mutex
	nonce *uint64
}

func (p *CacheNonceProvider) Acquire() (uint64, error) {
	p.Lock()
	nonce := p.nonce
	if nonce != nil {
		p.nonce = nil
	}
	p.Unlock()
	if nonce != nil {
		return *nonce, nil
	}
	return p.sdk.Node().NonceAt(context.Background(), p.address, nil)
}

func (p *CacheNonceProvider) Update(_success bool) (err error) {
	nonce, err := p.sdk.Node().NonceAt(context.Background(), p.address, nil)
	if err != nil {
		log.Error("Failed to fetch nonce for account", "err", err)
	} else {
		p.Lock()
		p.nonce = &nonce
		p.Unlock()
		log.Info("Updated account nonce cache", "account", p.address, "nonce", nonce)
	}
	return
}
