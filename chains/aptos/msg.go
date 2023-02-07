package aptos

type CrossChainGlobalConfig struct {
	Type string
	Data struct {
		AptosToPolyTxHashIndex string
		CurEpochEndHeight      string
		CurEpochStartHeight    string
		CurValidators          []string
		PolyId                 string
	}
}
