package gocardless

import (
	"context"
	"testing"
)

func TestInstitutionList(t *testing.T) {
	fixtureFile := "testdata/institutions.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := InstitutionListParams{}

	o, err :=
		client.Institutions.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.Institutions == nil {

		t.Fatalf("Expected list of Institutions, got nil")

	}
}
