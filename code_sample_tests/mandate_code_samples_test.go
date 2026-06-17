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
	server := RunCodeSampleServer("mandates", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	mandateCreateParams := gocardless.MandateCreateParams{
		Scheme: "bacs",
		Links: gocardless.MandateCreateParamsLinks{
			CustomerBankAccount: "BA123",
			Creditor:            "CR123",
		},
		Metadata: map[string]string{"contract": "ABCD1234"},
	}
	_ = mandateCreateParams

	mandate, err := client.Mandates.Create(ctx, mandateCreateParams)
	_ = mandate
	_ = err

}

func TestMandateListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("mandates", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	mandateListParams := gocardless.MandateListParams{
		Customer: "CU123",
	}
	_ = mandateListParams

	mandateListResult, err := client.Mandates.List(ctx, mandateListParams)
	_ = mandateListResult
	_ = err
	for _, mandate := range mandateListResult.Mandates {
		fmt.Println(mandate.Id)
	}

}

func TestMandateGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("mandates", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	mandate, err := client.Mandates.Get(ctx, "MD123")
	_ = mandate
	_ = err

}

func TestMandateUpdateCodeSample(t *testing.T) {
	server := RunCodeSampleServer("mandates", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	mandateUpdateParams := gocardless.MandateUpdateParams{
		Metadata: map[string]string{"key": "value"},
	}
	_ = mandateUpdateParams

	mandate, err := client.Mandates.Update(ctx, "MD123", mandateUpdateParams)
	_ = mandate
	_ = err

}

func TestMandateCancelCodeSample(t *testing.T) {
	server := RunCodeSampleServer("mandates", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	mandateCancelParams := gocardless.MandateCancelParams{}
	_ = mandateCancelParams
	mandate, err := client.Mandates.Cancel(ctx, "MD123", mandateCancelParams)
	_ = mandate
	_ = err

}

func TestMandateReinstateCodeSample(t *testing.T) {
	server := RunCodeSampleServer("mandates", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	mandateReinstateParams := gocardless.MandateReinstateParams{}
	_ = mandateReinstateParams
	mandate, err := client.Mandates.Reinstate(ctx, "MD123", mandateReinstateParams)
	_ = mandate
	_ = err

}
