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
	server := RunCodeSampleServer("customer_bank_accounts", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	customerBankAccountCreateParams := gocardless.CustomerBankAccountCreateParams{
		AccountNumber:     "55779911",
		BranchCode:        "200000",
		AccountHolderName: "Frank Osborne",
		CountryCode:       "GB",
		Links: gocardless.CustomerBankAccountCreateParamsLinks{
			Customer: "CU123",
		},
	}
	_ = customerBankAccountCreateParams

	customerBankAccount, err := client.CustomerBankAccounts.Create(ctx, customerBankAccountCreateParams)
	_ = customerBankAccount
	_ = err

}

func TestCustomerBankAccountListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("customer_bank_accounts", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	customerBankAccountListParams := gocardless.CustomerBankAccountListParams{
		Enabled: true,
	}
	_ = customerBankAccountListParams

	customerBankAccountListResult, err := client.CustomerBankAccounts.List(ctx, customerBankAccountListParams)
	_ = customerBankAccountListResult
	_ = err
	for _, customerBankAccount := range customerBankAccountListResult.CustomerBankAccounts {
		fmt.Println(customerBankAccount.Id)
	}

}

func TestCustomerBankAccountGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("customer_bank_accounts", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	customerBankAccount, err := client.CustomerBankAccounts.Get(ctx, "BA123")
	_ = customerBankAccount
	_ = err

}

func TestCustomerBankAccountUpdateCodeSample(t *testing.T) {
	server := RunCodeSampleServer("customer_bank_accounts", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	customerBankAccountUpdateParams := gocardless.CustomerBankAccountUpdateParams{
		Metadata: map[string]string{"key": "value"},
	}
	_ = customerBankAccountUpdateParams

	customerBankAccount, err := client.CustomerBankAccounts.Update(ctx, "BA123", customerBankAccountUpdateParams)
	_ = customerBankAccount
	_ = err

}

func TestCustomerBankAccountDisableCodeSample(t *testing.T) {
	server := RunCodeSampleServer("customer_bank_accounts", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	customerBankAccount, err := client.CustomerBankAccounts.Disable(ctx, "BA123")
	_ = customerBankAccount
	_ = err

}
