package zksync

import (
	"testing"
)

func TestClient_GetZkL1BatchNumber(t *testing.T) {
	c := New("https://zksync2-mainnet.zksync.io")
	height, err := c.GetLatestHeight()
	if err != nil {
		t.Fatal("Failed to get latest height", err)
	}
	t.Log("Chain height", height)

	//height = 9157
	number, err := c.GetZkL1BatchNumber(height)
	if err != nil {
		t.Fatal("Failed to GetZkL1BatchNumber", err)
	}
	t.Log("ZkL1BatchNumber", number)

}
