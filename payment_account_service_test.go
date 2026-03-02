package gocardless

import (
	"context"
	"testing"
)

func TestPaymentAccountGet(t *testing.T) {
	fixtureFile := "testdata/payment_accounts.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.PaymentAccounts.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected PaymentAccount, got nil")

	}
}

func TestPaymentAccountList(t *testing.T) {
	fixtureFile := "testdata/payment_accounts.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := PaymentAccountListParams{}

	o, err :=
		client.PaymentAccounts.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.PaymentAccounts == nil {

		t.Fatalf("Expected list of PaymentAccounts, got nil")

	}
}
