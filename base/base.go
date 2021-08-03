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
	case NEO3:
		return "NEO3"
	default:
		return fmt.Sprintf("Unknown(%d)", id)
	}
}
