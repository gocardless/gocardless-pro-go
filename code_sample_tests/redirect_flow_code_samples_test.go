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

func TestRedirectFlowCreateCodeSample(t *testing.T) {
	server := RunCodeSampleServer("redirect_flows", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	redirectFlowCreateParams := gocardless.RedirectFlowCreateParams{
		Description: "Cider Barrels",
		Links: &gocardless.RedirectFlowCreateParamsLinks{
			Creditor: "CR00006A7FRJA5",
		},
		PrefilledCustomer: &gocardless.RedirectFlowCreateParamsPrefilledCustomer{
			AddressLine1: "338-346 Goswell Road",
			City:         "London",
			GivenName:    "Tim",
			FamilyName:   "Rogers",
			Email:        "tim@gocardless.com",
			PostalCode:   "EC1V 7LQ",
		},
		Scheme:             "bacs",
		SessionToken:       "dummy_session_token",
		SuccessRedirectUrl: "https://developer.gocardless.com/example-redirect-uri/",
	}
	_ = redirectFlowCreateParams

	redirectFlow, err := client.RedirectFlows.Create(ctx, redirectFlowCreateParams)
	_ = redirectFlow
	_ = err

}

func TestRedirectFlowGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("redirect_flows", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	redirectFlow, err := client.RedirectFlows.Get(ctx, "RE123")
	_ = redirectFlow
	_ = err

}

func TestRedirectFlowCompleteCodeSample(t *testing.T) {
	server := RunCodeSampleServer("redirect_flows", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	redirectFlowCompleteParams := gocardless.RedirectFlowCompleteParams{
		SessionToken: "dummy_session_token",
	}
	_ = redirectFlowCompleteParams

	redirectFlow, err := client.RedirectFlows.Complete(ctx, "RE0003QNP5DE2101R80QZHJ2X12P93Q4", redirectFlowCompleteParams)
	_ = redirectFlow
	_ = err

}
