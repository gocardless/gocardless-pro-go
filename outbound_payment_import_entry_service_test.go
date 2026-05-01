package gocardless

import (
	"context"
	"testing"
)

func TestOutboundPaymentImportEntryList(t *testing.T) {
	fixtureFile := "testdata/outbound_payment_import_entries.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := OutboundPaymentImportEntryListParams{}

	o, err :=
		client.OutboundPaymentImportEntries.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.OutboundPaymentImportEntries == nil {

		t.Fatalf("Expected list of OutboundPaymentImportEntries, got nil")

	}
}
