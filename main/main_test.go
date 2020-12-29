package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/anuragkalra/alder/mock"
	"github.com/jarcoal/httpmock"
)

func ExampleRun() {
	Run()

	// Output:
	// {123.46 0 true 0.00 null}
	// {100.00 1 true 50.00 2020-08-15 00:00:00 +0000 UTC}
	// {4920.34 2 true 607.67 2020-08-22 00:00:00 +0000 UTC}
	// {12938.00 3 true 622.41 2020-08-01 00:00:00 +0000 UTC}
	// {9238.02 4 false 9238.02 null}
}

func TestComputeDebtInfo(t *testing.T) {
	httpmock.Activate()
	mock.RegisterResponders()
	defer httpmock.DeactivateAndReset()

	debts, err := getDebts()
	assert.NoError(t, err)
	paymentPlans, err := getPaymentPlans()
	assert.NoError(t, err)
	payments, err := getPayments()
	assert.NoError(t, err)

	err = ComputeDebtInfo(debts, paymentPlans, payments)
	assert.NoError(t, err)
	assert.Equal(t, 123.46, debts[0].Amount)
	assert.Equal(t, 0, debts[0].ID)
	assert.Equal(t, true, debts[0].IsInPaymentPlan)
	assert.Equal(t, float64(0), debts[0].RemainingAmount)
	assert.Equal(t, time.Time{}, debts[0].NextPaymentDueDate)

	assert.NoError(t, err)
	assert.Equal(t, 100.00, debts[1].Amount)
	assert.Equal(t, 1, debts[1].ID)
	assert.Equal(t, true, debts[1].IsInPaymentPlan)
	assert.Equal(t, float64(50.00), debts[1].RemainingAmount)
	assert.Equal(t, time.Date(2020, time.August, 15, 0, 0, 0, 0, time.UTC), debts[1].NextPaymentDueDate)
}
