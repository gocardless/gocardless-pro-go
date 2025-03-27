package gocardless

import (
	"context"
	"testing"
)

func TestBalanceList(t *testing.T) {
	fixtureFile := "testdata/balances.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BalanceListParams{}

	o, err :=
		client.Balances.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.Balances == nil {

		t.Fatalf("Expected list of Balances, got nil")

	}
}
