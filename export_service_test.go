package gocardless

import (
	"context"
	"testing"
)

func TestExportGet(t *testing.T) {
	fixtureFile := "testdata/exports.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.Exports.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Export, got nil")

	}
}

func TestExportList(t *testing.T) {
	fixtureFile := "testdata/exports.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := ExportListParams{}

	o, err :=
		client.Exports.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.Exports == nil {

		t.Fatalf("Expected list of Exports, got nil")

	}
}
