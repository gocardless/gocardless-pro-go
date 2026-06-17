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

func TestWebhookListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("webhooks", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	webhookListParams := gocardless.WebhookListParams{
		CreatedAt: &gocardless.WebhookListParamsCreatedAt{
			Gt: "2020-01-01T17:01:06.000Z",
		},
	}
	_ = webhookListParams

	webhookListResult, err := client.Webhooks.List(ctx, webhookListParams)
	_ = webhookListResult
	_ = err
	for _, webhook := range webhookListResult.Webhooks {
		fmt.Println(webhook.Id)
	}

}

func TestWebhookGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("webhooks", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	webhook, err := client.Webhooks.Get(ctx, "WB123")
	_ = webhook
	_ = err

}

func TestWebhookRetryCodeSample(t *testing.T) {
	server := RunCodeSampleServer("webhooks", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	webhook, err := client.Webhooks.Retry(ctx, "WB123")
	_ = webhook
	_ = err

}
