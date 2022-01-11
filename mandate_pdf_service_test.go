package gocardless

import (
	"context"
	"testing"
)

func TestMandatePdfCreate(t *testing.T) {
	fixtureFile := "testdata/mandate_pdfs.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := MandatePdfCreateParams{}

	o, err :=
		client.MandatePdfs.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected MandatePdf, got nil")

	}
}
