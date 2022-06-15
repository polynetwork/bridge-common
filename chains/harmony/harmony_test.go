package harmony

import (
	"context"
	"fmt"
	"testing"

	"github.com/polynetwork/bridge-common/chains/custom"
)

func TestNew(t *testing.T) {
	c, _ := custom.Dial("https://zksync2-testnet.zksync.dev/")
	resp, err := c.CallContextRaw(
		context.Background(), "zks_getMainContract")
		// custom.ToBlockNumArg(big.NewInt(int64(1365587))), true)
	if err != nil { t.Fatal(err) }
	t.Log(string(resp.Result))
}

func TestClient_HeaderByNumber(t *testing.T) {
	c := New("https://api.harmony.one")
	height, err := c.GetLatestHeight()
	if err != nil {
		t.Fatal("Failed to get latest height", err)
	}
	t.Log("Chain height", height)
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
