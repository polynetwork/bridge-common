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
	POLY  uint64 = 0
	BTC   uint64 = 1
	ETH   uint64 = 2
	ONT   uint64 = 3
	NEO   uint64 = 4
	BSC   uint64 = 6
	HECO  uint64 = 7
	O3    uint64 = 80
	NEO3  uint64 = 88
	OK    uint64 = 90
	MATIC uint64 = 13
	METIS uint64 = 300

	SWITCHEO uint64 = 1000

	ENV = "devnet"
)

var CHAINS = []uint64{POLY, ETH, ONT, NEO, BSC, HECO, O3, OK, MATIC, METIS}
