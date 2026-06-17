package code_sample_tests // Use a distinct package from the library itself to ensure code samples are tested in the same way as user code

// Code Sample Tests
// These tests verify that the documentation code samples are syntactically valid
// and can execute against a mocked API without errors.
//
// IMPORTANT: These tests do NOT verify business logic - they only verify that
// the code samples compile and execute without syntax errors.

import (
	"context"
	"fmt"
	"testing"

	gocardless "github.com/gocardless/gocardless-pro-go/v6"
)

func TestCurrencyExchangeRateListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("currency_exchange_rates", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	currencyExchangeRateListParams := gocardless.CurrencyExchangeRateListParams{
		Source: "EUR",
		Target: "GBP",
	}
	_ = currencyExchangeRateListParams

	currencyExchangeRateListResult, err := client.CurrencyExchangeRates.List(ctx, currencyExchangeRateListParams)
	_ = currencyExchangeRateListResult
	_ = err
	for _, currencyExchangeRate := range currencyExchangeRateListResult.CurrencyExchangeRates {
		fmt.Println(currencyExchangeRate.Rate)
	}

}
