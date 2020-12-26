package main

import (
	"net/http"

	"github.com/jarcoal/httpmock"
)

const mockDebtsURL = "https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/debts"
const mockPaymentPlansURL = "https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payment_plans"
const mockPaymentsURL = "https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payments"

//RegisterResponders registers all HTTP endpoints
func RegisterResponders() {
	httpmock.RegisterResponder(http.MethodGet, mockDebtsURL,
		httpmock.NewStringResponder(http.StatusOK, debts))

	httpmock.RegisterResponder(http.MethodGet, mockPaymentPlansURL,
		httpmock.NewStringResponder(http.StatusOK, paymentPlans))

	httpmock.RegisterResponder(http.MethodGet, mockPaymentsURL,
		httpmock.NewStringResponder(http.StatusOK, payments))
}
