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

func TestSubscriptionCreateCodeSample(t *testing.T) {
	server := RunCodeSampleServer("subscriptions", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	subscriptionCreateParams := gocardless.SubscriptionCreateParams{
		Amount:       1500, // 15 GBP in pence
		Currency:     "GBP",
		IntervalUnit: "monthly",
		DayOfMonth:   5,
		Links: gocardless.SubscriptionCreateParamsLinks{
			Mandate: "MD123",
		},
		Metadata: map[string]string{"subscription_number": "ABC123"},
	}
	_ = subscriptionCreateParams

	subscription, err := client.Subscriptions.Create(ctx, subscriptionCreateParams)
	_ = subscription
	_ = err

}

func TestSubscriptionListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("subscriptions", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	subscriptionListParams := gocardless.SubscriptionListParams{}
	_ = subscriptionListParams
	subscriptionListResult, err := client.Subscriptions.List(ctx, subscriptionListParams)
	_ = subscriptionListResult
	_ = err
	for _, subscription := range subscriptionListResult.Subscriptions {
		fmt.Println(subscription.Id)
	}

}

func TestSubscriptionGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("subscriptions", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	subscription, err := client.Subscriptions.Get(ctx, "SB123")
	_ = subscription
	_ = err

}

func TestSubscriptionUpdateCodeSample(t *testing.T) {
	server := RunCodeSampleServer("subscriptions", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	subscriptionUpdateParams := gocardless.SubscriptionUpdateParams{
		Amount: 4200,
		Name:   "New Name",
	}
	_ = subscriptionUpdateParams

	subscription, err := client.Subscriptions.Update(ctx, "SB123", subscriptionUpdateParams)
	_ = subscription
	_ = err

}

func TestSubscriptionPauseCodeSample(t *testing.T) {
	server := RunCodeSampleServer("subscriptions", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	subscriptionPauseParams := gocardless.SubscriptionPauseParams{}
	_ = subscriptionPauseParams
	subscription, err := client.Subscriptions.Pause(ctx, "SB123", subscriptionPauseParams)
	_ = subscription
	_ = err

}

func TestSubscriptionResumeCodeSample(t *testing.T) {
	server := RunCodeSampleServer("subscriptions", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	subscriptionResumeParams := gocardless.SubscriptionResumeParams{}
	_ = subscriptionResumeParams
	subscription, err := client.Subscriptions.Resume(ctx, "SB123", subscriptionResumeParams)
	_ = subscription
	_ = err

}

func TestSubscriptionCancelCodeSample(t *testing.T) {
	server := RunCodeSampleServer("subscriptions", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	subscriptionCancelParams := gocardless.SubscriptionCancelParams{}
	_ = subscriptionCancelParams
	subscription, err := client.Subscriptions.Cancel(ctx, "SB123", subscriptionCancelParams)
	_ = subscription
	_ = err

}
