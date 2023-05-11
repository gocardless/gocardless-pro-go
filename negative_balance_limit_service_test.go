package gocardless

import (
	"context"
	"testing"
)

func TestNegativeBalanceLimitList(t *testing.T) {
	fixtureFile := "testdata/negative_balance_limits.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := NegativeBalanceLimitListParams{}

	o, err :=
		client.NegativeBalanceLimits.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.NegativeBalanceLimits == nil {

		t.Fatalf("Expected list of NegativeBalanceLimits, got nil")

	}
}

func TestNegativeBalanceLimitCreate(t *testing.T) {
	fixtureFile := "testdata/negative_balance_limits.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := NegativeBalanceLimitCreateParams{}

	o, err :=
		client.NegativeBalanceLimits.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected NegativeBalanceLimit, got nil")

	}
}
