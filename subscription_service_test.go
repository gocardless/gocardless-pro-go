package gocardless

import (
	"context"
	"testing"
)

func TestSubscriptionCreate(t *testing.T) {
	fixtureFile := "testdata/subscriptions.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := SubscriptionCreateParams{}

	o, err :=
		client.Subscriptions.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Subscription, got nil")

	}
}

func TestSubscriptionList(t *testing.T) {
	fixtureFile := "testdata/subscriptions.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := SubscriptionListParams{}

	o, err :=
		client.Subscriptions.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.Subscriptions == nil {

		t.Fatalf("Expected list of Subscriptions, got nil")

	}
}

func TestSubscriptionGet(t *testing.T) {
	fixtureFile := "testdata/subscriptions.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.Subscriptions.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Subscription, got nil")

	}
}

func TestSubscriptionUpdate(t *testing.T) {
	fixtureFile := "testdata/subscriptions.json"
	server := runServer(t, fixtureFile, "update")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := SubscriptionUpdateParams{}

	o, err :=
		client.Subscriptions.Update(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Subscription, got nil")

	}
}

func TestSubscriptionPause(t *testing.T) {
	fixtureFile := "testdata/subscriptions.json"
	server := runServer(t, fixtureFile, "pause")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := SubscriptionPauseParams{}

	o, err :=
		client.Subscriptions.Pause(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Subscription, got nil")

	}
}

func TestSubscriptionResume(t *testing.T) {
	fixtureFile := "testdata/subscriptions.json"
	server := runServer(t, fixtureFile, "resume")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := SubscriptionResumeParams{}

	o, err :=
		client.Subscriptions.Resume(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Subscription, got nil")

	}
}

func TestSubscriptionCancel(t *testing.T) {
	fixtureFile := "testdata/subscriptions.json"
	server := runServer(t, fixtureFile, "cancel")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := SubscriptionCancelParams{}

	o, err :=
		client.Subscriptions.Cancel(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Subscription, got nil")

	}
}
