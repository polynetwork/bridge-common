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
	"fmt"

	sdk "github.com/polynetwork/poly-go-sdk"
)

type PolySignerConfig struct {
	Path     string
	Password string
}

func NewPolySigner(config *PolySignerConfig) (signer *sdk.Account, err error) {
	if config == nil {
		return nil, fmt.Errorf("Missing poly wallet config")
	}
	s := sdk.NewPolySdk()
	wallet, err := s.OpenWallet(config.Path)
	if err != nil {
		err = fmt.Errorf("Open poly wallet error %v", err)
		return nil, err
	}
	signer, err = wallet.GetDefaultAccount([]byte(config.Password))
	if err != nil {
		err = fmt.Errorf("Get poly default account error %v", err)
		return
	}
	return
}
