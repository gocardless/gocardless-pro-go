package gocardless

import (
	"context"
	"testing"
)

func TestOutboundPaymentCreate(t *testing.T) {
	fixtureFile := "testdata/outbound_payments.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := OutboundPaymentCreateParams{}

	o, err :=
		client.OutboundPayments.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected OutboundPayment, got nil")

	}
}

func TestOutboundPaymentWithdraw(t *testing.T) {
	fixtureFile := "testdata/outbound_payments.json"
	server := runServer(t, fixtureFile, "withdraw")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := OutboundPaymentWithdrawParams{}

	o, err :=
		client.OutboundPayments.Withdraw(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected OutboundPayment, got nil")

	}
}

func TestOutboundPaymentCancel(t *testing.T) {
	fixtureFile := "testdata/outbound_payments.json"
	server := runServer(t, fixtureFile, "cancel")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := OutboundPaymentCancelParams{}

	o, err :=
		client.OutboundPayments.Cancel(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected OutboundPayment, got nil")

	}
}

func TestOutboundPaymentApprove(t *testing.T) {
	fixtureFile := "testdata/outbound_payments.json"
	server := runServer(t, fixtureFile, "approve")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := OutboundPaymentApproveParams{}

	o, err :=
		client.OutboundPayments.Approve(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected OutboundPayment, got nil")

	}
}

func TestOutboundPaymentGet(t *testing.T) {
	fixtureFile := "testdata/outbound_payments.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.OutboundPayments.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected OutboundPayment, got nil")

	}
}

func TestOutboundPaymentList(t *testing.T) {
	fixtureFile := "testdata/outbound_payments.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := OutboundPaymentListParams{}

	o, err :=
		client.OutboundPayments.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.OutboundPayments == nil {

		t.Fatalf("Expected list of OutboundPayments, got nil")

	}
}

func TestOutboundPaymentUpdate(t *testing.T) {
	fixtureFile := "testdata/outbound_payments.json"
	server := runServer(t, fixtureFile, "update")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := OutboundPaymentUpdateParams{}

	o, err :=
		client.OutboundPayments.Update(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected OutboundPayment, got nil")

	}
}
