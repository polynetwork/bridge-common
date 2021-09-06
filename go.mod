module github.com/polynetwork/bridge-common

go 1.15

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/beego/beego/v2 v2.0.1
	github.com/btcsuite/btcd v0.21.0-beta // indirect
	github.com/cespare/cp v1.1.1 // indirect
	github.com/christianxiao/tendermint v0.25.0
	github.com/cosmos/cosmos-sdk v0.39.2
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/ethereum/go-ethereum v1.10.7
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/golang/mock v1.5.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/joeqian10/neo-gogogo v0.0.0-20201214075916-44b70d175579
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/okex/exchain-go-sdk v0.18.2
	github.com/ontio/ontology v1.11.1-0.20200812075204-26cf1fa5dd47
	github.com/ontio/ontology-go-sdk v1.11.4
	github.com/pkg/errors v0.9.1
	github.com/polynetwork/cosmos-poly-module v0.0.0-20200810030259-95d586518759
	github.com/polynetwork/poly v1.3.1
	github.com/polynetwork/poly-go-sdk v0.0.0-20210114035303-84e1615f4ad4
	github.com/polynetwork/poly-io-test v0.0.0-20200819093740-8cf514b07750 // indirect
	github.com/prometheus/tsdb v0.9.1 // indirect
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v1.1.1 // indirect
	github.com/spf13/viper v1.7.1 // indirect
	github.com/status-im/keycard-go v0.0.0-20190424133014-d95853db0f48 // indirect
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.9
	github.com/tendermint/tm-db v0.5.2 // indirect
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de
	gopkg.in/yaml.v2 v2.3.0
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/okex/cosmos-sdk v0.39.2-exchain9
	github.com/ethereum/go-ethereum => github.com/ethereum/go-ethereum v1.9.25
	github.com/okex/exchain => github.com/okex/exchain v0.18.4
	github.com/tendermint/iavl => github.com/okex/iavl v0.14.3-exchain
	github.com/tendermint/tendermint => github.com/okex/tendermint v0.33.9-exchain6
)
