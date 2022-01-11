package gocardless

import (
	"context"
	"testing"
)

func TestPaymentCreate(t *testing.T) {
	fixtureFile := "testdata/payments.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := PaymentCreateParams{}

	o, err :=
		client.Payments.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Payment, got nil")

	}
}

func TestPaymentList(t *testing.T) {
	fixtureFile := "testdata/payments.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := PaymentListParams{}

	o, err :=
		client.Payments.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.Payments == nil {

		t.Fatalf("Expected list of Payments, got nil")

	}
}

func TestPaymentGet(t *testing.T) {
	fixtureFile := "testdata/payments.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.Payments.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Payment, got nil")

	}
}

func TestPaymentUpdate(t *testing.T) {
	fixtureFile := "testdata/payments.json"
	server := runServer(t, fixtureFile, "update")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := PaymentUpdateParams{}

	o, err :=
		client.Payments.Update(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Payment, got nil")

	}
}

func TestPaymentCancel(t *testing.T) {
	fixtureFile := "testdata/payments.json"
	server := runServer(t, fixtureFile, "cancel")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := PaymentCancelParams{}

	o, err :=
		client.Payments.Cancel(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Payment, got nil")

	}
}

func TestPaymentRetry(t *testing.T) {
	fixtureFile := "testdata/payments.json"
	server := runServer(t, fixtureFile, "retry")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := PaymentRetryParams{}

	o, err :=
		client.Payments.Retry(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Payment, got nil")

	}
}
