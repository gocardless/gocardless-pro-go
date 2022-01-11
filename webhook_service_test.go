package gocardless

import (
	"context"
	"testing"
)

func TestWebhookList(t *testing.T) {
	fixtureFile := "testdata/webhooks.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := WebhookListParams{}

	o, err :=
		client.Webhooks.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.Webhooks == nil {

		t.Fatalf("Expected list of Webhooks, got nil")

	}
}

func TestWebhookGet(t *testing.T) {
	fixtureFile := "testdata/webhooks.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.Webhooks.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Webhook, got nil")

	}
}

func TestWebhookRetry(t *testing.T) {
	fixtureFile := "testdata/webhooks.json"
	server := runServer(t, fixtureFile, "retry")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.Webhooks.Retry(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Webhook, got nil")

	}
}
