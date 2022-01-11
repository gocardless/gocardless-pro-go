package gocardless

import (
	"context"
	"testing"
)

func TestEventList(t *testing.T) {
	fixtureFile := "testdata/events.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := EventListParams{}

	o, err :=
		client.Events.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.Events == nil {

		t.Fatalf("Expected list of Events, got nil")

	}
}

func TestEventGet(t *testing.T) {
	fixtureFile := "testdata/events.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.Events.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Event, got nil")

	}
}
