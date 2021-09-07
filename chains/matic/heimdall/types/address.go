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

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/polynetwork/bridge-common/chains/matic/cosmos"
	"gopkg.in/yaml.v2"
)

const (
	// AddrLen defines a valid address length
	AddrLen = 20
)

// Ensure that different address types implement the interface
var _ cosmos.Address = HeimdallAddress{}
var _ yaml.Marshaler = HeimdallAddress{}

// HeimdallAddress represents heimdall address
type HeimdallAddress common.Address

// ZeroHeimdallAddress represents zero address
var ZeroHeimdallAddress = HeimdallAddress{}

// EthAddress get eth address
func (aa HeimdallAddress) EthAddress() common.Address {
	return common.Address(aa)
}

// Equals returns boolean for whether two AccAddresses are Equal
func (aa HeimdallAddress) Equals(aa2 cosmos.Address) bool {
	if aa.Empty() && aa2.Empty() {
		return true
	}

	return bytes.Equal(aa.Bytes(), aa2.Bytes())
}

// Empty returns boolean for whether an AccAddress is empty
func (aa HeimdallAddress) Empty() bool {
	return bytes.Equal(aa.Bytes(), ZeroHeimdallAddress.Bytes())
}

// Marshal returns the raw address bytes. It is needed for protobuf
// compatibility.
func (aa HeimdallAddress) Marshal() ([]byte, error) {
	return aa.Bytes(), nil
}

// Unmarshal sets the address to the given data. It is needed for protobuf
// compatibility.
func (aa *HeimdallAddress) Unmarshal(data []byte) error {
	*aa = HeimdallAddress(common.BytesToAddress(data))
	return nil
}

// MarshalJSON marshals to JSON using Bech32.
func (aa HeimdallAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(aa.String())
}

// MarshalYAML marshals to YAML using Bech32.
func (aa HeimdallAddress) MarshalYAML() (interface{}, error) {
	return aa.String(), nil
}

// UnmarshalJSON unmarshals from JSON assuming Bech32 encoding.
func (aa *HeimdallAddress) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	*aa = HexToHeimdallAddress(s)
	return nil
}

// UnmarshalYAML unmarshals from JSON assuming Bech32 encoding.
func (aa *HeimdallAddress) UnmarshalYAML(data []byte) error {
	var s string
	err := yaml.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	*aa = HexToHeimdallAddress(s)
	return nil
}

// Bytes returns the raw address bytes.
func (aa HeimdallAddress) Bytes() []byte {
	return aa[:]
}

// String implements the Stringer interface.
func (aa HeimdallAddress) String() string {
	return "0x" + hex.EncodeToString(aa.Bytes())
}

// Format implements the fmt.Formatter interface.
// nolint: errcheck
func (aa HeimdallAddress) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		s.Write([]byte(aa.String()))
	case 'p':
		s.Write([]byte(fmt.Sprintf("%p", aa)))
	default:
		s.Write([]byte(fmt.Sprintf("%X", aa.Bytes())))
	}
}

//
// Address utils
//

// BytesToHeimdallAddress returns Address with value b.
func BytesToHeimdallAddress(b []byte) HeimdallAddress {
	return HeimdallAddress(common.BytesToAddress(b))
}

// HexToHeimdallAddress returns Address with value b.
func HexToHeimdallAddress(b string) HeimdallAddress {
	return HeimdallAddress(common.HexToAddress(b))
}

// AccAddressToHeimdallAddress returns Address with value b.
func AccAddressToHeimdallAddress(b cosmos.AccAddress) HeimdallAddress {
	return BytesToHeimdallAddress(b[:])
}

// HeimdallAddressToAccAddress returns Address with value b.
func HeimdallAddressToAccAddress(b HeimdallAddress) cosmos.AccAddress {
	return cosmos.AccAddress(b.Bytes())
}

// SampleHeimdallAddress returns sample address
func SampleHeimdallAddress(s string) HeimdallAddress {
	return BytesToHeimdallAddress([]byte(s))
}
