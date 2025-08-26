package gocardless

import (
	"context"
	"testing"
)

func TestCreditorBankAccountValidateValidate(t *testing.T) {
	fixtureFile := "testdata/creditor_bank_account_validates.json"
	server := runServer(t, fixtureFile, "validate")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := CreditorBankAccountValidateValidateParams{}

	o, err :=
		client.CreditorBankAccountValidates.Validate(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected CreditorBankAccountValidate, got nil")

	}
}
