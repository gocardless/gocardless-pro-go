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
	server := RunCodeSampleServer("mandate_imports", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	mandateImportCreateParams := gocardless.MandateImportCreateParams{
		Scheme: "bacs",
	}
	_ = mandateImportCreateParams

	mandateImport, err := client.MandateImports.Create(ctx, mandateImportCreateParams)
	_ = mandateImport
	_ = err

}

func TestMandateImportGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("mandate_imports", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	mandateImport, err := client.MandateImports.Get(ctx, "IM123", gocardless.MandateImportGetParams{})
	_ = mandateImport
	_ = err

}

func TestMandateImportSubmitCodeSample(t *testing.T) {
	server := RunCodeSampleServer("mandate_imports", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	mandateImportSubmitParams := gocardless.MandateImportSubmitParams{}
	_ = mandateImportSubmitParams
	mandateImport, err := client.MandateImports.Submit(ctx, "IM123", mandateImportSubmitParams)
	_ = mandateImport
	_ = err

}

func TestMandateImportCancelCodeSample(t *testing.T) {
	server := RunCodeSampleServer("mandate_imports", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	mandateImportCancelParams := gocardless.MandateImportCancelParams{}
	_ = mandateImportCancelParams
	mandateImport, err := client.MandateImports.Cancel(ctx, "IM123", mandateImportCancelParams)
	_ = mandateImport
	_ = err

}
