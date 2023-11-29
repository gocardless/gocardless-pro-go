package gocardless

import (
	"context"
	"testing"
)

func TestTransferredMandateTransferredMandates(t *testing.T) {
	fixtureFile := "testdata/transferred_mandates.json"
	server := runServer(t, fixtureFile, "transferred_mandates")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := TransferredMandateTransferredMandatesParams{}

	o, err :=
		client.TransferredMandates.TransferredMandates(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected TransferredMandate, got nil")

	}
}
