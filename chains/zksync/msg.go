package zksync

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Header []byte

type HeaderInfo struct {
	ParentHash    common.Hash  `json:"parentHash"       gencodec:"required"`
	Root          common.Hash  `json:"stateRoot"        gencodec:"required"`
	TxHash        common.Hash  `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash   common.Hash  `json:"receiptsRoot"     gencodec:"required"`
	Number        *hexutil.Big `json:"number"           gencodec:"required"`
	L1BatchNumber *hexutil.Big `json:"L1BatchNumber"           gencodec:"required"`
}

func (h Header) L1BatchNumber() (number uint64, err error) {
	info := new(HeaderInfo)
	err = json.Unmarshal(h, info)
	if err != nil {
		return
	}

	if info.L1BatchNumber != nil {
		number = info.L1BatchNumber.ToInt().Uint64()
	}
	return
}
