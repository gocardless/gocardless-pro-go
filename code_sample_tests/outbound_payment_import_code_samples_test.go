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

func TestOutboundPaymentImportCreateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("outbound_payment_imports", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	outboundPaymentImportCreateParams := gocardless.OutboundPaymentImportCreateParams{
		EntryItems: []gocardless.OutboundPaymentImportCreateParamsEntryItems{
			{
				Amount:                 1000,
				Scheme:                 "faster_payments",
				Reference:              "Invoice 123",
				RecipientBankAccountId: "BA123",
			},
			{
				Amount:                 2000,
				Scheme:                 "faster_payments",
				Reference:              "Invoice 124",
				RecipientBankAccountId: "BA456",
				Metadata: map[string]string{
					"order_id": "ORD-789",
				},
			},
		},
		Links: &gocardless.OutboundPaymentImportCreateParamsLinks{
			Creditor: "CR123",
		},
	}

	outboundPaymentImport, err := client.OutboundPaymentImports.Create(ctx, outboundPaymentImportCreateParams)

}

func TestOutboundPaymentImportGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("outbound_payment_imports", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	outboundPaymentImport, err := client.OutboundPaymentImports.Get(ctx, "IM123", gocardless.OutboundPaymentImportGetParams{})

}

func TestOutboundPaymentImportListCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("outbound_payment_imports", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	outboundPaymentImportListParams := gocardless.OutboundPaymentImportListParams{
		Limit: 10,
	}

	outboundPaymentImportListResult, err := client.OutboundPaymentImports.List(ctx, outboundPaymentImportListParams)

}
