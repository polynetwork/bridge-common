package wallet

import (
	"context"
	"testing"
	"time"

	"github.com/polynetwork/bridge-common/chains/apt"
	"github.com/portto/aptos-go-sdk/models"
	"github.com/stretchr/testify/assert"
)

func TestApt(t *testing.T) {
	sdk, err := apt.WithOptions(0, []string{"https://testnet.aptoslabs.com"}, time.Minute, 1000)
	assert.NoError(t, err)
	w := NewAptWallet(&Config{
		Path: "../.apt.dat",
	}, sdk)
	from, _ := models.HexToAccountAddress("0xcda3c83ef5376950794a0260d6be57a267c9325b356b47d1390b485a50ab61d5")
	to, _ := models.HexToAccountAddress("0x0e69e1d1069f086aca14daccbd3183848a1a446f5c3d3ea09bfa964e9324798c")
	
	/*
	hash, err := w.CreateAccount(context.Background(), &from, to, time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("hash %s", hash)
	*/

	hash, err := w.Transfer(context.Background(), from, to, 1, time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("hash %s", hash)
}
