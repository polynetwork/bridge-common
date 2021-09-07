/*
 * Copyright (C) 2020 The poly network Authors
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

package types

// query endpoints supported by the auth Querier
const (
	QueryParams        = "params"
	QuerySpan          = "span"
	QuerySpanList      = "span-list"
	QueryLatestSpan    = "latest-span"
	QueryNextSpan      = "next-span"
	QueryNextProducers = "next-producers"
	QueryNextSpanSeed  = "next-span-seed"

	ParamSpan          = "span"
	ParamSprint        = "sprint"
	ParamProducerCount = "producer-count"
	ParamLastEthBlock  = "last-eth-block"
)

// QuerySpanParams defines the params for querying accounts.
type QuerySpanParams struct {
	RecordID uint64
}

// NewQuerySpanParams creates a new instance of QuerySpanParams.
func NewQuerySpanParams(recordID uint64) QuerySpanParams {
	return QuerySpanParams{RecordID: recordID}
}
