package gocardless

import (
	"context"
	"testing"
)

func TestPaymentAccountTransactionList(t *testing.T) {
	fixtureFile := "testdata/payment_account_transactions.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := PaymentAccountTransactionListParams{}

	o, err :=
		client.PaymentAccountTransactions.List(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o.PaymentAccountTransactions == nil {

		t.Fatalf("Expected list of PaymentAccountTransactions, got nil")

	}
}
