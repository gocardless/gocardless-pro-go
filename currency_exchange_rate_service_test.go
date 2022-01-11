package gocardless

import (
	"context"
	"testing"
)

func TestCurrencyExchangeRateList(t *testing.T) {
	fixtureFile := "testdata/currency_exchange_rates.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := CurrencyExchangeRateListParams{}

	o, err :=
		client.CurrencyExchangeRates.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.CurrencyExchangeRates == nil {

		t.Fatalf("Expected list of CurrencyExchangeRates, got nil")

	}
}
