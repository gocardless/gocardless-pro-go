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
	server := RunCodeSampleServer("bank_account_holder_verifications", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	bankAccountHolderVerificationCreateParams := gocardless.BankAccountHolderVerificationCreateParams{
		Type: "confirmation_of_payee",
		Links: gocardless.BankAccountHolderVerificationCreateParamsLinks{
			BankAccount: "BA123",
		},
	}
	_ = bankAccountHolderVerificationCreateParams

	verification, err := client.BankAccountHolderVerifications.Create(ctx, bankAccountHolderVerificationCreateParams)
	_ = verification
	_ = err

}

func TestBankAccountHolderVerificationGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("bank_account_holder_verifications", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	verification, err := client.BankAccountHolderVerifications.Get(ctx, "BAHV123")
	_ = verification
	_ = err

}
