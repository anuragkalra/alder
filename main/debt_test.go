package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/anuragkalra/alder/mock"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetDebts(t *testing.T) {
	httpmock.Activate()
	mock.RegisterResponders()
	defer httpmock.DeactivateAndReset()
	debts, err := getDebts()
	assert.NoError(t, err)
	assert.Equal(t, 123.46, debts[0].Amount)
	assert.Equal(t, 0, debts[0].ID)
}

func TestDebtString(t *testing.T) {
	var d Debt
	{
		d.Amount = float64(34.2)
		d.ID = 2
		d.IsInPaymentPlan = true
		d.RemainingAmount = 23.1
		d.NextPaymentDueDate = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

		s := fmt.Sprint(d)
		assert.Equal(t, "{34.20 2 true 23.10 2009-11-10 23:00:00 +0000 UTC}", s)
	}

	{
		d.Amount = float64(39.9)
		d.ID = 2
		d.IsInPaymentPlan = false
		d.RemainingAmount = 25.7
		d.NextPaymentDueDate = time.Time{}

		s := fmt.Sprint(d)
		assert.Equal(t, "{39.90 2 false 25.70 null}", s)
	}

}
