package main

import (
	"testing"

	"github.com/anuragkalra/alder/mock"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetPayments(t *testing.T) {
	httpmock.Activate()
	mock.RegisterResponders()
	defer httpmock.DeactivateAndReset()
	payments, err := getPayments()
	assert.NoError(t, err)
	assert.Equal(t, 51.25, payments[0].Amount)
	assert.Equal(t, 0, payments[0].PaymentPlanID)
	assert.Equal(t, "2020-09-29", payments[0].Date)
}
