//go:build testnet
// +build testnet

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

const (
	POLY     uint64 = 0
	BTC      uint64 = 1
	ETH      uint64 = 2
	ONT      uint64 = 3
	NEO      uint64 = 5
	HECO     uint64 = 7
	BSC      uint64 = 79
	O3       uint64 = 82
	NEO3     uint64 = 88
	PLT      uint64 = 107
	ZILLIQA  uint64 = 111
	OK       uint64 = 200
	HEIMDALL uint64 = 201
	MATIC    uint64 = 202
	ARBITRUM uint64 = 205
	XDAI     uint64 = 206
	OPTIMISM uint64 = 210
	FANTOM   uint64 = 208
	AVA      uint64 = 209
	METIS    uint64 = 300
	PIXIE    uint64 = 316
	BOBA     uint64 = 400
	RINKEBY  uint64 = 402
	OASIS    uint64 = 500
	HSC      uint64 = 603
	BYTOM    uint64 = 701
	HARMONY  uint64 = 800

	SWITCHEO   uint64 = 1000
	BCSPALETTE uint64 = 1001

	ENV = "testnet"
)

var CHAINS = []uint64{
	POLY, ETH, BSC, HECO, OK, ONT, NEO, NEO3, HEIMDALL, MATIC, SWITCHEO, O3, PLT, ARBITRUM, XDAI, OPTIMISM, FANTOM, AVA, METIS, RINKEBY, BOBA, PIXIE, OASIS, HSC, HARMONY, BYTOM, BCSPALETTE,
}

var ETH_CHAINS = []uint64{
	ETH, BSC, HECO, OK, MATIC, O3, PLT, ARBITRUM, XDAI, OPTIMISM, FANTOM, AVA, METIS, RINKEBY, BOBA, PIXIE, OASIS, HSC, HARMONY, HARMONY, BYTOM, BCSPALETTE,
}
