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

func TestTaxRateListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("tax_rates", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	taxRateListParams := gocardless.TaxRateListParams{
		Jurisdiction: "GB",
	}
	_ = taxRateListParams

	taxRateListResult, err := client.TaxRates.List(ctx, taxRateListParams)
	_ = taxRateListResult
	_ = err

}

func TestTaxRateGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("tax_rates", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	taxRate, err := client.TaxRates.Get(ctx, "GB_VAT_1")
	_ = taxRate
	_ = err

}
