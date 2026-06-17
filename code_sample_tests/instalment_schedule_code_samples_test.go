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
	server := RunCodeSampleServer("instalment_schedules", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

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
	_ = instalmentScheduleCreateWithDatesParams

	requestOption := gocardless.WithIdempotencyKey("random_instalment_schedule_specific_string")
	_ = requestOption
	instalmentSchedule, err := client.InstalmentSchedules.CreateWithDates(ctx, instalmentScheduleCreateWithDatesParams, requestOption)
	_ = instalmentSchedule
	_ = err

}

func TestInstalmentScheduleCreateWithScheduleCodeSample(t *testing.T) {
	server := RunCodeSampleServer("instalment_schedules", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

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
	_ = instalmentScheduleCreateWithScheduleParams

	requestOption := gocardless.WithIdempotencyKey("random_instalment_schedule_specific_string")
	_ = requestOption
	instalmentSchedule, err := client.InstalmentSchedules.CreateWithSchedule(ctx, instalmentScheduleCreateWithScheduleParams, requestOption)
	_ = instalmentSchedule
	_ = err

}

func TestInstalmentScheduleListCodeSample(t *testing.T) {
	server := RunCodeSampleServer("instalment_schedules", true)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	instalmentScheduleListParams := gocardless.InstalmentScheduleListParams{}
	_ = instalmentScheduleListParams
	instalmentScheduleListResult, err := client.InstalmentSchedules.List(ctx, instalmentScheduleListParams)
	_ = instalmentScheduleListResult
	_ = err
	for _, instalmentSchedule := range instalmentScheduleListResult.InstalmentSchedules {
		fmt.Println(instalmentSchedule.Name)
	}

}

func TestInstalmentScheduleGetCodeSample(t *testing.T) {
	server := RunCodeSampleServer("instalment_schedules", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	instalmentSchedule, err := client.InstalmentSchedules.Get(ctx, "IS123")
	_ = instalmentSchedule
	_ = err

}

func TestInstalmentScheduleUpdateCodeSample(t *testing.T) {
	server := RunCodeSampleServer("instalment_schedules", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	instalmentScheduleUpdateParams := gocardless.InstalmentScheduleUpdateParams{
		Metadata: map[string]string{"key": "value"},
	}
	_ = instalmentScheduleUpdateParams

	instalmentSchedule, err := client.InstalmentSchedules.Update(ctx, "IS123", instalmentScheduleUpdateParams)
	_ = instalmentSchedule
	_ = err

}

func TestInstalmentScheduleCancelCodeSample(t *testing.T) {
	server := RunCodeSampleServer("instalment_schedules", false)
	_ = server
	defer server.Close()

	ctx := context.TODO()
	_ = ctx
	client, _ := gocardless.GetClient(t, server.URL)
	_ = client

	instalmentScheduleCancelParams := gocardless.InstalmentScheduleCancelParams{}
	_ = instalmentScheduleCancelParams
	instalmentSchedule, err := client.InstalmentSchedules.Cancel(ctx, "IS123", instalmentScheduleCancelParams)
	_ = instalmentSchedule
	_ = err

}
