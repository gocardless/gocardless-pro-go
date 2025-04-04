package gocardless

import (
	"context"
	"testing"
)

func TestBankAccountDetailGet(t *testing.T) {
	fixtureFile := "testdata/bank_account_details.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BankAccountDetailGetParams{}

	o, err :=
		client.BankAccountDetails.Get(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BankAccountDetail, got nil")

	}
}
