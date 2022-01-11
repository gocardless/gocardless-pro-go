package gocardless

import (
	"context"
	"testing"
)

func TestBlockCreate(t *testing.T) {
	fixtureFile := "testdata/blocks.json"
	server := runServer(t, fixtureFile, "create")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BlockCreateParams{}

	o, err :=
		client.Blocks.Create(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Block, got nil")

	}
}

func TestBlockGet(t *testing.T) {
	fixtureFile := "testdata/blocks.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.Blocks.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Block, got nil")

	}
}

func TestBlockList(t *testing.T) {
	fixtureFile := "testdata/blocks.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BlockListParams{}

	o, err :=
		client.Blocks.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.Blocks == nil {

		t.Fatalf("Expected list of Blocks, got nil")

	}
}

func TestBlockDisable(t *testing.T) {
	fixtureFile := "testdata/blocks.json"
	server := runServer(t, fixtureFile, "disable")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.Blocks.Disable(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Block, got nil")

	}
}

func TestBlockEnable(t *testing.T) {
	fixtureFile := "testdata/blocks.json"
	server := runServer(t, fixtureFile, "enable")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.Blocks.Enable(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Block, got nil")

	}
}

func TestBlockBlockByRef(t *testing.T) {
	fixtureFile := "testdata/blocks.json"
	server := runServer(t, fixtureFile, "block_by_ref")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := BlockBlockByRefParams{}

	o, err :=
		client.Blocks.BlockByRef(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected Block, got nil")

	}
}
