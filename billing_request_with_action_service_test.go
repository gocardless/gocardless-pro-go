package gocardless

import (
	"context"
	"testing"
)

func TestBillingRequestWithActionCreateWithActions(t *testing.T) {
	fixtureFile := "testdata/billing_request_with_actions.json"
	server := runServer(t, fixtureFile, "create_with_actions")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestWithActionCreateWithActionsParams{}

	o, err :=
		client.BillingRequestWithActions.CreateWithActions(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequestWithAction, got nil")

	}
}
