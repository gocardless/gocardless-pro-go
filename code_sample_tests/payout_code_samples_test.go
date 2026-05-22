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
	server := gocardless.RunCodeSampleServer("payouts", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	payoutListParams := gocardless.PayoutListParams{}
	payoutListResult, err := client.Payouts.List(ctx, payoutListParams)
	for _, payout := range payoutListResult.Payouts {
		fmt.Println(payout.Amount)
	}

}

func TestPayoutGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("payouts", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	payout, err := client.Payouts.Get(ctx, "PO123")

}

func TestPayoutUpdateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("payouts", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	payoutUpdateParams := gocardless.PayoutUpdateParams{
		Metadata: map[string]string{"key": "value"},
	}

	payout, err := client.Payouts.Update(ctx, "PO123", payoutUpdateParams)

}
