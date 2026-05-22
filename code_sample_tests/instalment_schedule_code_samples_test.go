package code_sample_tests // Use a distinct package from the library itself to ensure code samples are tested in the same way as user code

// Code Sample Tests
// These tests verify that the documentation code samples are syntactically valid
// and can execute against a mocked API without errors.
//
// IMPORTANT: These tests do NOT verify business logic - they only verify that
// the code samples compile and execute without syntax errors.

import (
	"context"
	"fmt"
	"testing"

	gocardless "github.com/gocardless/gocardless-pro-go/v6"
)

func TestInstalmentScheduleCreateWithDatesCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("instalment_schedules", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	instalmentScheduleCreateWithDatesParams := gocardless.InstalmentScheduleCreateWithDatesParams{
		TotalAmount: 10000,
		AppFee:      10,
		Currency:    "GBP",
		Instalments: []gocardless.InstalmentScheduleCreateWithDatesParamsInstalments{
			{
				Amount:     3400,
				ChargeDate: "2019-08-20",
			},
			{
				Amount:     3400,
				ChargeDate: "2019-09-03",
			},
			{
				Amount:     3400,
				ChargeDate: "2019-09-17",
			},
		},
		Links: gocardless.InstalmentScheduleCreateWithDatesParamsLinks{
			Mandate: "MD123",
		},
		Metadata: map[string]string{"invoiceId": "001"},
	}

	requestOption := gocardless.WithIdempotencyKey("random_instalment_schedule_specific_string")
	instalmentSchedule, err := client.InstalmentSchedules.CreateWithDates(ctx, instalmentScheduleCreateWithDatesParams, requestOption)

}

func TestInstalmentScheduleCreateWithScheduleCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("instalment_schedules", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	instalmentScheduleCreateWithScheduleParams := gocardless.InstalmentScheduleCreateWithScheduleParams{
		TotalAmount: 10000,
		AppFee:      10,
		Currency:    "GBP",
		Instalments: gocardless.InstalmentScheduleCreateWithScheduleParamsInstalments{
			Interval:     2,
			IntervalUnit: "monthly",
			Amounts:      []int{3400, 3400, 3200},
		},
		Metadata: map[string]string{"invoiceId": "001"},
	}

	requestOption := gocardless.WithIdempotencyKey("random_instalment_schedule_specific_string")
	instalmentSchedule, err := client.InstalmentSchedules.CreateWithSchedule(ctx, instalmentScheduleCreateWithScheduleParams, requestOption)

}

func TestInstalmentScheduleListCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("instalment_schedules", true)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	instalmentScheduleListParams := gocardless.InstalmentScheduleListParams{}
	instalmentScheduleListResult, err := client.InstalmentSchedules.List(ctx, instalmentScheduleListParams)
	for _, instalmentSchedule := range instalmentScheduleListResult.InstalmentSchedules {
		fmt.Println(instalmentSchedule.Name)
	}

}

func TestInstalmentScheduleGetCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("instalment_schedules", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	instalmentSchedule, err := client.InstalmentSchedules.Get(ctx, "IS123")

}

func TestInstalmentScheduleUpdateCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("instalment_schedules", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	instalmentScheduleUpdateParams := gocardless.InstalmentScheduleUpdateParams{
		Metadata: map[string]string{"key": "value"},
	}

	instalmentSchedule, err := client.InstalmentSchedules.Update(ctx, "IS123", instalmentScheduleUpdateParams)

}

func TestInstalmentScheduleCancelCodeSample(t *testing.T) {
	server := gocardless.RunCodeSampleServer("instalment_schedules", false)
	defer server.Close()

	ctx := context.TODO()
	client, _ := gocardless.GetClient(t, server.URL)

	instalmentScheduleCancelParams := gocardless.InstalmentScheduleCancelParams{}
	instalmentSchedule, err := client.InstalmentSchedules.Cancel(ctx, "IS123", instalmentScheduleCancelParams)

}
