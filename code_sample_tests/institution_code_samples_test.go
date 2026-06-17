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

func TestInstitutionListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("institutions", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	institutionListParams := gocardless.InstitutionListParams{}
	_ = institutionListParams
	institutionListResult, err := client.Institutions.List(ctx, institutionListParams)
	_ = institutionListResult
	_ = err
	for _, institution := range institutionListResult.Institutions {
		fmt.Println(institution.Id)
	}

}

func TestInstitutionListForBillingRequestCodeSample(t *testing.T) {
	server := RunCodeSampleServer("institutions", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	institutionListForBillingRequestParams := gocardless.InstitutionListForBillingRequestParams{
		CountryCode: "GB",
	}
	_ = institutionListForBillingRequestParams
	institutionListForBillingRequestResult, err := client.Institutions.ListForBillingRequest(ctx, "BR123", institutionListForBillingRequestParams)
	_ = institutionListForBillingRequestResult
	_ = err
	for _, institution := range institutionListForBillingRequestResult.Institutions {
		fmt.Println(institution.Id)
	}

}
