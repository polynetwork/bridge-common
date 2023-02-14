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
	"github.com/polynetwork/bridge-common/base"
	"github.com/polynetwork/bridge-common/log"
)

// London upgrade - support
type EthWallet struct {
	*Wallet
}

func (w *EthWallet) Estimate(account accounts.Account, addr common.Address, amount *big.Int, gasLimit uint64, price GasPriceOracle, gasPriceX *big.Float, data []byte) (tip, cap *big.Int, limit uint64, err error) {
	if price == nil {
		tip, err = w.GasTip()
		if err != nil { return }
		cap, err = w.GasPrice()
		if err != nil { return }
	} else {
		cap, tip = price.PriceWithTip()
	}
	base := new(big.Int).Sub(cap, tip)
	base = big.NewInt(0).Quo(big.NewInt(0).Mul(base, big.NewInt(12)), big.NewInt(10)) // max base fee

	if gasPriceX != nil {
		tip, _ = new(big.Float).Mul(new(big.Float).SetInt(tip), gasPriceX).Int(nil)
	}

	cap = big.NewInt(0).Add(base, tip)

	if gasLimit == 0 {
		msg := ethereum.CallMsg{
			From: account.Address, To: &addr, Value: big.NewInt(0), Data: data,
			GasFeeCap: cap, GasTipCap: tip,
		}
		gasLimit, err = w.sdk.Node().EstimateGas(context.Background(), msg)
		if err != nil {
			return
		}
	}

	gasLimit = uint64(1.3 * float32(gasLimit))
	limit = GetChainGasLimit(w.chainId, gasLimit)
	if limit < gasLimit {
		err = fmt.Errorf("Send tx estimated gas limit(%v) higher than max %v", gasLimit, limit)
		return
	}
	return
}

// NOTE: gasPrice, gasPriceX used as gas tip here!
func (w *EthWallet) QuikcSendWithAccount(account accounts.Account, addr common.Address, amount *big.Int, gasLimit uint64, price GasPriceOracle, gasPriceX *big.Float, data []byte) (hash string, err error) {
	tip, cap, limit, err := w.Estimate(account, addr, amount, gasLimit, price, gasPriceX, data)
	if err != nil {
		if strings.Contains(err.Error(), "has been executed") {
			log.Info("Transaction already executed")
			err = nil
			return
		}
		err = fmt.Errorf("Estimate gas limit error %v", err)
		return
	}

	provider, nonces := w.GetAccount(account)
	nonce, err := nonces.Acquire()
	if err != nil {
		return
	}
	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:     nonce,
		GasTipCap: tip,
		GasFeeCap: cap,
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

	if w.Broadcast {
		err = w.sdk.Broadcast(tx)
	} else {
		err = w.sdk.Node().SendTransaction(context.Background(), tx)
	}
	//TODO: Check err here before update nonces
	nonces.Update(err == nil)
	log.Info("Sending tx", "hash", tx.Hash(), "account", account.Address, "nonce", tx.Nonce(), "limit", tx.Gas(), "gasPrice", tx.GasPrice(), "err", err)
	return tx.Hash().String(), err
}

func (w *EthWallet) QuickSendWithMaxLimit(chainId uint64, account accounts.Account, addr common.Address, amount *big.Int, maxLimit *big.Int, price GasPriceOracle, gasPriceX *big.Float, data []byte) (hash string, err error) {
	tip, cap, limit, err := w.Estimate(account, addr, amount, 0, price, gasPriceX, data)
	if err != nil {
		if strings.Contains(err.Error(), "has been executed") {
			log.Info("Transaction already executed")
			err = nil
			return
		}
		err = fmt.Errorf("Estimate gas limit error %v", err)
		return
	}

	if chainId == base.OPTIMISM {
		limit = uint64(1.6 * float32(limit))
	}

	delta := new(big.Int).Sub(maxLimit, new(big.Int).Mul(big.NewInt(int64(limit)), cap))
	if delta.Sign() < 0 {
		err = fmt.Errorf("Send tx estimated gas (limit %v, price %s) higher than max limit %s", limit, cap, maxLimit)
		return
	}

	provider, nonces := w.GetAccount(account)
	nonce, err := nonces.Acquire()
	if err != nil {
		return
	}
	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:     nonce,
		GasTipCap: tip,
		GasFeeCap: cap,
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

	if w.Broadcast {
		err = w.sdk.Broadcast(tx)
	} else {
		err = w.sdk.Node().SendTransaction(context.Background(), tx)
	}
	
	//TODO: Check err here before update nonces
	nonces.Update(err == nil)
	log.Info("Compose dst chain tx", "hash", tx.Hash(), "account", account.Address, "nonce", tx.Nonce(), "limit", tx.Gas(), "gasPrice", tx.GasPrice())
	log.Info("Sent tx with limit", "chainId", chainId, "hash", hash, "delta", delta, "err", err)
	return tx.Hash().String(), err
}