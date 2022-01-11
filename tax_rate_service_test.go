package gocardless

import (
	"context"
	"testing"
)

func TestTaxRateList(t *testing.T) {
	fixtureFile := "testdata/tax_rates.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := TaxRateListParams{}

	o, err :=
		client.TaxRates.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.TaxRates == nil {

		t.Fatalf("Expected list of TaxRates, got nil")

	}
}

func TestTaxRateGet(t *testing.T) {
	fixtureFile := "testdata/tax_rates.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.TaxRates.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected TaxRate, got nil")

	}
}
