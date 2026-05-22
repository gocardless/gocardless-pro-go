package code_sample_tests // Use a distinct package from the library itself to ensure code samples are tested in the same way as user code

// Code Sample Tests
// These tests verify that the documentation code samples are syntactically valid
// and can execute against a mocked API without errors.
//
// IMPORTANT: These tests do NOT verify business logic - they only verify that
// the code samples compile and execute without syntax errors.

import (
	"context"
	"testing"

	gocardless "github.com/gocardless/gocardless-pro-go/v6"
)

func TestBankAccountHolderVerificationCreateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("bank_account_holder_verifications", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	bankAccountHolderVerificationCreateParams := gocardless.BankAccountHolderVerificationCreateParams{
		Type: "confirmation_of_payee",
		Links: gocardless.BankAccountHolderVerificationCreateParamsLinks{
			BankAccount: "BA123",
		},
	}

	verification, err := client.BankAccountHolderVerifications.Create(ctx, bankAccountHolderVerificationCreateParams)

}

func TestBankAccountHolderVerificationGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("bank_account_holder_verifications", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	verification, err := client.BankAccountHolderVerifications.Get(ctx, "BAHV123")

}
