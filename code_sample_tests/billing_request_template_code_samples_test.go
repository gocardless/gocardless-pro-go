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

func TestBillingRequestTemplateListCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("billing_request_templates", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	billingRequestTemplateListParams := gocardless.BillingRequestTemplateListParams{}
	billingRequestTemplateListResult, err := client.BillingRequestTemplates.List(ctx, billingRequestTemplateListParams)
	for _, billingRequestTemplate := range billingRequestTemplateListResult.BillingRequestTemplates {
		fmt.Println(billingRequestTemplate.Id)
	}

}

func TestBillingRequestTemplateGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("billing_request_templates", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	billingRequestTemplate, err := client.BillingRequestTemplates.Get(ctx, "BRT123")

}

func TestBillingRequestTemplateCreateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("billing_request_templates", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	billingRequestTemplateCreateParams := gocardless.BillingRequestTemplateCreateParams{
		Name:                      "12 Month Gold Plan",
		PaymentRequestDescription: "One-time joining fee",
		PaymentRequestCurrency:    "GBP",
		PaymentRequestAmount:      "6999",
		MandateRequestCurrency:    "GBP",
		RedirectUri:               "https://my-company.com/landing",
	}

	billingRequestTemplate, err := client.BillingRequestTemplates.Create(ctx, billingRequestTemplateCreateParams)

}

func TestBillingRequestTemplateUpdateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("billing_request_templates", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	billingRequestTemplateUpdateParams := gocardless.BillingRequestTemplateUpdateParams{
		Name:                 "12 Month Silver Plan",
		PaymentRequestAmount: "4999",
	}

	billingRequestTemplate, err := client.BillingRequestTemplates.Update(ctx, "BRT123", billingRequestTemplateUpdateParams)

}
