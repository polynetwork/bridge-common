//go:build mainnet
// +build mainnet

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
	NEO      uint64 = 4
	SWITCHEO uint64 = 5
	BSC      uint64 = 6
	HECO     uint64 = 7
	PLT      uint64 = 8
	O3       uint64 = 10
	OK       uint64 = 12
	NEO3     uint64 = 14
	HEIMDALL uint64 = 15
	MATIC    uint64 = 17
	ZILLIQA  uint64 = 18
	ARBITRUM uint64 = 19
	XDAI     uint64 = 20
	AVA      uint64 = 21
	FANTOM   uint64 = 22
	OPTIMISM uint64 = 23
	METIS    uint64 = 24
	BOBA     uint64 = 25
	OASIS    uint64 = 26
	HARMONY  uint64 = 27
	HSC      uint64 = 28
	BYTOM    uint64 = 29
	KCC      uint64 = 30
	STARCOIN uint64 = 31
	KAVA     uint64 = 32
	MILKO    uint64 = 34
	CUBE     uint64 = 35
	CELO     uint64 = 36
	CLOVER   uint64 = 37
	CONFLUX  uint64 = 38
	ASTAR    uint64 = 40
	APTOS    uint64 = 41
	BRISE    uint64 = 42
	DEXIT    uint64 = 43
	CLOUDTX  uint64 = 44
	ZKSYNC   uint64 = 45

	// Invalid chain IDs
	RINKEBY    uint64 = 1000000
	PIXIE      uint64 = 2000000
	BCSPALETTE uint64 = 1001001
	ONTEVM     uint64 = 1001333
	FLOW       uint64 = 1000444
	PLT2       uint64 = 1000108

	// CEX
	BINANCE uint64 = 9001
	OKX     uint64 = 9002
	GATE    uint64 = 9003
	KUCOIN  uint64 = 9004

	// Others
	CORE uint64 = 101
	KLAY uint64 = 102
	CANTO uint64 = 103

	XDC       uint64 = 201
	POLYGONZK uint64 = 202
	FEVM      uint64 = 203

	ENV = "mainnet"
)

var CHAINS = []uint64{
	POLY, ETH, BSC, HECO, OK, ONT, NEO, NEO3, HEIMDALL, MATIC, SWITCHEO, O3, PLT, ARBITRUM, XDAI, OPTIMISM, FANTOM, AVA,
	METIS, BOBA, PIXIE, OASIS, HSC, HARMONY, BYTOM, BCSPALETTE, STARCOIN, ONTEVM, KCC, MILKO, CUBE, KAVA, FLOW, ZKSYNC,
	CELO, CLOVER, CONFLUX, ASTAR, APTOS, BRISE,
}

var ETH_CHAINS = []uint64{
	ETH, BSC, HECO, OK, MATIC, O3, PLT, ARBITRUM, XDAI, OPTIMISM, FANTOM, AVA, METIS, BOBA, PIXIE, OASIS, HSC, HARMONY,
	BYTOM, BCSPALETTE, KCC, ONTEVM, MILKO, CUBE, KAVA, ZKSYNC, CELO, CLOVER, CONFLUX, ASTAR, BRISE, CORE, KLAY, CANTO, XDC,
	DEXIT, CLOUDTX, ZKSYNC, POLYGONZK, FEVM,
}
