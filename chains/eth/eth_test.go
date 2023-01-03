package eth

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/polynetwork/bridge-common/util"
)

func TestEth(t *testing.T) {
	key := crypto.Keccak256(common.LeftPadBytes([]byte{0}, 32), common.LeftPadBytes([]byte{10}, 32))
	c := New("https://rpc.ankr.com/eth")
	proofKey := hexutil.Encode(key)
	proof, err := c.GetProof("0xcf2afe102057ba5c16f899271045a0a37fcb10f2", proofKey, 0)
	if err != nil { t.Fatal(err) }
	ap, sp := proof.AccountProof, proof.StorageProofs
	proof.AccountProof, proof.StorageProofs = nil, nil
	t.Logf("%+v",util.Verbose(proof))
	t.Log("Account proof")
	for i, p := range ap {
		t.Logf("%v: %v", i, p)
	}

	t.Log("Storage proof")
	for i, p := range sp {
		t.Logf("%v: %v", i, util.Verbose(p))
	}

	bytes, err := c.StorageAt(context.Background(), common.HexToAddress("0xcf2afe102057ba5c16f899271045a0a37fcb10f2"), common.HexToHash(proofKey), nil)
	if err != nil { t.Fatal(err) }
	t.Logf("Slot %x", bytes)
}