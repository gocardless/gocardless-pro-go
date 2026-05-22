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

func TestPayerAuthorisationGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("payer_authorisations", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	payerAuthorisation, err := client.PayerAuthorisations.Get(ctx, "PAU123")

}

func TestPayerAuthorisationCreateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("payer_authorisations", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	payerAuthorisationCreateParams := gocardless.PayerAuthorisationCreateParams{
		Customer: gocardless.PayerAuthorisationCreateParamsCustomer{
			Email:      "mail@example.com",
			GivenName:  "Name",
			FamilyName: "Surname",
			Metadata:   map[string]string{"salesforce_id": "EFGH5678"},
		},
		BankAccount: gocardless.PayerAuthorisationCreateParamsBankAccount{
			AccountHolderName: "Name Surname",
			BranchCode:        "200000",
			AccountNumber:     "55779911",
		},
		Mandate: gocardless.PayerAuthorisationCreateParamsMandate{
			Reference: "XYZ789",
		},
	}

	payerAuthorisation, err := client.PayerAuthorisations.Create(ctx, payerAuthorisationCreateParams)

}

func TestPayerAuthorisationUpdateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("payer_authorisations", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	payerAuthorisationUpdateParams := gocardless.PayerAuthorisationUpdateParams{
		Customer: gocardless.PayerAuthorisationUpdateParamsCustomer{
			Email:      "mail@example.com",
			GivenName:  "Name",
			FamilyName: "Surname",
			Metadata:   map[string]string{"salesforce_id": "EFGH5678"},
		},
		BankAccount: gocardless.PayerAuthorisationUpdateParamsBankAccount{
			AccountHolderName: "Name Surname",
			BranchCode:        "200000",
			AccountNumber:     "55779911",
		},
		Mandate: gocardless.PayerAuthorisationUpdateParamsMandate{
			Reference: "XYZ789",
		},
	}

	payerAuthorisation, err := client.PayerAuthorisations.Update(ctx, "PA123", payerAuthorisationUpdateParams)

}

func TestPayerAuthorisationSubmitCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("payer_authorisations", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	payerAuthorisation, err := client.PayerAuthorisations.Submit(ctx, "PAU123")

}

func TestPayerAuthorisationConfirmCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("payer_authorisations", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	payerAuthorisation, err := client.PayerAuthorisations.Confirm(ctx, "PAU123")

}
