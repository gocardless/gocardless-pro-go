package gocardless

import (
	"context"
	"testing"
)

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
