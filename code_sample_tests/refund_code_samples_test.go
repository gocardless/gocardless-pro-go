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

func TestRefundCreateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("refunds", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	refundCreateParams := gocardless.RefundCreateParams{
		Amount:                  100,
		TotalAmountConfirmation: 150,
		Reference:               "Acme refund",
		Metadata:                map[string]string{"reason": "late delivery"},
		Links: gocardless.RefundCreateParamsLinks{
			Payment: "PM123",
		},
	}

	refund, err := client.Refunds.Create(ctx, refundCreateParams)

}

func TestRefundListCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("refunds", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	refundListParams := gocardless.RefundListParams{
		Mandate: "MD123",
	}
	refundListResult, err := client.Refunds.List(ctx, refundListParams)
	for _, refund := range refundListResult.Refunds {
		fmt.Println(refund.Id)
	}

}

func TestRefundGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("refunds", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	refund, err := client.Refunds.Get(ctx, "RF123")

}

func TestRefundUpdateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("refunds", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	refundUpdateParams := gocardless.RefundUpdateParams{
		Metadata: map[string]string{"key": "value"},
	}
	refund, err := client.Refunds.Update(ctx, "RF123", refundUpdateParams)

}
