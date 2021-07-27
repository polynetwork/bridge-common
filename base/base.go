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
)

const (
	ADDRESS_LENGTH = 64
)
