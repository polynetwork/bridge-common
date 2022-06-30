package eth

import (
	"fmt"
	"testing"
)

func TestClient_HeaderByNumber(t *testing.T) {
	c := New("https://bsc-dataseed1.ninicoin.io")
	height, err := c.GetLatestHeight()
	if err != nil {
		t.Fatal("Failed to get latest height", err)
	}
	t.Log("Chain height", height)
	header, err := c.GetHeader(19123778)
	if err != nil {
		t.Fatal("Failed to get header by number ", err)
	}
	t.Log(fmt.Sprintf("%s", header))
	h, err := header.GetHeight()
	if err != nil { t.Fatal(err) }
	t.Log(h)
}
