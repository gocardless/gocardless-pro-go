package gocardless

import (
	"context"
	"testing"
)

func TestRefundCreate(t *testing.T) {
	fixtureFile := "testdata/refunds.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := RefundCreateParams{}

	o, err :=
		client.Refunds.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Refund, got nil")

	}
}

func TestRefundList(t *testing.T) {
	fixtureFile := "testdata/refunds.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := RefundListParams{}

	o, err :=
		client.Refunds.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.Refunds == nil {

		t.Fatalf("Expected list of Refunds, got nil")

	}
}

func TestRefundGet(t *testing.T) {
	fixtureFile := "testdata/refunds.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.Refunds.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Refund, got nil")

	}
}

func TestRefundUpdate(t *testing.T) {
	fixtureFile := "testdata/refunds.json"
	server := runServer(t, fixtureFile, "update")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := RefundUpdateParams{}

	o, err :=
		client.Refunds.Update(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Refund, got nil")

	}
}
