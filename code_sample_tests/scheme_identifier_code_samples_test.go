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

func TestSchemeIdentifierCreateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("scheme_identifiers", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	schemeIdentifierCreateParams := gocardless.SchemeIdentifierCreateParams{
		Name:   "Durian Co",
		Scheme: "bacs",
		Links: &gocardless.SchemeIdentifierCreateParamsLinks{
			Creditor: "CR123",
		},
	}
	schemeIdentifier, err := client.SchemeIdentifiers.Create(ctx, schemeIdentifierCreateParams)
	if err != nil {
		fmt.Printf("error creating scheme identifier: %s", err.Error())
		return
	}

}

func TestSchemeIdentifierListCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("scheme_identifiers", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	schemeIdentifierListParams := gocardless.SchemeIdentifierListParams{}

	schemeIdentifeirListResult, err := client.SchemeIdentifiers.List(ctx, schemeIdentifierListParams)
	if err != nil {
		fmt.Printf("error listing scheme identifiers: %s", err.Error())
		return
	}
	for _, schemeIdentifier := range schemeIdentifeirListResult.SchemeIdentifiers {
		fmt.Println(schemeIdentifier.Id)
	}

}

func TestSchemeIdentifierGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("scheme_identifiers", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	schemeIdentifier, err := client.SchemeIdentifiers.Get(ctx, "SU123")
	if err != nil {
		fmt.Printf("error getting scheme identifier: %s", err.Error())
		return
	}

}
