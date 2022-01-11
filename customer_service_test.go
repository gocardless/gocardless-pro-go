package gocardless

import (
	"context"
	"testing"
)

func TestCustomerCreate(t *testing.T) {
	fixtureFile := "testdata/customers.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := CustomerCreateParams{}

	o, err :=
		client.Customers.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Customer, got nil")

	}
}

func TestCustomerList(t *testing.T) {
	fixtureFile := "testdata/customers.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := CustomerListParams{}

	o, err :=
		client.Customers.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.Customers == nil {

		t.Fatalf("Expected list of Customers, got nil")

	}
}

func TestCustomerGet(t *testing.T) {
	fixtureFile := "testdata/customers.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.Customers.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Customer, got nil")

	}
}

func TestCustomerUpdate(t *testing.T) {
	fixtureFile := "testdata/customers.json"
	server := runServer(t, fixtureFile, "update")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := CustomerUpdateParams{}

	o, err :=
		client.Customers.Update(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Customer, got nil")

	}
}

func TestCustomerRemove(t *testing.T) {
	fixtureFile := "testdata/customers.json"
	server := runServer(t, fixtureFile, "remove")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := CustomerRemoveParams{}

	o, err :=
		client.Customers.Remove(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Customer, got nil")

	}
}
