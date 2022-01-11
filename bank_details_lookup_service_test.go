package gocardless

import (
	"context"
	"testing"
)

func TestBankDetailsLookupCreate(t *testing.T) {
	fixtureFile := "testdata/bank_details_lookups.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BankDetailsLookupCreateParams{}

	o, err :=
		client.BankDetailsLookups.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BankDetailsLookup, got nil")

	}
}
