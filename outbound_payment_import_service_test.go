package gocardless

import (
	"context"
	"testing"
)

func TestOutboundPaymentImportCreate(t *testing.T) {
	fixtureFile := "testdata/outbound_payment_imports.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := OutboundPaymentImportCreateParams{}

	o, err :=
		client.OutboundPaymentImports.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected OutboundPaymentImport, got nil")

	}
}

func TestOutboundPaymentImportGet(t *testing.T) {
	fixtureFile := "testdata/outbound_payment_imports.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := OutboundPaymentImportGetParams{}

	o, err :=
		client.OutboundPaymentImports.Get(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected OutboundPaymentImport, got nil")

	}
}

func TestOutboundPaymentImportList(t *testing.T) {
	fixtureFile := "testdata/outbound_payment_imports.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := OutboundPaymentImportListParams{}

	o, err :=
		client.OutboundPaymentImports.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.OutboundPaymentImports == nil {

		t.Fatalf("Expected list of OutboundPaymentImports, got nil")

	}
}
