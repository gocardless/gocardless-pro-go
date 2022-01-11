package gocardless

import (
	"context"
	"testing"
)

func TestBankAuthorisationGet(t *testing.T) {
	fixtureFile := "testdata/bank_authorisations.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.BankAuthorisations.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BankAuthorisation, got nil")

	}
}

func TestBankAuthorisationCreate(t *testing.T) {
	fixtureFile := "testdata/bank_authorisations.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BankAuthorisationCreateParams{}

	o, err :=
		client.BankAuthorisations.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BankAuthorisation, got nil")

	}
}
