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

func TestBillingRequestFlowCreateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("billing_request_flows", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	billingRequestFlowCreateParams := gocardless.BillingRequestFlowCreateParams{
		RedirectUri: "https://my-company.com/landing",
		ExitUri:     "https://my-company.com/exit",
		PrefilledCustomer: &gocardless.BillingRequestFlowCreateParamsPrefilledCustomer{
			AddressLine1: "338-346 Goswell Road",
			City:         "London",
			GivenName:    "Tim",
			FamilyName:   "Rogers",
			Email:        "tim@gocardless.com",
			PostalCode:   "EC1V 7LQ",
		},
		Links: gocardless.BillingRequestFlowCreateParamsLinks{
			BillingRequest: "BR123",
		},
	}

	billingRequestFlow, err := client.BillingRequestFlows.Create(ctx, billingRequestFlowCreateParams)

}

func TestBillingRequestFlowInitialiseCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("billing_request_flows", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	billingRequestFlowInitialiseParams := gocardless.BillingRequestFlowInitialiseParams{}
	billingRequestFlow, err := client.BillingRequestFlows.Initialise(ctx, "BRF123", billingRequestFlowInitialiseParams)

}
