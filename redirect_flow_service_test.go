package gocardless

import (
	"context"
	"testing"
)

func TestRedirectFlowCreate(t *testing.T) {
	fixtureFile := "testdata/redirect_flows.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := RedirectFlowCreateParams{}

	o, err :=
		client.RedirectFlows.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected RedirectFlow, got nil")

	}
}

func TestRedirectFlowGet(t *testing.T) {
	fixtureFile := "testdata/redirect_flows.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.RedirectFlows.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected RedirectFlow, got nil")

	}
}

func TestRedirectFlowComplete(t *testing.T) {
	fixtureFile := "testdata/redirect_flows.json"
	server := runServer(t, fixtureFile, "complete")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := RedirectFlowCompleteParams{}

	o, err :=
		client.RedirectFlows.Complete(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected RedirectFlow, got nil")

	}
}
