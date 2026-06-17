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

func TestPayoutListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("payouts", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	payoutListParams := gocardless.PayoutListParams{}
	_ = payoutListParams
	payoutListResult, err := client.Payouts.List(ctx, payoutListParams)
	_ = payoutListResult
	_ = err
	for _, payout := range payoutListResult.Payouts {
		fmt.Println(payout.Amount)
	}

}

func TestPayoutGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("payouts", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	payout, err := client.Payouts.Get(ctx, "PO123")
	_ = payout
	_ = err

}

func TestPayoutUpdateCodeSample(t *testing.T) {
	server := RunCodeSampleServer("payouts", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	payoutUpdateParams := gocardless.PayoutUpdateParams{
		Metadata: map[string]string{"key": "value"},
	}
	_ = payoutUpdateParams

	payout, err := client.Payouts.Update(ctx, "PO123", payoutUpdateParams)
	_ = payout
	_ = err

}
