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
	"math/big"
	"sync"
	"time"

	"github.com/polynetwork/bridge-common/base"
	"github.com/polynetwork/bridge-common/chains/apt"
	"github.com/polynetwork/bridge-common/chains/eth"
	"github.com/polynetwork/bridge-common/log"
	"github.com/polynetwork/bridge-common/util"
)

var softGasLimits = map[uint64]uint64{}

func HardGasLimit(chain uint64) uint64 {
	switch chain {
	case base.ETH:
		return 300000
	default:
		return 10000000000000
	}
}

// NOTE: call on init
func SetGasLimit(chain, limit uint64) {
	softGasLimits[chain] = limit
}

func GetChainGasLimit(chain, limit uint64) uint64 {
	soft := softGasLimits[chain]
	if soft > 0 && limit > soft {
		limit = soft
	}
	hard := HardGasLimit(chain)
	if limit > hard {
		limit = hard
	}
	return limit
}

var balanceMinLimits = map[uint64]*big.Int{}

// NOTE: call on init
func SetBalanceLimit(chain uint64, limit *big.Int) {
	balanceMinLimits[chain] = limit
}

func HasBalance(chain uint64, balance *big.Int) bool {
	limit := balanceMinLimits[chain]
	if limit == nil {
		limit = util.SetDecimals(big.NewInt(1), 17)
	}
	return balance != nil && balance.Cmp(limit) >= 0
}

type GasPriceOracle interface {
	Price() *big.Int
	PriceWithTip() (*big.Int, *big.Int)
}

type RemoteGasPriceOracle struct {
	sdk *eth.SDK
	price, tip *big.Int
	sync.RWMutex
	upgrade bool
}

func NewRemoteGasPriceOracle(sdk *eth.SDK, upgrade bool, interval time.Duration) (o *RemoteGasPriceOracle, err error) {
	price, err := sdk.Node().SuggestGasPrice(context.Background())
	if err != nil {
		return
	}
	var tip *big.Int
	if upgrade {
		tip, err = sdk.Node().SuggestGasTipCap(context.Background())
		if err != nil { return}
	}
	o = &RemoteGasPriceOracle{sdk: sdk, price: price, upgrade: upgrade, tip: tip}
	go o.update(interval, upgrade)
	return
}

func (o *RemoteGasPriceOracle) update(interval time.Duration, upgrade bool) {
	timer := time.NewTicker(interval)
	for range timer.C {
		var tip *big.Int
		price, err := o.sdk.Node().SuggestGasPrice(context.Background())
		if err == nil && upgrade {
			tip, err = o.sdk.Node().SuggestGasTipCap(context.Background())
		}

		if err != nil {
			log.Error("Failed to update gas price", "err", err)
		} else {
			o.Lock()
			o.price = price
			o.tip = tip
			o.Unlock()
		}
	}
}

func (o *RemoteGasPriceOracle) PriceWithTip() (price *big.Int, tip *big.Int) {
	o.RLock()
	defer o.RUnlock()
	return o.price, o.tip
}

func (o *RemoteGasPriceOracle) Price() *big.Int {
	o.RLock()
	defer o.RUnlock()
	return o.price
}

type AptGasPriceOracle interface {
	Price() (uint64, uint64, uint64)
}

type RemoteAptGasPriceOracle struct {
	sdk *apt.SDK
	price, low, high uint64
	sync.RWMutex
}

func NewRemoteAptGasPriceOracle(sdk *apt.SDK, interval time.Duration) (o *RemoteAptGasPriceOracle, err error) {
	price, low, high, err := sdk.Node().EstimateGasPrice(context.Background())
	if err != nil {
		return
	}
	o = &RemoteAptGasPriceOracle{sdk: sdk, price: price, low: low, high: high}
	go o.update(interval)
	return
}

func (o *RemoteAptGasPriceOracle) update(interval time.Duration) {
	timer := time.NewTicker(interval)
	for range timer.C {
		price, low, high, err := o.sdk.Node().EstimateGasPrice(context.Background())
		if err != nil {
			log.Error("Failed to update gas price", "err", err)
		} else {
			o.Lock()
			o.price = price
			o.low = low
			o.high = high
			o.Unlock()
		}
	}
}

func (o *RemoteAptGasPriceOracle) Price() (uint64, uint64, uint64) {
	o.RLock()
	defer o.RUnlock()
	return o.price, o.low, o.high
}