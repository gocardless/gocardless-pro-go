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
	server := RunCodeSampleServer("blocks", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	blockCreateParams := gocardless.BlockCreateParams{
		BlockType:         "email",
		ReasonType:        "no_intent_to_pay",
		ResourceReference: "example@example.com",
	}
	_ = blockCreateParams

	block, err := client.Blocks.Create(ctx, blockCreateParams)
	_ = block
	_ = err

}

func TestBlockGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("blocks", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	block, err := client.Blocks.Get(ctx, "BLC456")
	_ = block
	_ = err

}

func TestBlockListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("blocks", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	blockListParams := gocardless.BlockListParams{}
	_ = blockListParams
	blockListResult, err := client.Blocks.List(ctx, blockListParams)
	_ = blockListResult
	_ = err
	for _, block := range blockListResult.Blocks {
		fmt.Println(block.Id)
	}

}

func TestBlockDisableCodeSample(t *testing.T) {
	server := RunCodeSampleServer("blocks", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	block, err := client.Blocks.Disable(ctx, "BLC123")
	_ = block
	_ = err

}

func TestBlockEnableCodeSample(t *testing.T) {
	server := RunCodeSampleServer("blocks", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	block, err := client.Blocks.Enable(ctx, "BLC123")
	_ = block
	_ = err

}

func TestBlockBlockByRefCodeSample(t *testing.T) {
	server := RunCodeSampleServer("blocks", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	blockBlockByRefParams := gocardless.BlockBlockByRefParams{
		ReferenceType:  "customer",
		ReferenceValue: "CU123",
		ReasonType:     "no_intent_to_pay",
	}
	_ = blockBlockByRefParams

	blockBlockByRefResult, err := client.Blocks.BlockByRef(ctx, blockBlockByRefParams)
	_ = blockBlockByRefResult
	_ = err
	for _, block := range blockBlockByRefResult.Blocks {
		fmt.Println(block.Id)
	}

}
