package gocardless

import (
	"context"
	"testing"
)

func TestPayerThemeCreateForCreditor(t *testing.T) {
	fixtureFile := "testdata/payer_themes.json"
	server := runServer(t, fixtureFile, "create_for_creditor")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := PayerThemeCreateForCreditorParams{}

	o, err :=
		client.PayerThemes.CreateForCreditor(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected PayerTheme, got nil")

	}
}
