package gocardless

import (
	"context"
	"testing"
)

func TestBillingRequestFlowCreate(t *testing.T) {
	fixtureFile := "testdata/billing_request_flows.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestFlowCreateParams{}

	o, err :=
		client.BillingRequestFlows.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequestFlow, got nil")

	}
}

func TestBillingRequestFlowInitialise(t *testing.T) {
	fixtureFile := "testdata/billing_request_flows.json"
	server := runServer(t, fixtureFile, "initialise")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestFlowInitialiseParams{}

	o, err :=
		client.BillingRequestFlows.Initialise(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequestFlow, got nil")

	}
}
