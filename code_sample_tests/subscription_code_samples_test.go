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
	server := gocardless.RunCodeSampleServer("subscriptions", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

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

	subscription, err := client.Subscriptions.Create(ctx, subscriptionCreateParams)

}

func TestSubscriptionListCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("subscriptions", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	subscriptionListParams := gocardless.SubscriptionListParams{}
	subscriptionListResult, err := client.Subscriptions.List(ctx, subscriptionListParams)
	for _, subscription := range subscriptionListResult.Subscriptions {
		fmt.Println(subscription.Id)
	}

}

func TestSubscriptionGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("subscriptions", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	subscription, err := client.Subscriptions.Get(ctx, "SB123")

}

func TestSubscriptionUpdateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("subscriptions", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	subscriptionUpdateParams := gocardless.SubscriptionUpdateParams{
		Amount: 4200,
		Name:   "New Name",
	}

	subscription, err := client.Subscriptions.Update(ctx, "SB123", subscriptionUpdateParams)

}

func TestSubscriptionPauseCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("subscriptions", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	subscriptionPauseParams := gocardless.SubscriptionPauseParams{}
	subscription, err := client.Subscriptions.Pause(ctx, "SB123", subscriptionPauseParams)

}

func TestSubscriptionResumeCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("subscriptions", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	subscriptionResumeParams := gocardless.SubscriptionResumeParams{}
	subscription, err := client.Subscriptions.Resume(ctx, "SB123", subscriptionResumeParams)

}

func TestSubscriptionCancelCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("subscriptions", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	subscriptionCancelParams := gocardless.SubscriptionCancelParams{}
	subscription, err := client.Subscriptions.Cancel(ctx, "SB123", subscriptionCancelParams)

}
