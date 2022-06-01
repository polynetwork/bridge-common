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

import (
	"github.com/polynetwork/bridge-common/tools"
)

type CheckFeeStatus int

const (
	SKIP        CheckFeeStatus = -2 // Skip since not our tx
	NOT_PAID    CheckFeeStatus = -1 // Not paid or paid too low
	MISSING     CheckFeeStatus = 0  // Tx not received yet
	PAID        CheckFeeStatus = 1  // Paid and enough pass
	EstimatePay CheckFeeStatus = 2  // Paid but need EstimateGas
)

type CheckFeeRequest struct {
	ChainId  uint64
	TxId     string
	PolyHash string
	Paid     float64
	Min      float64
	PaidGas  float64
	Status   CheckFeeStatus
}

func (r *CheckFeeRequest) Pass() bool {
	return r != nil && r.Status == PAID
}

func (r *CheckFeeRequest) Skip() bool {
	return r != nil && r.Status == SKIP
}

func (r *CheckFeeRequest) Missing() bool {
	return r == nil || r.Status == MISSING
}

func (r *CheckFeeRequest) EstimatePay() bool {
	return r != nil && r.Status == EstimatePay
}

func (c *Client) CheckFee(req map[string]*CheckFeeRequest) (err error) {
	return tools.PostJsonFor(c.address+"/newcheckfee", req, &req)
}

type GetFeeRequest struct {
	SrcChainId    uint64
	DstChainId    uint64
	SwapTokenHash string
	Hash          string
}

type GetFeeResponse struct {
	SrcChainId               uint64
	DstChainId               uint64
	SwapTokenHash            string
	TokenAmount              string
	UsdtAmount               string
	TokenAmountWithPrecision string
	Balance                  string
	BalanceWithPrecision     string
}

func (c *Client) GetFee(req *GetFeeRequest) (res *GetFeeResponse, err error) {
	res = new(GetFeeResponse)
	err = tools.PostJsonFor(c.address+"/getfee", req, res)
	if err != nil {
		return nil, err
	}
	return
}
