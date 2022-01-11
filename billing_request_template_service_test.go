package gocardless

import (
	"context"
	"testing"
)

func TestBillingRequestTemplateList(t *testing.T) {
	fixtureFile := "testdata/billing_request_templates.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestTemplateListParams{}

	o, err :=
		client.BillingRequestTemplates.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.BillingRequestTemplates == nil {

		t.Fatalf("Expected list of BillingRequestTemplates, got nil")

	}
}

func TestBillingRequestTemplateGet(t *testing.T) {
	fixtureFile := "testdata/billing_request_templates.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.BillingRequestTemplates.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequestTemplate, got nil")

	}
}

func TestBillingRequestTemplateCreate(t *testing.T) {
	fixtureFile := "testdata/billing_request_templates.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestTemplateCreateParams{}

	o, err :=
		client.BillingRequestTemplates.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequestTemplate, got nil")

	}
}

func TestBillingRequestTemplateUpdate(t *testing.T) {
	fixtureFile := "testdata/billing_request_templates.json"
	server := runServer(t, fixtureFile, "update")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestTemplateUpdateParams{}

	o, err :=
		client.BillingRequestTemplates.Update(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequestTemplate, got nil")

	}
}
