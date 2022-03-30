//go:build devnet
// +build devnet

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
	POLY     uint64 = 0
	ZION     uint64 = 1
	ETH      uint64 = 202
	KOVAN    uint64 = 302
	RINKBY   uint64 = 402
	GOERLI   uint64 = 502
	ONT      uint64 = 3
	NEO      uint64 = 5
	HECO     uint64 = 207
	BSC      uint64 = 1000006
	SIDE     uint64 = 77
	O3       uint64 = 82
	NEO3     uint64 = 88
	PLT      uint64 = 208
	ZILLIQA  uint64 = 111
	OK       uint64 = 200
	HEIMDALL uint64 = 2015
	MATIC    uint64 = 20016
	ARBITRUM uint64 = 300
	XDAI     uint64 = 600
	OPTIMISM uint64 = 1200
	FANTOM   uint64 = 400
	AVA      uint64 = 500
	METIS    uint64 = 301
	SWITCHEO uint64 = 1000

	ENV = "devnet"
)

var CHAINS = []uint64{
	POLY, ETH, BSC, HECO, OK, ONT, NEO, NEO3, HEIMDALL, MATIC, SWITCHEO, O3, PLT, ARBITRUM, XDAI, OPTIMISM, FANTOM, AVA, KOVAN, RINKBY, GOERLI, SIDE, METIS,
}
