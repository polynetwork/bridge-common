package wallet

import (
	"github.com/polynetwork/bridge-common/chains/aptos"
)

type AptosWallet struct {
	sdk        *aptos.SDK
	Address    string
	PrivateKey string
	PublicKey  string
	config     *Config
}

func NewAptosWallet(config *Config, sdk *aptos.SDK) *AptosWallet {
	return &AptosWallet{sdk: sdk, Address: config.Address, PrivateKey: config.PrivateKey, PublicKey: config.PublicKey, config: config}
}

func (w *AptosWallet) Init() (err error) {
	return
}
