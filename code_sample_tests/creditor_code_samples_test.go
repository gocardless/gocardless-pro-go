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
	server := RunCodeSampleServer("creditors", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	creditorCreateParams := gocardless.CreditorCreateParams{
		Name:                "Acme",
		CountryCode:         "GB",
		CreditorType:        "company",
		BankReferencePrefix: "ACME",
	}
	_ = creditorCreateParams

	creditor, err := client.Creditors.Create(ctx, creditorCreateParams)
	_ = creditor
	_ = err

}

func TestCreditorListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("creditors", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	creditorListParams := gocardless.CreditorListParams{
		Limit: 3,
	}
	_ = creditorListParams

	// List the first three creditors.
	creditorListResult, err := client.Creditors.List(ctx, creditorListParams)
	_ = creditorListResult
	_ = err
	for _, creditor := range creditorListResult.Creditors {
		fmt.Println(creditor.Name)
	}

}

func TestCreditorGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("creditors", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	creditor, err := client.Creditors.Get(ctx, "CR123", gocardless.CreditorGetParams{})
	_ = creditor
	_ = err

}

func TestCreditorUpdateCodeSample(t *testing.T) {
	server := RunCodeSampleServer("creditors", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	creditorUpdateParams := gocardless.CreditorUpdateParams{
		Links: &gocardless.CreditorUpdateParamsLinks{
			DefaultGbpPayoutAccount: "BA789",
		},
	}
	_ = creditorUpdateParams

	creditor, err := client.Creditors.Update(ctx, "CR123", creditorUpdateParams)
	_ = creditor
	_ = err

}
