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

func TestNegativeBalanceLimitListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("negative_balance_limits", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	params := gocardless.NegativeBalanceLimitListParams{
		Currency: "GBP",
		Creditor: "CR123",
	}
	_ = params

	response, err := client.NegativeBalanceLimits.List(ctx, params)
	_ = response
	_ = err

}
