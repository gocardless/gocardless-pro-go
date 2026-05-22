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

func TestCustomerBankAccountCreateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("customer_bank_accounts", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	customerBankAccountCreateParams := gocardless.CustomerBankAccountCreateParams{
		AccountNumber:     "55779911",
		BranchCode:        "200000",
		AccountHolderName: "Frank Osborne",
		CountryCode:       "GB",
		Links: gocardless.CustomerBankAccountCreateParamsLinks{
			Customer: "CU123",
		},
	}

	customerBankAccount, err := client.CustomerBankAccounts.Create(ctx, customerBankAccountCreateParams)

}

func TestCustomerBankAccountListCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("customer_bank_accounts", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	customerBankAccountListParams := gocardless.CustomerBankAccountListParams{
		Enabled: true,
	}

	customerBankAccountListResult, err := client.CustomerBankAccounts.List(ctx, customerBankAccountListParams)
	for _, customerBankAccount := range customerBankAccountListResult.CustomerBankAccounts {
		fmt.Println(customerBankAccount.Id)
	}

}

func TestCustomerBankAccountGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("customer_bank_accounts", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	customerBankAccount, err := client.CustomerBankAccounts.Get(ctx, "BA123")

}

func TestCustomerBankAccountUpdateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("customer_bank_accounts", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	customerBankAccountUpdateParams := gocardless.CustomerBankAccountUpdateParams{
		Metadata: map[string]string{"key": "value"},
	}

	customerBankAccount, err := client.CustomerBankAccounts.Update(ctx, "BA123", customerBankAccountUpdateParams)

}

func TestCustomerBankAccountDisableCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("customer_bank_accounts", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	customerBankAccount, err := client.CustomerBankAccounts.Disable(ctx, "BA123")

}
