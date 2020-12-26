package main

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// For testing std out
func ExampleHello() {
	Hello()

	// Output:
	// hello
}

func TestGetDebts(t *testing.T) {
	httpmock.Activate()
	RegisterResponders()
	defer httpmock.DeactivateAndReset()
	debts := getDebts()
	assert.Equal(t, 123.46, debts[0].Amount)
	assert.Equal(t, 0, debts[0].ID)
}

func TestGetPayments(t *testing.T) {
	httpmock.Activate()
	RegisterResponders()
	defer httpmock.DeactivateAndReset()
	payments := getPayments()
	assert.Equal(t, 51.25, payments[0].Amount)
	assert.Equal(t, 0, payments[0].PaymentPlanID)
	assert.Equal(t, "2020-09-29", payments[0].Date)
}

func TestGetPaymentPlans(t *testing.T) {
	httpmock.Activate()
	RegisterResponders()
	defer httpmock.DeactivateAndReset()
	paymentPlans := getPaymentPlans()
	assert.Equal(t, 102.5, paymentPlans[0].AmountToPay)
	assert.Equal(t, 0, paymentPlans[0].DebtID)
	assert.Equal(t, 0, paymentPlans[0].ID)
	assert.Equal(t, 51.25, paymentPlans[0].InstallmentAmount)
	assert.Equal(t, "WEEKLY", paymentPlans[0].InstallmentFrequency)
	assert.Equal(t, "2020-09-28", paymentPlans[0].StartDate)
}

func TestOutputDebts(t *testing.T) {
	httpmock.Activate()
	RegisterResponders()
	defer httpmock.DeactivateAndReset()

	debts := getDebts()
	paymentPlans := getPaymentPlans()

	assert.Equal(t, false, debts[0].IsInPaymentPlan)
	assert.Equal(t, false, debts[1].IsInPaymentPlan)
	assert.Equal(t, false, debts[2].IsInPaymentPlan)
	assert.Equal(t, false, debts[3].IsInPaymentPlan)
	assert.Equal(t, false, debts[4].IsInPaymentPlan)

	err := outputDebts(debts, paymentPlans)

	assert.NoError(t, err)
	assert.Equal(t, true, debts[0].IsInPaymentPlan)
	assert.Equal(t, true, debts[1].IsInPaymentPlan)
	assert.Equal(t, true, debts[2].IsInPaymentPlan)
	assert.Equal(t, true, debts[3].IsInPaymentPlan)
	assert.Equal(t, false, debts[4].IsInPaymentPlan)
}
