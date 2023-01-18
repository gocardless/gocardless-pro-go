package gocardless

import (
	"context"
	"testing"
)

func TestSchemeIdentifierList(t *testing.T) {
	fixtureFile := "testdata/scheme_identifiers.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := SchemeIdentifierListParams{}

	o, err :=
		client.SchemeIdentifiers.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.SchemeIdentifiers == nil {

		t.Fatalf("Expected list of SchemeIdentifiers, got nil")

	}
}

func TestSchemeIdentifierGet(t *testing.T) {
	fixtureFile := "testdata/scheme_identifiers.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.SchemeIdentifiers.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected SchemeIdentifier, got nil")

	}
}
