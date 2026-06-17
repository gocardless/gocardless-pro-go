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

func TestBillingRequestCreateCodeSample(t *testing.T) {
	server := RunCodeSampleServer("billing_requests", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	billingRequestCreateParams := gocardless.BillingRequestCreateParams{
		PaymentRequest: &gocardless.BillingRequestCreateParamsPaymentRequest{
			Amount:      1000,
			Currency:    "GBP",
			Description: "First Payment",
		},
		MandateRequest: &gocardless.BillingRequestCreateParamsMandateRequest{
			Scheme: "bacs",
		},
	}
	_ = billingRequestCreateParams

	billingRequest, err := client.BillingRequests.Create(ctx, billingRequestCreateParams)
	_ = billingRequest
	_ = err

}

func TestBillingRequestCollectCustomerDetailsCodeSample(t *testing.T) {
	server := RunCodeSampleServer("billing_requests", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	billingRequestCollectCustomerDetailsParams := gocardless.BillingRequestCollectCustomerDetailsParams{
		Customer: &gocardless.BillingRequestCollectCustomerDetailsParamsCustomer{
			GivenName:  "Alice",
			FamilyName: "Smith",
			Email:      "alice@example.com",
		},
		CustomerBillingDetail: &gocardless.BillingRequestCollectCustomerDetailsParamsCustomerBillingDetail{
			AddressLine1: "1 Somewhere Lane",
		},
	}
	_ = billingRequestCollectCustomerDetailsParams

	billingRequest, err := client.BillingRequests.CollectCustomerDetails(ctx, "BR123", billingRequestCollectCustomerDetailsParams)
	_ = billingRequest
	_ = err

}

func TestBillingRequestCollectBankAccountCodeSample(t *testing.T) {
	server := RunCodeSampleServer("billing_requests", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	billingRequestCollectBankAccountParams := gocardless.BillingRequestCollectBankAccountParams{
		BranchCode:        "200000",
		CountryCode:       "GB",
		AccountNumber:     "55779911",
		AccountHolderName: "Frank Osborne",
	}
	_ = billingRequestCollectBankAccountParams

	billingRequest, err := client.BillingRequests.CollectBankAccount(ctx, "BR123", billingRequestCollectBankAccountParams)
	_ = billingRequest
	_ = err

}

func TestBillingRequestConfirmPayerDetailsCodeSample(t *testing.T) {
	server := RunCodeSampleServer("billing_requests", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	BillingRequestConfirmPayerDetailsParams := gocardless.BillingRequestConfirmPayerDetailsParams{}
	_ = BillingRequestConfirmPayerDetailsParams
	billingRequest, err := client.BillingRequests.ConfirmPayerDetails(ctx, "BR123", BillingRequestConfirmPayerDetailsParams)
	_ = billingRequest
	_ = err

}

func TestBillingRequestFulfilCodeSample(t *testing.T) {
	server := RunCodeSampleServer("billing_requests", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	billingRequestFulfilParams := gocardless.BillingRequestFulfilParams{}
	_ = billingRequestFulfilParams
	billingRequest, err := client.BillingRequests.Fulfil(ctx, "BR123", billingRequestFulfilParams)
	_ = billingRequest
	_ = err

}

func TestBillingRequestCancelCodeSample(t *testing.T) {
	server := RunCodeSampleServer("billing_requests", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	billingRequestCancelParams := gocardless.BillingRequestCancelParams{}
	_ = billingRequestCancelParams
	billingRequest, err := client.BillingRequests.Cancel(ctx, "BR123", billingRequestCancelParams)
	_ = billingRequest
	_ = err

}

func TestBillingRequestListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("billing_requests", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	billingRequestListParams := gocardless.BillingRequestListParams{}
	_ = billingRequestListParams
	billingRequestListResult, err := client.BillingRequests.List(ctx, billingRequestListParams)
	_ = billingRequestListResult
	_ = err
	for _, billingRequest := range billingRequestListResult.BillingRequests {
		fmt.Println(billingRequest.Id)
	}

}

func TestBillingRequestGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("billing_requests", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	billingRequest, err := client.BillingRequests.Get(ctx, "BR123")
	_ = billingRequest
	_ = err

}

func TestBillingRequestNotifyCodeSample(t *testing.T) {
	server := RunCodeSampleServer("billing_requests", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	billingRequestNotifyParams := gocardless.BillingRequestNotifyParams{
		NotificationType: "email",
		RedirectUri:      "https://my-company.com",
	}
	_ = billingRequestNotifyParams

	billingRequest, err := client.BillingRequests.Notify(ctx, "BR123", billingRequestNotifyParams)
	_ = billingRequest
	_ = err

}

func TestBillingRequestFallbackCodeSample(t *testing.T) {
	server := RunCodeSampleServer("billing_requests", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	billingRequest, err := client.BillingRequests.Fallback(ctx, "BR123", gocardless.BillingRequestFallbackParams{})
	_ = billingRequest
	_ = err

}

func TestBillingRequestChooseCurrencyCodeSample(t *testing.T) {
	server := RunCodeSampleServer("billing_requests", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	billingRequestChooseCurrencyParams := gocardless.BillingRequestChooseCurrencyParams{
		Currency: "GBP",
	}
	_ = billingRequestChooseCurrencyParams

	billingRequest, err := client.BillingRequests.ChooseCurrency(ctx, "BR123", billingRequestChooseCurrencyParams)
	_ = billingRequest
	_ = err

}
