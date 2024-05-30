package gocardless

import (
	"context"
	"testing"
)

func TestLogoCreateForCreditor(t *testing.T) {
	fixtureFile := "testdata/logos.json"
	server := runServer(t, fixtureFile, "create_for_creditor")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := LogoCreateForCreditorParams{}

	o, err :=
		client.Logos.CreateForCreditor(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Logo, got nil")

	}
}
