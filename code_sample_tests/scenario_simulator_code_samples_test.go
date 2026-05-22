package code_sample_tests // Use a distinct package from the library itself to ensure code samples are tested in the same way as user code

// Code Sample Tests
// These tests verify that the documentation code samples are syntactically valid
// and can execute against a mocked API without errors.
//
// IMPORTANT: These tests do NOT verify business logic - they only verify that
// the code samples compile and execute without syntax errors.

import (
	"context"
	"testing"

	gocardless "github.com/gocardless/gocardless-pro-go/v6"
)

func TestScenarioSimulatorRunCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("scenario_simulators", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	scenarioSimulatorRunParams := gocardless.ScenarioSimulatorRunParams{
		Links: &gocardless.ScenarioSimulatorRunParamsLinks{
			Resource: "PM123",
		},
	}

	action := "payment_failed"
	scenarioSimulator, err := client.ScenarioSimulators.Run(ctx, action, scenarioSimulatorRunParams)

}
