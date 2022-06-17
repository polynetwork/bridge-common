package flow
/*
import (
	"crypto/ecdsa"
	"fmt"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ontio/ontology-crypto/ec"
	"github.com/polynetwork/poly/common"
	polytypes "github.com/polynetwork/poly/core/types"
)

type Args struct {
	AssetAddress []byte
	ToAddress    []byte
	Value        uint64
}

func (a *Args) Deserialization(source *common.ZeroCopySource) error {
	assetAddress, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("Args deserialize assetAddress error")
	}
	toAddress, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("Args deserialize toAddress error")
	}
	value, eof := source.NextUint64()
	if eof {
		return fmt.Errorf("Args deserialize value error")
	}
	a.AssetAddress = assetAddress
	a.ToAddress = toAddress
	a.Value = value
	return nil
}

type ResourceRoute struct {
	Address string
	Path    string
}

func (r *ResourceRoute) Deserialization(source *common.ZeroCopySource) error {
	addressUint64, eof := source.NextUint64()
	if eof {
		return fmt.Errorf("ResourceAddress deserialize address error")
	}
	path, eof := source.NextString()
	if eof {
		return fmt.Errorf("ResourceAddress deserialize path error")
	}
	r.Address = fmt.Sprintf("%x", addressUint64)
	r.Path = path
	return nil
}

func RecoverPubkeyFromFlowSig(hash, sig []byte, polyAddr common.Address) (pub []byte, err error) {
	ethSig := make([]byte, ethcrypto.SignatureLength)
	copy(ethSig, sig)

	var pubKey *ecdsa.PublicKey
	for i := 0; i <= 1; i++ {
		ethSig[64] = byte(i)
		pub, err = ethcrypto.Ecrecover(hash[:], ethSig)
		if err != nil {
			err = fmt.Errorf("recover from flow sig failed:%v", err)
			return
		}

		pubKey, err = ethcrypto.UnmarshalPubkey(pub)
		if err != nil {
			err = fmt.Errorf("unmarshal pubkey failed:%v", err)
			return
		}

		addr := polytypes.AddressFromPubKey(&ec.PublicKey{Algorithm: ec.ECDSA, PublicKey: pubKey})
		if polyAddr == addr {
			fmt.Println("polyAddr == addr")
			pub = pub[1:]
			return
		}
	}

	err = fmt.Errorf("didn't found matching pubkey")
	return

}
*/