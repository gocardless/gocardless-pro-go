package gocardless

import (
	"context"
	"testing"
)

func TestCustomerNotificationHandle(t *testing.T) {
	fixtureFile := "testdata/customer_notifications.json"
	server := runServer(t, fixtureFile, "handle")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := CustomerNotificationHandleParams{}

	o, err :=
		client.CustomerNotifications.Handle(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected CustomerNotification, got nil")

	}
}
