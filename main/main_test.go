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

// MAY NEED TO MODIFY
func TestAddRemainingAmount(t *testing.T) {
	httpmock.Activate()
	RegisterResponders()
	defer httpmock.DeactivateAndReset()

	debts := getDebts()
	paymentPlans := getPaymentPlans()
	payments := getPayments()

	err := addRemainingAmount(debts, paymentPlans, payments)

	assert.NoError(t, err)
	assert.Equal(t, 0, debts[0].RemainingAmount)
	assert.Equal(t, 50, debts[1].RemainingAmount)
	assert.Equal(t, 607.67, debts[2].RemainingAmount)
	assert.Equal(t, 622.415, debts[3].RemainingAmount)
	assert.Equal(t, 9238.02, debts[4].RemainingAmount)
}

func TestAddNextPaymentDue(t *testing.T) {
	httpmock.Activate()
	RegisterResponders()
	defer httpmock.DeactivateAndReset()

	debts := getDebts()
	paymentPlans := getPaymentPlans()
	payments := getPayments()

	err := addNextPaymentDue(debts, paymentPlans, payments)

	assert.NoError(t, err)

	//TODO Move dates to ISO 8601 UTC Date Format
	assert.Equal(t, nil, debts[0].NextPaymentDueDate)    //Debt has been paid off
	assert.Equal(t, "8/15", debts[1].NextPaymentDueDate) //(8/8) + 1 Week
	assert.Equal(t, "8/22", debts[2].NextPaymentDueDate) //(8/8) + 2 Weeks
	//Payment Plan Start Date. Already made 3 payments but hasnt reached total
	assert.Equal(t, "8/1", debts[3].NextPaymentDueDate)
	assert.Equal(t, nil, debts[4].NextPaymentDueDate) //No Payment Plan
}
