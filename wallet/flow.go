package wallet
/*
import (
	"context"
	"fmt"
	sdk "github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/polynetwork/bridge-common/chains/flow"
)

var VerifySigAndExecuteTxToLockProxyScripTemplate = `
import LockProxy from %s
import CrossChainManager from %s

transaction(
    userReceiverPublicPath: PublicPath,
    pathStr: String,
    sigs: [String],
    signers: [String],
    toMerkleValueBs: String
) {
    prepare(acct: AuthAccount) {

        let p = LockProxy.getPathFromStr(String.encodeHex(pathStr.utf8))
        if (p == nil) {
            LockProxy.registerReceiverPath(pathStr: pathStr, path: userReceiverPublicPath)
        }
        assert(LockProxy.getPathFromStr(String.encodeHex(pathStr.utf8))!.toString() == userReceiverPublicPath.toString(), 
            message: "fail to regesiter receiver path")

        assert(CrossChainManager.verifySigAndExecuteTx(
            sigs: strArrayToBytesArray(sigs), 
            signers: strArrayToBytesArray(signers), 
            toMerkleValueBs: toMerkleValueBs.decodeHex()
        ), message: "fail to verifySigAndExecuteTx")
    }
}

pub fun strArrayToBytesArray(_ strs: [String]): [[UInt8]] {
    let res: [[UInt8]] = []
    for str in strs {
        res.append(str.decodeHex())
    }
    return res
}
`

type ServiceAccount struct {
	ServiceAcctAddr sdk.Address
	ServiceAcctKey  *sdk.AccountKey
	ServiceSigner   crypto.Signer
}

type FlowWallet struct {
	sdk        *flow.SDK
	Address    string
	PrivateKey string
	config     *Config
}

func NewFlowWallet(config *Config, sdk *flow.SDK) *FlowWallet {
	return &FlowWallet{sdk: sdk, Address: config.Address, PrivateKey: config.PrivateKey, config: config}
}

func (w *FlowWallet) CreateServiceAccount() (serviceAccount *ServiceAccount, err error) {
	privateKey, err := crypto.DecodePrivateKeyHex(crypto.ECDSA_P256, w.PrivateKey)
	if err != nil {
		err = fmt.Errorf("failed to decode flow private key, err=%s", err)
		return
	}

	addr := sdk.HexToAddress(w.Address)
	acc, err := w.sdk.Node().GetAccount(context.Background(), addr)
	if err != nil {
		err = fmt.Errorf("failed get account from flow address, err=%s", err)
		return
	}

	accountKey := acc.Keys[0]
	signer := crypto.NewInMemorySigner(privateKey, accountKey.HashAlgo)

	return &ServiceAccount{
		ServiceAcctAddr: addr,
		ServiceAcctKey:  accountKey,
		ServiceSigner:   signer,
	}, nil
}

func (w *FlowWallet) Init() (err error) {
	return
}
 */
