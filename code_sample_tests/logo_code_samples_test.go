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

func TestLogoCreateForCreditorCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("logos", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	logoCreateForCreditorParams := gocardless.LogoCreateForCreditorParams{
		Image: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAYAAADgdz34AAAABmJLR0QA/wD/AP+gvaeTAAAA",
		Links: &gocardless.LogoCreateForCreditorParamsLinks{
			Creditor: "CR123",
		},
	}

	logo, err := client.Logos.CreateForCreditor(ctx, logoCreateForCreditorParams)

}
