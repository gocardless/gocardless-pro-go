package gocardless

import (
	"context"
	"testing"
)

func TestBillingRequestCreate(t *testing.T) {
	fixtureFile := "testdata/billing_requests.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestCreateParams{}

	o, err :=
		client.BillingRequests.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequest, got nil")

	}
}

func TestBillingRequestCollectCustomerDetails(t *testing.T) {
	fixtureFile := "testdata/billing_requests.json"
	server := runServer(t, fixtureFile, "collect_customer_details")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestCollectCustomerDetailsParams{}

	o, err :=
		client.BillingRequests.CollectCustomerDetails(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequest, got nil")

	}
}

func TestBillingRequestCollectBankAccount(t *testing.T) {
	fixtureFile := "testdata/billing_requests.json"
	server := runServer(t, fixtureFile, "collect_bank_account")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestCollectBankAccountParams{}

	o, err :=
		client.BillingRequests.CollectBankAccount(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequest, got nil")

	}
}

func TestBillingRequestConfirmPayerDetails(t *testing.T) {
	fixtureFile := "testdata/billing_requests.json"
	server := runServer(t, fixtureFile, "confirm_payer_details")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestConfirmPayerDetailsParams{}

	o, err :=
		client.BillingRequests.ConfirmPayerDetails(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequest, got nil")

	}
}

func TestBillingRequestFulfil(t *testing.T) {
	fixtureFile := "testdata/billing_requests.json"
	server := runServer(t, fixtureFile, "fulfil")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestFulfilParams{}

	o, err :=
		client.BillingRequests.Fulfil(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequest, got nil")

	}
}

func TestBillingRequestCancel(t *testing.T) {
	fixtureFile := "testdata/billing_requests.json"
	server := runServer(t, fixtureFile, "cancel")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestCancelParams{}

	o, err :=
		client.BillingRequests.Cancel(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequest, got nil")

	}
}

func TestBillingRequestList(t *testing.T) {
	fixtureFile := "testdata/billing_requests.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestListParams{}

	o, err :=
		client.BillingRequests.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.BillingRequests == nil {

		t.Fatalf("Expected list of BillingRequests, got nil")

	}
}

func TestBillingRequestGet(t *testing.T) {
	fixtureFile := "testdata/billing_requests.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.BillingRequests.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequest, got nil")

	}
}

func TestBillingRequestNotify(t *testing.T) {
	fixtureFile := "testdata/billing_requests.json"
	server := runServer(t, fixtureFile, "notify")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestNotifyParams{}

	o, err :=
		client.BillingRequests.Notify(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequest, got nil")

	}
}

func TestBillingRequestFallback(t *testing.T) {
	fixtureFile := "testdata/billing_requests.json"
	server := runServer(t, fixtureFile, "fallback")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestFallbackParams{}

	o, err :=
		client.BillingRequests.Fallback(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequest, got nil")

	}
}

func TestBillingRequestChooseCurrency(t *testing.T) {
	fixtureFile := "testdata/billing_requests.json"
	server := runServer(t, fixtureFile, "choose_currency")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestChooseCurrencyParams{}

	o, err :=
		client.BillingRequests.ChooseCurrency(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequest, got nil")

	}
}

func TestBillingRequestSelectInstitution(t *testing.T) {
	fixtureFile := "testdata/billing_requests.json"
	server := runServer(t, fixtureFile, "select_institution")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestSelectInstitutionParams{}

	o, err :=
		client.BillingRequests.SelectInstitution(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequest, got nil")

	}
}

func TestBillingRequestCreateWithActions(t *testing.T) {
	fixtureFile := "testdata/billing_requests.json"
	server := runServer(t, fixtureFile, "create_with_actions")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BillingRequestCreateWithActionsParams{}

	o, err :=
		client.BillingRequests.CreateWithActions(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected BillingRequest, got nil")

	}
}
