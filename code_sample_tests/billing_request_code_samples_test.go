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
	server := gocardless.RunCodeSampleServer("billing_requests", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

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

	billingRequest, err := client.BillingRequests.Create(ctx, billingRequestCreateParams)

}

func TestBillingRequestCollectCustomerDetailsCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("billing_requests", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

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

	billingRequest, err := client.BillingRequests.CollectCustomerDetails(ctx, "BR123", billingRequestCollectCustomerDetailsParams)

}

func TestBillingRequestCollectBankAccountCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("billing_requests", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	billingRequestCollectBankAccountParams := gocardless.BillingRequestCollectBankAccountParams{
		BranchCode:        "200000",
		CountryCode:       "GB",
		AccountNumber:     "55779911",
		AccountHolderName: "Frank Osborne",
	}

	billingRequest, err := client.BillingRequests.CollectBankAccount(ctx, "BR123", billingRequestCollectBankAccountParams)

}

func TestBillingRequestConfirmPayerDetailsCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("billing_requests", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	BillingRequestConfirmPayerDetailsParams := gocardless.BillingRequestConfirmPayerDetailsParams{}
	billingRequest, err := client.BillingRequests.ConfirmPayerDetails(ctx, "BR123", BillingRequestConfirmPayerDetailsParams)

}

func TestBillingRequestFulfilCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("billing_requests", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	billingRequestFulfilParams := gocardless.BillingRequestFulfilParams{}
	billingRequest, err := client.BillingRequests.Fulfil(ctx, "BR123", billingRequestFulfilParams)

}

func TestBillingRequestCancelCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("billing_requests", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	billingRequestCancelParams := gocardless.BillingRequestCancelParams{}
	billingRequest, err := client.BillingRequests.Cancel(ctx, "BR123", billingRequestCancelParams)

}

func TestBillingRequestListCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("billing_requests", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	billingRequestListParams := gocardless.BillingRequestListParams{}
	billingRequestListResult, err := client.BillingRequests.List(ctx, billingRequestListParams)
	for _, billingRequest := range billingRequestListResult.BillingRequests {
		fmt.Println(billingRequest.Id)
	}

}

func TestBillingRequestGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("billing_requests", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	billingRequest, err := client.BillingRequests.Get(ctx, "BR123")

}

func TestBillingRequestNotifyCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("billing_requests", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	billingRequestNotifyParams := gocardless.BillingRequestNotifyParams{
		NotificationType: "email",
		RedirectUri:      "https://my-company.com",
	}

	billingRequest, err := client.BillingRequests.Notify(ctx, "BR123", billingRequestNotifyParams)

}
