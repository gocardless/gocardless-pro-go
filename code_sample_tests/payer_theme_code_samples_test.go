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

func TestPayerThemeCreateForCreditorCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("payer_themes", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	payerThemeCreateForCreditorParams := gocardless.PayerThemeCreateForCreditorParams{
		ButtonBackgroundColour: "#128DAA",
		ContentBoxBorderColour: "#BD10E0",
		HeaderBackgroundColour: "#BD10E0",
		LinkTextColour:         "#7ED321",
		Links: &gocardless.PayerThemeCreateForCreditorParamsLinks{
			Creditor: "CR123",
		},
	}

	payerThemes, err := client.PayerThemes.CreateForCreditor(ctx, payerThemeCreateForCreditorParams)

}
