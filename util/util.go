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

package util

func Min(nums ...uint64) (min uint64) {
	if len(nums) > 0 {
		min = nums[0]
	} else {
		return
	}
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	return
}

func NonZeroMin(nums ...uint64) (min uint64) {
	for _, num := range nums {
		if min == 0 || (min > num && num != 0) {
			min = num
		}
	}
	return
}
