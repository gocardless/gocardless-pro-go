package code_sample_tests // Use a distinct package from the library itself to ensure code samples are tested in the same way as user code

// Code Sample Tests
// These tests verify that the documentation code samples are syntactically valid
// and can execute against a mocked API without errors.
//
// IMPORTANT: These tests do NOT verify business logic - they only verify that
// the code samples compile and execute without syntax errors.

import (
	"context"
	"fmt"
	"testing"

	gocardless "github.com/gocardless/gocardless-pro-go/v6"
)

func TestCreditorBankAccountCreateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("creditor_bank_accounts", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	creditorBankAccountCreateParams := gocardless.CreditorBankAccountCreateParams{
		AccountNumber:     "55779911",
		BranchCode:        "200000",
		CountryCode:       "GB",
		AccountHolderName: "Acme",
	}

	creditorBankAccount, err := client.CreditorBankAccounts.Create(ctx, creditorBankAccountCreateParams)

}

func TestCreditorBankAccountListCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("creditor_bank_accounts", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	creditorBankAccountListParams := gocardless.CreditorBankAccountListParams{}
	creditorBankAccountListResult, err := client.CreditorBankAccounts.List(ctx, creditorBankAccountListParams)
	for _, creditorBankAccount := range creditorBankAccountListResult.CreditorBankAccounts {
		fmt.Println(creditorBankAccount.Id)
	}

}

func TestCreditorBankAccountGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("creditor_bank_accounts", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	creditorBankAccount, err := client.CreditorBankAccounts.Get(ctx, "BA123")

}

func TestCreditorBankAccountDisableCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("creditor_bank_accounts", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	creditorBankAccount, err := client.CreditorBankAccounts.Disable(ctx, "BA123")

}
