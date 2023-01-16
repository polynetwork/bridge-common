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

const (
	POLY       uint64 = 0
	ZION       uint64 = POLY
	ZIONCHAIN  uint64 = 1
	ETH        uint64 = 2
	ONT        uint64 = 3
	NEO        uint64 = 4
	BSC        uint64 = 6
	HECO       uint64 = 7
	O3         uint64 = 80
	NEO3       uint64 = 888
	OK         uint64 = 90
	MATIC      uint64 = 13
	METIS      uint64 = 300
	PIXIE      uint64 = 316
	RINKEBY    uint64 = 402
	HSC        uint64 = 603
	BYTOM      uint64 = 701
	KCC        uint64 = 900
	ONTEVM     uint64 = 5555
	FLOW       uint64 = 444
	KAVA       uint64 = 920
	CUBE       uint64 = 930
	SWITCHEO   uint64 = 1000
	HARMONY    uint64 = 801
	BCSPALETTE uint64 = 1001
	STARCOIN   uint64 = 318
	ZKSYNC     uint64 = 940
	CELO       uint64 = 960
	CLOVER     uint64 = 970
	CONFLUX    uint64 = 980
	ASTAR      uint64 = 990
	APTOS      uint64 = 998
	BRISE      uint64 = 1010

	// Invalid chain ids
	HEIMDALL uint64 = 9000001
	PLT      uint64 = 9000002
	ARBITRUM uint64 = 9000003
	ZILLIQA  uint64 = 9000004
	XDAI     uint64 = 9000005
	OPTIMISM uint64 = 9000006
	FANTOM   uint64 = 9000007
	AVA      uint64 = 9000008
	BOBA     uint64 = 9000009
	OASIS    uint64 = 9000010
	MILKO    uint64 = 9000011
	KOVAN    uint64 = 9000012
	GOERLI   uint64 = 9000013

	ENV = "devnet"
)

var CHAINS = []uint64{
	ZION, ZIONCHAIN, ETH, GOERLI, ONT, NEO, BSC, HECO, O3, OK, MATIC, METIS, RINKEBY, PIXIE, HSC, HARMONY, BYTOM, STARCOIN,
	ONTEVM, CUBE, KAVA, FLOW, ZKSYNC, CELO, CLOVER, CONFLUX, ASTAR, APTOS, BRISE, RIPPLE,
}

var ETH_CHAINS = []uint64{
	ZION, ZIONCHAIN, ETH, GOERLI, BSC, HECO, OK, MATIC, O3, METIS, RINKEBY, PIXIE, HSC, HARMONY, BYTOM, KCC, ONTEVM, CUBE, KAVA,
	ZKSYNC, CELO, CLOVER, CONFLUX, ASTAR, BRISE,
}
