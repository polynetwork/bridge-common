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
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/polynetwork/bridge-common/log"
)

// London upgrade - support
type EthWallet struct {
	Wallet
}

// NOTE: gasPrice, gasPriceX used as gas tip here!
func (w *EthWallet) SendWithAccount(account accounts.Account, addr common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, gasPriceX *big.Float, data []byte) (hash string, err error) {
	if gasPrice == nil || gasPrice.Sign() <= 0 {
		gasPrice, err = w.GasTip()
		if err != nil {
			err = fmt.Errorf("Get gas tip error %v", err)
			return
		}
		if gasPriceX != nil {
			gasPrice, _ = new(big.Float).Mul(new(big.Float).SetInt(gasPrice), gasPriceX).Int(nil)
		}
	}

	gasCap, err := w.GasPrice()
	if err != nil {
		err = fmt.Errorf("Get gas price error %v", err)
		return
	}
	// TODO: Make this configurable
	gasCap = big.NewInt(0).Quo(big.NewInt(0).Mul(gasCap, big.NewInt(30)), big.NewInt(10)) // max gas price

	provider, nonces := w.GetAccount(account)
	nonce, err := nonces.Acquire()
	if err != nil {
		return
	}
	if gasLimit == 0 {
		msg := ethereum.CallMsg{
			From: account.Address, To: &addr, Value: big.NewInt(0), Data: data,
			GasFeeCap: gasCap, GasTipCap: gasPrice,
		}
		gasLimit, err = w.sdk.Node().EstimateGas(context.Background(), msg)
		if err != nil {
			nonces.Update(false)
			if strings.Contains(err.Error(), "has been executed") {
				log.Info("Transaction already executed")
				return "", nil
			}
			return "", fmt.Errorf("Estimate gas limit error %v", err)
		}
	}

	gasLimit = uint64(1.3 * float32(gasLimit))
	limit := GetChainGasLimit(w.chainId, gasLimit)
	if limit < gasLimit {
		nonces.Update(false)
		return "", fmt.Errorf("Send tx estimated gas limit(%v) higher than max %v", gasLimit, limit)
	}
	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:     nonce,
		GasTipCap: gasPrice,
		GasFeeCap: gasCap,
		Gas:       limit,
		To:        &addr,
		Value:     amount,
		Data:      data,
	})
	// tx := types.NewTransaction(nonce, addr, amount, limit, gasPrice, data)
	tx, err = provider.SignTx(account, tx, big.NewInt(int64(w.chainId)))
	if err != nil {
		nonces.Update(false)
		return "", fmt.Errorf("Sign tx error %v", err)
	}
	log.Info("Compose dst chain tx", "hash", tx.Hash(), "account", account.Address)
	err = w.sdk.Node().SendTransaction(context.Background(), tx)
	//TODO: Check err here before update nonces
	nonces.Update(true)
	return tx.Hash().String(), err
}
