package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//PaymentPlan represents a struct from the paymentPlansURL endpoint
type PaymentPlan struct {
	AmountToPay          float64 `json:"amount_to_pay"`
	DebtID               int     `json:"debt_id"`
	ID                   int     `json:"id"`
	InstallmentAmount    float64 `json:"installment_amount"`
	InstallmentFrequency string  `json:"installment_frequency"`
	StartDate            string  `json:"start_date"`
}

func getPaymentPlans() ([]PaymentPlan, error) {
	resp, err := http.Get(paymentPlansURL)
	if err != nil {
		return nil, err
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
		log.Fatalln(err)
	}
	paymentPlans := make([]PaymentPlan, 0)
	if err = json.Unmarshal(body, &paymentPlans); err != nil {
		return nil, err
		log.Fatalln(err)
	}
	return paymentPlans, nil
}
