package gocardless

import (
	"context"
	"testing"
)

func TestCreditorCreate(t *testing.T) {
	fixtureFile := "testdata/creditors.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := CreditorCreateParams{}

	o, err :=
		client.Creditors.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Creditor, got nil")

	}
}

func TestCreditorList(t *testing.T) {
	fixtureFile := "testdata/creditors.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := CreditorListParams{}

	o, err :=
		client.Creditors.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.Creditors == nil {

		t.Fatalf("Expected list of Creditors, got nil")

	}
}

func TestCreditorGet(t *testing.T) {
	fixtureFile := "testdata/creditors.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := CreditorGetParams{}

	o, err :=
		client.Creditors.Get(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Creditor, got nil")

	}
}

func TestCreditorUpdate(t *testing.T) {
	fixtureFile := "testdata/creditors.json"
	server := runServer(t, fixtureFile, "update")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := CreditorUpdateParams{}

	o, err :=
		client.Creditors.Update(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Creditor, got nil")

	}
}
