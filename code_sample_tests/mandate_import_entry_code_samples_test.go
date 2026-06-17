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

func TestMandateImportEntryCreateCodeSample(t *testing.T) {
	server := RunCodeSampleServer("mandate_import_entries", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	mandateImportEntryCreateParams := gocardless.MandateImportEntryCreateParams{
		Customer: gocardless.MandateImportEntryCreateParamsCustomer{
			CompanyName: "Théâtre du Palais-Royal",
			Email:       "moliere@tdpr.fr",
		},
		BankAccount: gocardless.MandateImportEntryCreateParamsBankAccount{
			AccountHolderName: "Jean-Baptiste Poquelin",
			Iban:              "FR14BARC20000055779911",
		},
		Amendment: &gocardless.MandateImportEntryCreateParamsAmendment{
			OriginalMandateReference: "REFMANDATE",
			OriginalCreditorId:       "FR123OTHERBANK",
			OriginalCreditorName:     "Amphitryon",
		},
		Links: gocardless.MandateImportEntryCreateParamsLinks{
			MandateImport: "IM000010790WX1",
		},
	}
	_ = mandateImportEntryCreateParams

	mandateImportEntry, err := client.MandateImportEntries.Create(ctx, mandateImportEntryCreateParams)
	_ = mandateImportEntry
	_ = err

}

func TestMandateImportEntryListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("mandate_import_entries", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	mandateImportEntryListParams := gocardless.MandateImportEntryListParams{
		MandateImport: "IM000010790WX1",
	}
	_ = mandateImportEntryListParams

	mandateImportEntryListPagingIterator := client.MandateImportEntries.All(ctx, mandateImportEntryListParams)
	_ = mandateImportEntryListPagingIterator
	for mandateImportEntryListPagingIterator.Next() {
		mandateImportEntryListResult, err := mandateImportEntryListPagingIterator.Value(ctx)
		_ = mandateImportEntryListResult
		_ = err
		for _, mandateImportEntry := range mandateImportEntryListResult.MandateImportEntries {
			fmt.Println(mandateImportEntry.RecordIdentifier)
		}
	}

}
