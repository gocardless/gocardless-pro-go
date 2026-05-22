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

func TestMandatePdfCreateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("mandate_pdfs", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	mandatePdfCreateParams := gocardless.MandatePdfCreateParams{
		Links: &gocardless.MandatePdfCreateParamsLinks{
			Mandate: "MD123",
		},
	}

	mandatePdf, err := client.MandatePdfs.Create(ctx, mandatePdfCreateParams)

	requestOption := gocardless.WithIdempotencyKey("mandate_pdfs_idempotency_key")
	mandatePdf, err = client.MandatePdfs.Create(ctx, mandatePdfCreateParams, requestOption)

	headers := map[string]string{"Accept-Language": "fr"}
	requestOption = gocardless.WithHeaders(headers)
	mandatePdf, err = client.MandatePdfs.Create(ctx, mandatePdfCreateParams, requestOption)

}
