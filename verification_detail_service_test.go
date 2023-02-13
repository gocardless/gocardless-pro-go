package gocardless

import (
	"context"
	"testing"
)

func TestVerificationDetailCreate(t *testing.T) {
	fixtureFile := "testdata/verification_details.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := VerificationDetailCreateParams{}

	o, err :=
		client.VerificationDetails.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected VerificationDetail, got nil")

	}
}

func TestVerificationDetailList(t *testing.T) {
	fixtureFile := "testdata/verification_details.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := VerificationDetailListParams{}

	o, err :=
		client.VerificationDetails.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.VerificationDetails == nil {

		t.Fatalf("Expected list of VerificationDetails, got nil")

	}
}
