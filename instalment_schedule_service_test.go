package gocardless

import (
	"context"
	"testing"
)

func TestInstalmentScheduleCreateWithDates(t *testing.T) {
	fixtureFile := "testdata/instalment_schedules.json"
	server := runServer(t, fixtureFile, "create_with_dates")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := InstalmentScheduleCreateWithDatesParams{}

	o, err :=
		client.InstalmentSchedules.CreateWithDates(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected InstalmentSchedule, got nil")

	}
}

func TestInstalmentScheduleCreateWithSchedule(t *testing.T) {
	fixtureFile := "testdata/instalment_schedules.json"
	server := runServer(t, fixtureFile, "create_with_schedule")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := InstalmentScheduleCreateWithScheduleParams{}

	o, err :=
		client.InstalmentSchedules.CreateWithSchedule(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected InstalmentSchedule, got nil")

	}
}

func TestInstalmentScheduleList(t *testing.T) {
	fixtureFile := "testdata/instalment_schedules.json"
	server := runServer(t, fixtureFile, "list")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := InstalmentScheduleListParams{}

	o, err :=
		client.InstalmentSchedules.List(
			ctx, p)

	if err != nil {
		t.Fatal(err)
	}

	if o.InstalmentSchedules == nil {

		t.Fatalf("Expected list of InstalmentSchedules, got nil")

	}
}

func TestInstalmentScheduleGet(t *testing.T) {
	fixtureFile := "testdata/instalment_schedules.json"
	server := runServer(t, fixtureFile, "get")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	o, err :=
		client.InstalmentSchedules.Get(
			ctx, "ID123")

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected InstalmentSchedule, got nil")

	}
}

func TestInstalmentScheduleUpdate(t *testing.T) {
	fixtureFile := "testdata/instalment_schedules.json"
	server := runServer(t, fixtureFile, "update")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := InstalmentScheduleUpdateParams{}

	o, err :=
		client.InstalmentSchedules.Update(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected InstalmentSchedule, got nil")

	}
}

func TestInstalmentScheduleCancel(t *testing.T) {
	fixtureFile := "testdata/instalment_schedules.json"
	server := runServer(t, fixtureFile, "cancel")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := InstalmentScheduleCancelParams{}

	o, err :=
		client.InstalmentSchedules.Cancel(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected InstalmentSchedule, got nil")

	}
}
