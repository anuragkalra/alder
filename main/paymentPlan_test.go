package main

import (
	"testing"

	"github.com/anuragkalra/alder/mock"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetPaymentPlans(t *testing.T) {
	httpmock.Activate()
	mock.RegisterResponders()
	defer httpmock.DeactivateAndReset()
	paymentPlans, err := getPaymentPlans()
	assert.NoError(t, err)
	assert.Equal(t, 102.5, paymentPlans[0].AmountToPay)
	assert.Equal(t, 0, paymentPlans[0].DebtID)
	assert.Equal(t, 0, paymentPlans[0].ID)
	assert.Equal(t, 51.25, paymentPlans[0].InstallmentAmount)
	assert.Equal(t, "WEEKLY", paymentPlans[0].InstallmentFrequency)
	assert.Equal(t, "2020-09-28", paymentPlans[0].StartDate)
}
