package gocardless

import (
	"context"
	"testing"
)

func TestCustomerBankAccountCreate(t *testing.T) {
	fixtureFile := "testdata/customer_bank_accounts.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := CustomerBankAccountCreateParams{}

	o, err :=
		client.CustomerBankAccounts.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected CustomerBankAccount, got nil")

	}
}

func TestCustomerBankAccountList(t *testing.T) {
	fixtureFile := "testdata/customer_bank_accounts.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := CustomerBankAccountListParams{}

	o, err :=
		client.CustomerBankAccounts.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.CustomerBankAccounts == nil {

		t.Fatalf("Expected list of CustomerBankAccounts, got nil")

	}
}

func TestCustomerBankAccountGet(t *testing.T) {
	fixtureFile := "testdata/customer_bank_accounts.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.CustomerBankAccounts.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected CustomerBankAccount, got nil")

	}
}

func TestCustomerBankAccountUpdate(t *testing.T) {
	fixtureFile := "testdata/customer_bank_accounts.json"
	server := runServer(t, fixtureFile, "update")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := CustomerBankAccountUpdateParams{}

	o, err :=
		client.CustomerBankAccounts.Update(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected CustomerBankAccount, got nil")

	}
}

func TestCustomerBankAccountDisable(t *testing.T) {
	fixtureFile := "testdata/customer_bank_accounts.json"
	server := runServer(t, fixtureFile, "disable")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.CustomerBankAccounts.Disable(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected CustomerBankAccount, got nil")

	}
}
