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

func TestMandateCreateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("mandates", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	mandateCreateParams := gocardless.MandateCreateParams{
		Scheme: "bacs",
		Links: gocardless.MandateCreateParamsLinks{
			CustomerBankAccount: "BA123",
			Creditor:            "CR123",
		},
		Metadata: map[string]string{"contract": "ABCD1234"},
	}

	mandate, err := client.Mandates.Create(ctx, mandateCreateParams)

}

func TestMandateListCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("mandates", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	mandateListParams := gocardless.MandateListParams{
		Customer: "CU123",
	}

	mandateListResult, err := client.Mandates.List(ctx, mandateListParams)
	for _, mandate := range mandateListResult.Mandates {
		fmt.Println(mandate.Id)
	}

}

func TestMandateGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("mandates", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	mandate, err := client.Mandates.Get(ctx, "MD123")

}

func TestMandateUpdateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("mandates", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	mandateUpdateParams := gocardless.MandateUpdateParams{
		Metadata: map[string]string{"key": "value"},
	}

	mandate, err := client.Mandates.Update(ctx, "MD123", mandateUpdateParams)

}

func TestMandateCancelCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("mandates", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	mandateCancelParams := gocardless.MandateCancelParams{}
	mandate, err := client.Mandates.Cancel(ctx, "MD123", mandateCancelParams)

}

func TestMandateReinstateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("mandates", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	mandateReinstateParams := gocardless.MandateReinstateParams{}
	mandate, err := client.Mandates.Reinstate(ctx, "MD123", mandateReinstateParams)

}
