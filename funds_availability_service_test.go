package gocardless

import (
	"context"
	"testing"
)

func TestFundsAvailabilityCheck(t *testing.T) {
	fixtureFile := "testdata/funds_availabilities.json"
	server := runServer(t, fixtureFile, "check")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := FundsAvailabilityCheckParams{}

	o, err :=
		client.FundsAvailabilities.Check(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected FundsAvailability, got nil")

	}
}
