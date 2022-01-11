package gocardless

import (
	"context"
	"testing"
)

func TestCreditorBankAccountCreate(t *testing.T) {
	fixtureFile := "testdata/creditor_bank_accounts.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := CreditorBankAccountCreateParams{}

	o, err :=
		client.CreditorBankAccounts.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected CreditorBankAccount, got nil")

	}
}

func TestCreditorBankAccountList(t *testing.T) {
	fixtureFile := "testdata/creditor_bank_accounts.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := CreditorBankAccountListParams{}

	o, err :=
		client.CreditorBankAccounts.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.CreditorBankAccounts == nil {

		t.Fatalf("Expected list of CreditorBankAccounts, got nil")

	}
}

func TestCreditorBankAccountGet(t *testing.T) {
	fixtureFile := "testdata/creditor_bank_accounts.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.CreditorBankAccounts.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected CreditorBankAccount, got nil")

	}
}

func TestCreditorBankAccountDisable(t *testing.T) {
	fixtureFile := "testdata/creditor_bank_accounts.json"
	server := runServer(t, fixtureFile, "disable")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.CreditorBankAccounts.Disable(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected CreditorBankAccount, got nil")

	}
}
