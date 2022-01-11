package gocardless

import (
	"context"
	"testing"
)

func TestMandateImportEntryCreate(t *testing.T) {
	fixtureFile := "testdata/mandate_import_entries.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := MandateImportEntryCreateParams{}

	o, err :=
		client.MandateImportEntries.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected MandateImportEntry, got nil")

	}
}

func TestMandateImportEntryList(t *testing.T) {
	fixtureFile := "testdata/mandate_import_entries.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := MandateImportEntryListParams{}

	o, err :=
		client.MandateImportEntries.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.MandateImportEntries == nil {

		t.Fatalf("Expected list of MandateImportEntries, got nil")

	}
}
