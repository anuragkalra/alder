package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Payment represents a struct from the paymentsURL endpoint
type Payment struct {
	Amount        float64 `json:"amount"`
	Date          string  `json:"date"`
	PaymentPlanID int     `json:"payment_plan_id"`
}

func getPayments() []Payment {
	resp, err := http.Get(paymentsURL)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	payments := make([]Payment, 0)
	if err = json.Unmarshal(body, &payments); err != nil {
		log.Fatalln(err)
	}
	return payments
}
