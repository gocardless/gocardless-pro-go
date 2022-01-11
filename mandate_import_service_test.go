package gocardless

import (
	"context"
	"testing"
)

func TestMandateImportCreate(t *testing.T) {
	fixtureFile := "testdata/mandate_imports.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := MandateImportCreateParams{}

	o, err :=
		client.MandateImports.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected MandateImport, got nil")

	}
}

func TestMandateImportGet(t *testing.T) {
	fixtureFile := "testdata/mandate_imports.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := MandateImportGetParams{}

	o, err :=
		client.MandateImports.Get(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected MandateImport, got nil")

	}
}

func TestMandateImportSubmit(t *testing.T) {
	fixtureFile := "testdata/mandate_imports.json"
	server := runServer(t, fixtureFile, "submit")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := MandateImportSubmitParams{}

	o, err :=
		client.MandateImports.Submit(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected MandateImport, got nil")

	}
}

func TestMandateImportCancel(t *testing.T) {
	fixtureFile := "testdata/mandate_imports.json"
	server := runServer(t, fixtureFile, "cancel")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := MandateImportCancelParams{}

	o, err :=
		client.MandateImports.Cancel(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected MandateImport, got nil")

	}
}
