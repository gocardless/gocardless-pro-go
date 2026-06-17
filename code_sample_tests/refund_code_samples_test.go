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
	server := RunCodeSampleServer("refunds", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	refundCreateParams := gocardless.RefundCreateParams{
		Amount:                  100,
		TotalAmountConfirmation: 150,
		Reference:               "Acme refund",
		Metadata:                map[string]string{"reason": "late delivery"},
		Links: gocardless.RefundCreateParamsLinks{
			Payment: "PM123",
		},
	}
	_ = refundCreateParams

	refund, err := client.Refunds.Create(ctx, refundCreateParams)
	_ = refund
	_ = err

}

func TestRefundListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("refunds", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	refundListParams := gocardless.RefundListParams{
		Mandate: "MD123",
	}
	_ = refundListParams
	refundListResult, err := client.Refunds.List(ctx, refundListParams)
	_ = refundListResult
	_ = err
	for _, refund := range refundListResult.Refunds {
		fmt.Println(refund.Id)
	}

}

func TestRefundGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("refunds", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	refund, err := client.Refunds.Get(ctx, "RF123")
	_ = refund
	_ = err

}

func TestRefundUpdateCodeSample(t *testing.T) {
	server := RunCodeSampleServer("refunds", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	refundUpdateParams := gocardless.RefundUpdateParams{
		Metadata: map[string]string{"key": "value"},
	}
	_ = refundUpdateParams
	refund, err := client.Refunds.Update(ctx, "RF123", refundUpdateParams)
	_ = refund
	_ = err

}
