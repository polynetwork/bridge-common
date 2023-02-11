package bridge

import (
	"fmt"
	"testing"

	"github.com/polynetwork/bridge-common/util"
)

func TestTokens(t *testing.T) {
	b := New("https://bridge.poly.network/v1")
	res, err := b.Tokens(2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(util.Json(res))
}

func TestAnyTokens(t *testing.T) {
	b := New("https://bridgeapi.multichain.org/v4")
	res, err := b.List(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(util.Json(res))
}
