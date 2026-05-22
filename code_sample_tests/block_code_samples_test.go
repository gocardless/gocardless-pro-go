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

func TestBlockCreateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("blocks", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	blockCreateParams := gocardless.BlockCreateParams{
		BlockType:         "email",
		ReasonType:        "no_intent_to_pay",
		ResourceReference: "example@example.com",
	}

	block, err := client.Blocks.Create(ctx, blockCreateParams)

}

func TestBlockGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("blocks", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	block, err := client.Blocks.Get(ctx, "BLC456")

}

func TestBlockListCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("blocks", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	blockListParams := gocardless.BlockListParams{}
	blockListResult, err := client.Blocks.List(ctx, blockListParams)
	for _, block := range blockListResult.Blocks {
		fmt.Println(block.Id)
	}

}

func TestBlockDisableCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("blocks", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	block, err := client.Blocks.Disable(ctx, "BLC123")

}

func TestBlockEnableCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("blocks", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	block, err := client.Blocks.Enable(ctx, "BLC123")

}

func TestBlockBlockByRefCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("blocks", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	blockBlockByRefParams := gocardless.BlockBlockByRefParams{
		ReferenceType:  "customer",
		ReferenceValue: "CU123",
		ReasonType:     "no_intent_to_pay",
	}

	blockBlockByRefResult, err := client.Blocks.BlockByRef(ctx, blockBlockByRefParams)
	for _, block := range blockBlockByRefResult.Blocks {
		fmt.Println(block.Id)
	}

}
