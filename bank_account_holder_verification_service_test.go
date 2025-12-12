package gocardless

import (
	"context"
	"testing"
)

func TestBankAccountHolderVerificationCreate(t *testing.T) {
	fixtureFile := "testdata/bank_account_holder_verifications.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BankAccountHolderVerificationCreateParams{}

	o, err :=
		client.BankAccountHolderVerifications.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BankAccountHolderVerification, got nil")

	}
}

func TestBankAccountHolderVerificationGet(t *testing.T) {
	fixtureFile := "testdata/bank_account_holder_verifications.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.BankAccountHolderVerifications.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BankAccountHolderVerification, got nil")

	}
}
