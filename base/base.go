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

package base

import "fmt"

var (
	PRICE_PRECISION = int64(100000000)
	FEE_PRECISION   = int64(100000000)
)

var (
	MARKET_COINMARKETCAP = "coinmarketcap"
	MARKET_BINANCE       = "binance"
	MARKET_HUOBI         = "huobi"
	MARKET_SELF          = "self"
)

const (
	STATE_FINISHED = iota
	STATE_PENDDING
	STATE_SOURCE_DONE
	STATE_SOURCE_CONFIRMED
	STATE_POLY_CONFIRMED
	STATE_DESTINATION_DONE

	STATE_WAIT = 100
	STATE_SKIP = 101
)

func GetStateName(state int) string {
	switch state {
	case STATE_FINISHED:
		return "Finished"
	case STATE_PENDDING:
		return "Pending"
	case STATE_SOURCE_DONE:
		return "SrcDone"
	case STATE_SOURCE_CONFIRMED:
		return "SrcConfirmed"
	case STATE_POLY_CONFIRMED:
		return "PolyConfirmed"
	case STATE_DESTINATION_DONE:
		return "DestDone"
	case STATE_WAIT:
		return "WAIT"
	case STATE_SKIP:
		return "SKIP"
	default:
		return fmt.Sprintf("Unknown(%d)", state)
	}
}

func GetChainName(id uint64) string {
	switch id {
	case POLY:
		return "Poly"
	case ETH:
		return "Ethereum"
	case RINKEBY:
		return "Ethereum-Rinkeby"
	case ONT:
		return "Ontology"
	case NEO:
		return "Neo"
	case BSC:
		return "Bsc"
	case HECO:
		return "Heco"
	case O3:
		return "O3"
	case OK:
		return "OK"
	case MATIC:
		return "Polygon"
	case HEIMDALL:
		return "Heimdall"
	case NEO3:
		return "NEO3"
	case SWITCHEO:
		return "Switcheo"
	case PLT:
		return "Palette"
	case ARBITRUM:
		return "Arbitrum"
	case ZILLIQA:
		return "Zilliqa"
	case XDAI:
		return "Xdai"
	case OPTIMISM:
		return "Optimism"
	case FANTOM:
		return "Fantom"
	case METIS:
		return "Metis"
	case AVA:
		return "Avalanche"
	case BOBA:
		return "Boba"
	case OASIS:
		return "Oasis"
	default:
		return fmt.Sprintf("Unknown(%d)", id)
	}
}

func BlocksToSkip(chainId uint64) uint64 {
	switch chainId {
	case MATIC:
		return 120
	case ETH:
		return 8
	case BSC, HECO:
		return 17
	case O3:
		return 8
	case PLT:
		return 5
	case ONT:
		return 0
	default:
		return 1
	}
}

func BlocksToWait(chainId uint64) uint64 {
	switch chainId {
	case ETH:
		return 12
	case BSC, HECO:
		return 21
	case ONT, NEO, NEO3, OK, SWITCHEO:
		return 1
	case PLT:
		return 4
	case O3:
		return 12
	case MATIC:
		return 128
	default:
		return 100000000
	}
}

func SameAsETH(chainId uint64) bool {
	for _, chain := range ETH_CHAINS {
		if chain == chainId {
			return true
		}
	}
	return false
}
