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

func TestOutboundPaymentCreateCodeSample(t *testing.T) {
	server := RunCodeSampleServer("outbound_payments", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	outboundPaymentCreateParams := gocardless.OutboundPaymentCreateParams{
		Amount:      1000,
		Scheme:      "faster_payments",
		Description: "Reward Payment (August 2024)",
		Reference:   "Invoice 123",
		Links: gocardless.OutboundPaymentCreateParamsLinks{
			Creditor:             "CR123",
			RecipientBankAccount: "BA123",
		},
	}
	_ = outboundPaymentCreateParams

	outboundPayment, err := client.OutboundPayments.Create(ctx, outboundPaymentCreateParams)
	_ = outboundPayment
	_ = err

}

func TestOutboundPaymentWithdrawCodeSample(t *testing.T) {
	server := RunCodeSampleServer("outbound_payments", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	outboundPaymentWithdrawParams := gocardless.OutboundPaymentWithdrawParams{
		Amount:      5000,
		Scheme:      "faster_payments",
		Description: "Withdraw funds to business account",
		Links: &gocardless.OutboundPaymentWithdrawParamsLinks{
			Creditor: "CR123",
		},
	}
	_ = outboundPaymentWithdrawParams

	outboundPayment, err := client.OutboundPayments.Withdraw(ctx, outboundPaymentWithdrawParams)
	_ = outboundPayment
	_ = err

}

func TestOutboundPaymentCancelCodeSample(t *testing.T) {
	server := RunCodeSampleServer("outbound_payments", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	outboundPayment, err := client.OutboundPayments.Cancel(ctx, "OUT123", gocardless.OutboundPaymentCancelParams{})
	_ = outboundPayment
	_ = err

}

func TestOutboundPaymentApproveCodeSample(t *testing.T) {
	server := RunCodeSampleServer("outbound_payments", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	outboundPayment, err := client.OutboundPayments.Approve(ctx, "OUT123", gocardless.OutboundPaymentApproveParams{})
	_ = outboundPayment
	_ = err

}

func TestOutboundPaymentGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("outbound_payments", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	outboundPayment, err := client.OutboundPayments.Get(ctx, "OUT123")
	_ = outboundPayment
	_ = err

}

func TestOutboundPaymentListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("outbound_payments", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	outboundPaymentListParams := gocardless.OutboundPaymentListParams{
		Limit: 10,
	}
	_ = outboundPaymentListParams

	outboundPaymentListResult, err := client.OutboundPayments.List(ctx, outboundPaymentListParams)
	_ = outboundPaymentListResult
	_ = err

}

func TestOutboundPaymentUpdateCodeSample(t *testing.T) {
	server := RunCodeSampleServer("outbound_payments", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	outboundPaymentUpdateParams := gocardless.OutboundPaymentUpdateParams{
		Metadata: map[string]string{
			"invoice_id": "INV-1234",
		},
	}
	_ = outboundPaymentUpdateParams

	outboundPayment, err := client.OutboundPayments.Update(ctx, "OUT123", outboundPaymentUpdateParams)
	_ = outboundPayment
	_ = err

}

func TestOutboundPaymentStatsCodeSample(t *testing.T) {
	server := RunCodeSampleServer("outbound_payments", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	stats, err := client.OutboundPayments.Stats(ctx, gocardless.OutboundPaymentStatsParams{})
	_ = stats
	_ = err

}
