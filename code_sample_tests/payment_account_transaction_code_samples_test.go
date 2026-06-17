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

func TestPaymentAccountTransactionGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("payment_account_transactions", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	transaction, err := client.PaymentAccountTransactions.Get(ctx, "PATR1234")
	_ = transaction
	_ = err

}

func TestPaymentAccountTransactionListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("payment_account_transactions", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	paymentAccountTransactionListParams := gocardless.PaymentAccountTransactionListParams{
		ValueDateFrom: "2024-01-01",
		ValueDateTo:   "2024-01-31",
	}
	_ = paymentAccountTransactionListParams

	transactionListResult, err := client.PaymentAccountTransactions.List(ctx, "BA12345", paymentAccountTransactionListParams)
	_ = transactionListResult
	_ = err

}
