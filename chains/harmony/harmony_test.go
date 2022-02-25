package harmony

import (
	"fmt"
	"testing"
)

func TestClient_HeaderByNumber(t *testing.T) {
	c := New("https://api.harmony.one")
	header, err := c.HeaderByNumber(23361200)
	if err != nil {
		t.Fatal("Failed to header by number ", err)
	}
	t.Log(fmt.Sprintf("%+v", *header))
	sig, err := header.GetLastCommitSignature()
	if len(sig) == 0 {
		t.Fatal("Failed to decode last commit sig", err)
	}
	bitmap, err := header.GetLastCommitBitmap()
	if len(sig) == 0 {
		t.Fatal("Failed to decode last commit bitmap", err)
	}
	t.Log(sig, bitmap)
}
