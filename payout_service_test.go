package gocardless

import (
	"context"
	"testing"
)

func TestPayoutList(t *testing.T) {
	fixtureFile := "testdata/payouts.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := PayoutListParams{}

	o, err :=
		client.Payouts.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.Payouts == nil {

		t.Fatalf("Expected list of Payouts, got nil")

	}
}

func TestPayoutGet(t *testing.T) {
	fixtureFile := "testdata/payouts.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.Payouts.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Payout, got nil")

	}
}

func TestPayoutUpdate(t *testing.T) {
	fixtureFile := "testdata/payouts.json"
	server := runServer(t, fixtureFile, "update")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := PayoutUpdateParams{}

	o, err :=
		client.Payouts.Update(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Payout, got nil")

	}
}
