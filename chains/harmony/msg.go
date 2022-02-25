package harmony

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type HeaderWithSig struct {
	Header *Header
	HeaderRLP []byte
	Sig []byte
	Bitmap []byte
}

func (hs *HeaderWithSig) Encode() ([]byte, error) {
	hs.Header = nil
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
	Extra               []byte         `json:"extraData"`
	MixDigest           common.Hash    `json:"mixHash"`
	// Additional Fields
	ViewID              *big.Int `json:"viewID"`
	Epoch               *big.Int `json:"epoch"`
	ShardID             uint32   `json:"shardID"`
	LastCommitSignature []byte   `json:"lastCommitSignature"`
	LastCommitBitmap    []byte   `json:"lastCommitBitmap"` // Contains which validator signed
	Vrf                 []byte   `json:"vrf"`
	Vdf                 []byte   `json:"vdf"`
	ShardState          []byte   `json:"shardState"`
}