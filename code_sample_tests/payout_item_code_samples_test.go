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

func TestPayoutItemListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("payout_items", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	payoutItemsListParams := gocardless.PayoutItemListParams{
		Payout: "PO123",
	}
	_ = payoutItemsListParams

	payoutItemsListResult, err := client.PayoutItems.List(ctx, payoutItemsListParams)
	_ = payoutItemsListResult
	_ = err
	for _, payoutItem := range payoutItemsListResult.PayoutItems {
		fmt.Println(payoutItem.Amount)
	}

}
