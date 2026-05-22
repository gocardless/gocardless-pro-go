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

func TestOutboundPaymentImportEntryListCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("outbound_payment_import_entries", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	outboundPaymentImportEntryListParams := gocardless.OutboundPaymentImportEntryListParams{
		OutboundPaymentImport: "IM123",
		Limit:                 10,
	}

	outboundPaymentImportEntryListResult, err := client.OutboundPaymentImportEntries.List(ctx, outboundPaymentImportEntryListParams)

}
