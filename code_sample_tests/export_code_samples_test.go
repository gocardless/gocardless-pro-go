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

func TestExportGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("exports", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	export, err := client.Exports.Get(ctx, "EX123")
	_ = export
	_ = err

}

func TestExportListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("exports", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	exportListParams := gocardless.ExportListParams{}
	_ = exportListParams
	exportListResult, err := client.Exports.List(ctx, exportListParams)
	_ = exportListResult
	_ = err
	for _, export := range exportListResult.Exports {
		fmt.Println(export.Id)
	}

}
