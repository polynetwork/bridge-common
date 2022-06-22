module github.com/polynetwork/bridge-common

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.39.2
	github.com/ethereum/go-ethereum v1.10.8
	github.com/joeqian10/neo-gogogo v0.0.0-20201214075916-44b70d175579
	github.com/mattn/go-colorable v0.1.8
	github.com/mattn/go-isatty v0.0.12
	github.com/okex/exchain v1.5.6
	github.com/okex/exchain-go-sdk v1.5.6
	github.com/ontio/ontology v1.11.1-0.20200812075204-26cf1fa5dd47
	github.com/ontio/ontology-go-sdk v1.11.4
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/polynetwork/cosmos-poly-module v0.0.0-20200810030259-95d586518759
	github.com/polynetwork/poly v1.3.1
	github.com/polynetwork/poly-go-sdk v0.0.0-20210114035303-84e1615f4ad4
	github.com/polynetwork/poly-io-test v0.0.0-20200819093740-8cf514b07750 // indirect
	github.com/tendermint/tendermint v0.33.9
	golang.org/x/mod v0.5.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)

replace (
	github.com/ethereum/go-ethereum => github.com/okex/go-ethereum v1.10.8-oec3
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tendermint/go-amino => github.com/okex/go-amino v0.15.1-exchain7
)
