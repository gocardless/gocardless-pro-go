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
	server := RunCodeSampleServer("mandate_pdfs", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	mandatePdfCreateParams := gocardless.MandatePdfCreateParams{
		Links: &gocardless.MandatePdfCreateParamsLinks{
			Mandate: "MD123",
		},
	}
	_ = mandatePdfCreateParams

	mandatePdf, err := client.MandatePdfs.Create(ctx, mandatePdfCreateParams)
	_ = mandatePdf
	_ = err

	requestOption := gocardless.WithIdempotencyKey("mandate_pdfs_idempotency_key")
	_ = requestOption
	mandatePdf, err = client.MandatePdfs.Create(ctx, mandatePdfCreateParams, requestOption)

	headers := map[string]string{"Accept-Language": "fr"}
	_ = headers
	requestOption = gocardless.WithHeaders(headers)
	mandatePdf, err = client.MandatePdfs.Create(ctx, mandatePdfCreateParams, requestOption)

}
