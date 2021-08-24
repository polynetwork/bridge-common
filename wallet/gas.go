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
	"github.com/polynetwork/bridge-common/base"
)

var softGasLimits = map[uint64]uint64{}

func HardGasLimit(chain uint64) uint64 {
	switch chain {
	case base.ETH:
		return 1000000
	default:
		return 1000000000
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
