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

func TestVerificationDetailCreateCodeSample(t *testing.T) {
	server := RunCodeSampleServer("verification_details", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	verificationDetailCreateParams := gocardless.VerificationDetailCreateParams{
		Name:          "Acme",
		CompanyNumber: "03768189",
		AddressLine1:  "12 Drury lane",
		City:          "London",
		Description:   "wine and cheese seller",
		PostalCode:    "B4 7NJ",
		Directors: []gocardless.VerificationDetailCreateParamsDirectors{
			{
				GivenName:   "Gandalf",
				FamilyName:  "Grey",
				City:        "London",
				DateOfBirth: "1986-02-19",
				Street:      "Drury Lane",
				PostalCode:  "B4 7NJ",
				CountryCode: "GB",
			},
		},
		Links: gocardless.VerificationDetailCreateParamsLinks{
			Creditor: "CR123",
		},
	}
	_ = verificationDetailCreateParams

	verificationDetail, err := client.VerificationDetails.Create(ctx, verificationDetailCreateParams)
	_ = verificationDetail
	_ = err
	if err != nil {
		fmt.Printf("error creating verification detail: %s", err.Error())
		return
	}

}

func TestVerificationDetailListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("verification_details", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	verificationDetailListParams := gocardless.VerificationDetailListParams{
		Creditor: "CR123",
	}
	_ = verificationDetailListParams
	verificationDetailList, err := client.VerificationDetails.List(ctx, verificationDetailListParams)
	_ = verificationDetailList
	_ = err
	if err != nil {
		fmt.Printf("error listing verification details: %+v", err)
		return
	}

	for _, verificationDetail := range verificationDetailList.VerificationDetails {
		fmt.Println(verificationDetail.Name)
		fmt.Println(verificationDetail.Description)
	}

}
