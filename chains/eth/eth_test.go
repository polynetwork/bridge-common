package eth

import (
	"fmt"
	"testing"
)

func TestClient_HeaderByNumber(t *testing.T) {
	c := New("https://ethereum-mainnet-rpc.allthatnode.com/")
	height, err := c.GetLatestHeight()
	if err != nil {
		t.Fatal("Failed to get latest height", err)
	}
	t.Log("Chain height", height)
	header, err := c.GetHeader(15_050_000)
	if err != nil {
		t.Fatal("Failed to get header by number ", err)
	}
	t.Log(fmt.Sprintf("%s", header))
	t.Log(fmt.Sprintf("%x", header))
	h, err := header.GetHeight()
	if err != nil { t.Fatal(err) }
	t.Log(h)
}
