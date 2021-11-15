module github.com/polynetwork/bridge-common

go 1.15

require (
	github.com/beego/beego/v2 v2.0.1
	github.com/cosmos/cosmos-sdk v0.39.2
	github.com/ethereum/go-ethereum v1.10.7
	github.com/joeqian10/neo-gogogo v0.0.0-20201214075916-44b70d175579
	github.com/mattn/go-colorable v0.1.7
	github.com/mattn/go-isatty v0.0.12
	github.com/okex/exchain-go-sdk v0.18.2
	github.com/ontio/ontology v1.11.1-0.20200812075204-26cf1fa5dd47
	github.com/ontio/ontology-go-sdk v1.11.4
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/polynetwork/cosmos-poly-module v0.0.0-20200810030259-95d586518759
	github.com/polynetwork/poly v1.3.1
	github.com/polynetwork/poly-go-sdk v0.0.0-20210114035303-84e1615f4ad4
	github.com/polynetwork/poly-io-test v0.0.0-20200819093740-8cf514b07750 // indirect
	github.com/tendermint/tendermint v0.33.9
	golang.org/x/mod v0.5.0 // indirect
	golang.org/x/sys v0.0.0-20210915083310-ed5796bab164 // indirect
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/okex/cosmos-sdk v0.39.2-exchain9
	github.com/ethereum/go-ethereum => github.com/ethereum/go-ethereum v1.9.25
	github.com/okex/exchain => github.com/okex/exchain v0.18.4
	github.com/tendermint/iavl => github.com/okex/iavl v0.14.3-exchain
	github.com/tendermint/tendermint => github.com/okex/tendermint v0.33.9-exchain6
)
