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

type CheckTxRequest struct {
	Hash string
}

type CheckTxResponse struct {
	Hash           string
	SrcChainId     uint64
	DstChainId     uint64
	BlockHeight    uint64
	Time           int64
	User           string
	FeeAmount      string
	TransferAmount string
	DstUser        string
	Token          *struct {
		Hash           string
		ChainId        uint64
		Name           string
		TokenBasicName string
	}
	TransactionState []*struct {
		Hash       string
		ChainId    uint64
		Blocks     uint64
		NeedBlocks uint64
		Time       int64
	}
}

func (c *Client) CheckTx(req map[string]*CheckTxRequest) (res *CheckTxResponse, err error) {
	res = new(CheckTxResponse)
	err = tools.PostJsonFor(c.address+"/transactionofhash", req, res)
	if err != nil {
		res = nil
	}
	return
}
