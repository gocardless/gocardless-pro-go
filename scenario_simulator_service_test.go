package gocardless

import (
	"context"
	"testing"
)

func TestScenarioSimulatorRun(t *testing.T) {
	fixtureFile := "testdata/scenario_simulators.json"
	server := runServer(t, fixtureFile, "run")
	defer server.Close()

	ctx := context.TODO()
	client, err := getClient(t, server.URL)
	if err != nil {
		t.Fatal(err)
	}

	p := ScenarioSimulatorRunParams{}

	o, err :=
		client.ScenarioSimulators.Run(
			ctx, "ID123", p)

	if err != nil {
		t.Fatal(err)
	}

	if o == nil {

		t.Fatalf("Expected ScenarioSimulator, got nil")

	}
}
