package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//Debt represents a struct from the debtsURL endpoint
type Debt struct {
	Amount             float64   `json:"amount"`
	ID                 int       `json:"id"`
	IsInPaymentPlan    bool      `json:"is_in_payment_plan"`
	RemainingAmount    float64   `json:"remaining_amount"`
	NextPaymentDueDate time.Time `json:"next_payment_due_date"`
}

type PrettyDebt struct {
	Amount             float64 `json:"amount"`
	ID                 int     `json:"id"`
	IsInPaymentPlan    bool    `json:"is_in_payment_plan"`
	RemainingAmount    float64 `json:"remaining_amount"`
	NextPaymentDueDate string  `json:"next_payment_due_date"`
}

func getDebts() []Debt {
	resp, err := http.Get(debtsURL)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	debts := make([]Debt, 0)
	if err = json.Unmarshal(body, &debts); err != nil {
		log.Fatalln(err)
	}
	return debts
}
