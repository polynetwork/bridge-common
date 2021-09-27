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

package bridge

import "github.com/polynetwork/bridge-common/tools"

type CheckFeeStatus int

const (
	SKIP     CheckFeeStatus = -2 // Skip since not our tx
	NOT_PAID CheckFeeStatus = -1 // Not paid or paid too low
	MISSING  CheckFeeStatus = 0  // Tx not received yet
	PAID     CheckFeeStatus = 1  // Paid and enough pass
)

type CheckFeeRequest struct {
	ChainId  uint64
	TxId     string
	PolyHash string
	Paid     float32
	Min      float64
	Status   CheckFeeStatus
}

func (r *CheckFeeRequest) Pass() bool {
	return r != nil && r.Status == PAID
}

func (r *CheckFeeRequest) Skip() bool {
	return r != nil && r.Status == SKIP
}

func (r *CheckFeeRequest) Missing() bool {
	return r != nil || r.Status == MISSING
}

func (c *Client) CheckFee(req map[string]*CheckFeeRequest) (err error) {
	return tools.PostJsonFor(c.address, req, &req)
}
