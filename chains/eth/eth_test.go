package eth

import (
	"testing"
)


func TestWs(t *testing.T) {
	sdk := new(Client)
	err := sdk.Listen("")
	if err != nil {
		t.Fatal(err)
	}
	ch := make(chan uint64)
	sdk.Subscribe(ch)
	for h := range ch {
		t.Log(h)
	}
}