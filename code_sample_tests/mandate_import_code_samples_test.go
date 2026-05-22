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

func TestMandateImportCreateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("mandate_imports", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	mandateImportCreateParams := gocardless.MandateImportCreateParams{
		Scheme: "bacs",
	}

	mandateImport, err := client.MandateImports.Create(ctx, mandateImportCreateParams)

}

func TestMandateImportGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("mandate_imports", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	mandateImport, err := client.MandateImports.Get(ctx, "IM123", gocardless.MandateImportGetParams{})

}

func TestMandateImportSubmitCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("mandate_imports", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	mandateImportSubmitParams := gocardless.MandateImportSubmitParams{}
	mandateImport, err := client.MandateImports.Submit(ctx, "IM123", mandateImportSubmitParams)

}

func TestMandateImportCancelCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("mandate_imports", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	mandateImportCancelParams := gocardless.MandateImportCancelParams{}
	mandateImport, err := client.MandateImports.Cancel(ctx, "IM123", mandateImportCancelParams)

}
