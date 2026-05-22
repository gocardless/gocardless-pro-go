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

func TestPaymentCreateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("payments", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	paymentCreateParams := gocardless.PaymentCreateParams{
		Amount:     100,
		Currency:   "GBP",
		ChargeDate: "2014-05-19",
		Reference:  "WINEBOX001",
		Metadata:   map[string]string{"order_dispatch_date": "2014-05-22"},
		Links: gocardless.PaymentCreateParamsLinks{
			Mandate: "MD123",
		},
	}

	payment, err := client.Payments.Create(ctx, paymentCreateParams)

}

func TestPaymentListCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("payments", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	paymentListParams := gocardless.PaymentListParams{
		CreatedAt: &gocardless.PaymentListParamsCreatedAt{
			Gt: "2020-01-01T17:01:06.000Z",
		},
	}

	paymentListResult, err := client.Payments.List(ctx, paymentListParams)
	for _, payment := range paymentListResult.Payments {
		fmt.Println(payment.Id)
	}

}

func TestPaymentGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("payments", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	payment, err := client.Payments.Get(ctx, "PM123")

}

func TestPaymentUpdateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("payments", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	paymentUpdateParams := gocardless.PaymentUpdateParams{
		Metadata: map[string]string{"key": "value"},
	}

	payment, err := client.Payments.Update(ctx, "PM123", paymentUpdateParams)

}

func TestPaymentCancelCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("payments", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	paymentCancelParams := gocardless.PaymentCancelParams{}
	payment, err := client.Payments.Cancel(ctx, "PM123", paymentCancelParams)

}

func TestPaymentRetryCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("payments", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	paymentRetryParams := gocardless.PaymentRetryParams{}
	payment, err := client.Payments.Retry(ctx, "PM123", paymentRetryParams)

}
