package gocardless

import (
	"context"
	"testing"
)

func TestPayoutItemList(t *testing.T) {
	fixtureFile := "testdata/payout_items.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := PayoutItemListParams{}

	o, err :=
		client.PayoutItems.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.PayoutItems == nil {

		t.Fatalf("Expected list of PayoutItems, got nil")

	}
}
