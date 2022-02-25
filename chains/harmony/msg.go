package harmony

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type HeaderWithSig struct {
	HeaderRLP []byte
	Sig []byte
	Bitmap []byte
}

func (hs *HeaderWithSig) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(hs)
}

type Header struct {
	ParentHash          common.Hash    `json:"parentHash"`
	Coinbase            common.Address `json:"miner"`
	Root                common.Hash    `json:"stateRoot"`
	TxHash              common.Hash    `json:"transactionsRoot"`
	ReceiptHash         common.Hash    `json:"receiptsRoot"`
	OutgoingReceiptHash common.Hash    `json:"outgoingReceiptsRoot"`
	IncomingReceiptHash common.Hash    `json:"incomingReceiptsRoot"`
	Number              *big.Int       `json:"number"`
	GasLimit            uint64         `json:"gasLimit"`
	GasUsed             uint64         `json:"gasUsed"`
	Extra               string   	   `json:"extraData"`
	MixDigest           common.Hash    `json:"mixHash"`
	// Additional Fields
	ViewID              *big.Int `json:"viewID"`
	Epoch               *big.Int `json:"epoch"`
	ShardID             uint32   `json:"shardID"`
	LastCommitSignature string   `json:"lastCommitSignature"`
	LastCommitSig 		string   `json:"lastCommitSig"`
	LastCommitBitmap    string   `json:"lastCommitBitmap"` // Contains which validator signed
	Vrf                 string   `json:"vrf"`
	Vdf                	string   `json:"vdf"`
	ShardState          string   `json:"shardState"`
}

func (h *Header) GetLastCommitSignature() ([]byte, error) {
	sig := h.LastCommitSignature
	if len(sig)  == 0 {
		sig = h.LastCommitSig
	}
	return hex.DecodeString(sig)
}

func (h *Header) GetLastCommitBitmap() ([]byte, error) {
	return hex.DecodeString(h.LastCommitBitmap)
}