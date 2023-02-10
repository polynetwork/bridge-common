package aptos

type EventFilter struct {
	Address        string
	CreationNumber string
	Query          map[string]interface{}
}

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
