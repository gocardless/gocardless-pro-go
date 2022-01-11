package gocardless

import (
	"context"
	"testing"
)

func TestMandateCreate(t *testing.T) {
	fixtureFile := "testdata/mandates.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := MandateCreateParams{}

	o, err :=
		client.Mandates.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Mandate, got nil")

	}
}

func TestMandateList(t *testing.T) {
	fixtureFile := "testdata/mandates.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := MandateListParams{}

	o, err :=
		client.Mandates.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.Mandates == nil {

		t.Fatalf("Expected list of Mandates, got nil")

	}
}

func TestMandateGet(t *testing.T) {
	fixtureFile := "testdata/mandates.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.Mandates.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Mandate, got nil")

	}
}

func TestMandateUpdate(t *testing.T) {
	fixtureFile := "testdata/mandates.json"
	server := runServer(t, fixtureFile, "update")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := MandateUpdateParams{}

	o, err :=
		client.Mandates.Update(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Mandate, got nil")

	}
}

func TestMandateCancel(t *testing.T) {
	fixtureFile := "testdata/mandates.json"
	server := runServer(t, fixtureFile, "cancel")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := MandateCancelParams{}

	o, err :=
		client.Mandates.Cancel(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Mandate, got nil")

	}
}

func TestMandateReinstate(t *testing.T) {
	fixtureFile := "testdata/mandates.json"
	server := runServer(t, fixtureFile, "reinstate")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := MandateReinstateParams{}

	o, err :=
		client.Mandates.Reinstate(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Mandate, got nil")

	}
}
