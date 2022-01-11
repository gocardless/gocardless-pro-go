package gocardless

import (
	"context"
	"testing"
)

func TestPayerAuthorisationGet(t *testing.T) {
	fixtureFile := "testdata/payer_authorisations.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.PayerAuthorisations.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected PayerAuthorisation, got nil")

	}
}

func TestPayerAuthorisationCreate(t *testing.T) {
	fixtureFile := "testdata/payer_authorisations.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := PayerAuthorisationCreateParams{}

	o, err :=
		client.PayerAuthorisations.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected PayerAuthorisation, got nil")

	}
}

func TestPayerAuthorisationUpdate(t *testing.T) {
	fixtureFile := "testdata/payer_authorisations.json"
	server := runServer(t, fixtureFile, "update")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := PayerAuthorisationUpdateParams{}

	o, err :=
		client.PayerAuthorisations.Update(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected PayerAuthorisation, got nil")

	}
}

func TestPayerAuthorisationSubmit(t *testing.T) {
	fixtureFile := "testdata/payer_authorisations.json"
	server := runServer(t, fixtureFile, "submit")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.PayerAuthorisations.Submit(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected PayerAuthorisation, got nil")

	}
}

func TestPayerAuthorisationConfirm(t *testing.T) {
	fixtureFile := "testdata/payer_authorisations.json"
	server := runServer(t, fixtureFile, "confirm")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.PayerAuthorisations.Confirm(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected PayerAuthorisation, got nil")

	}
}
