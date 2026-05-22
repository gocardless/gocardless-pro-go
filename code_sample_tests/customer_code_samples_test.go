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

func TestCustomerCreateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("customers", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	customerCreateParams := gocardless.CustomerCreateParams{
		AddressLine1: "27 Acer Road",
		AddressLine2: "Apt 2",
		City:         "London",
		PostalCode:   "E8 3GX",
		CountryCode:  "GB",
		Email:        "user@example.com",
		GivenName:    "Frank",
		FamilyName:   "Osborne",
	}

	customer, err := client.Customers.Create(ctx, customerCreateParams)

}

func TestCustomerListCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("customers", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	customerListParams := gocardless.CustomerListParams{
		Currency: "GBP",
	}

	customerListResult, err := client.Customers.List(ctx, customerListParams)
	for _, customer := range customerListResult.Customers {
		fmt.Println(customer.GivenName)
	}

}

func TestCustomerGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("customers", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	customer, err := client.Customers.Get(ctx, "CU123")

}

func TestCustomerUpdateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("customers", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	customerUpdateParams := gocardless.CustomerUpdateParams{
		Email: "updated_user@example.com",
	}

	customer, err := client.Customers.Update(ctx, "CU123", customerUpdateParams)

}

func TestCustomerRemoveCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("customers", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	customer, err := client.Customers.Remove(ctx, "CU123", gocardless.CustomerRemoveParams{})

}
