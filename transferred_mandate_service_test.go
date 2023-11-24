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

	o, err :=
		client.TransferredMandates.TransferredMandates(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected TransferredMandate, got nil")

	}
}
