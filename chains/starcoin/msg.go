package starcoin

import (
	"fmt"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/bcs"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
)

type CrossChainEvent struct {
	Sender               []byte
	TxId                 []byte
	ProxyOrAssetContract []byte
	ToChainId            uint64
	ToContract           []byte
	RawData              []byte
}

type StarcoinToPolyHeaderOrCrossChainMsg struct {
	EventIndex *int    `json:"event_index,omitempty"`
	AccessPath *string `json:"access_path,omitempty"`
}

func DeserializeCrossChainEvent(input []byte) (CrossChainEvent, error) {
	if input == nil {
		var obj CrossChainEvent
		return obj, fmt.Errorf("cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := deserializeCrossChainEvent(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("some input bytes were not read")
	}
	return obj, err
}

func deserializeCrossChainEvent(deserializer serde.Deserializer) (CrossChainEvent, error) {
	var obj CrossChainEvent
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.Sender = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.TxId = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.ProxyOrAssetContract = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.ToChainId = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.ToContract = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.RawData = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}
