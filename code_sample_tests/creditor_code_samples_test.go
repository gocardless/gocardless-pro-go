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

func TestCreditorCreateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("creditors", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	creditorCreateParams := gocardless.CreditorCreateParams{
		Name:                "Acme",
		CountryCode:         "GB",
		CreditorType:        "company",
		BankReferencePrefix: "ACME",
	}

	creditor, err := client.Creditors.Create(ctx, creditorCreateParams)

}

func TestCreditorListCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("creditors", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	creditorListParams := gocardless.CreditorListParams{
		Limit: 3,
	}

	// List the first three creditors.
	creditorListResult, err := client.Creditors.List(ctx, creditorListParams)
	for _, creditor := range creditorListResult.Creditors {
		fmt.Println(creditor.Name)
	}

}

func TestCreditorGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("creditors", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	creditor, err := client.Creditors.Get(ctx, "CR123", gocardless.CreditorGetParams{})

}

func TestCreditorUpdateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("creditors", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	creditorUpdateParams := gocardless.CreditorUpdateParams{
		Links: &gocardless.CreditorUpdateParamsLinks{
			DefaultGbpPayoutAccount: "BA789",
		},
	}

	creditor, err := client.Creditors.Update(ctx, "CR123", creditorUpdateParams)

}
