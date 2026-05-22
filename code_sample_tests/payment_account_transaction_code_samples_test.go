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
	server := gocardless.RunCodeSampleServer("payment_account_transactions", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	transaction, err := client.PaymentAccountTransactions.Get(ctx, "PATR1234")

}

func TestPaymentAccountTransactionListCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("payment_account_transactions", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	paymentAccountTransactionListParams := gocardless.PaymentAccountTransactionListParams{
		ValueDateFrom: "2024-01-01",
		ValueDateTo:   "2024-01-31",
	}

	transactionListResult, err := client.PaymentAccountTransactions.List(ctx, "BA12345", paymentAccountTransactionListParams)

}
