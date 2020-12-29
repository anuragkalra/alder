package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	debtsURL        = "https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/debts"
	paymentPlansURL = "https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payment_plans"
	paymentsURL     = "https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payments"
)

func main() {
	debts := getDebts()
	//Simple way to print the json array line by line (each line is an array element)
	for _, d := range debts {
		fmt.Println(d)
	}
	fmt.Println("----")

	paymentPlans := getPaymentPlans()
	//fmt.Println(paymentPlans)

	outputDebts(debts, paymentPlans)
	//fmt.Println(debts)
	for _, d := range debts {
		fmt.Println(d)
	}

	fmt.Println("----")

	//addRemainingAmount(debts, paymentPlans)
}

// For testing std out example
func Hello() {
	fmt.Println("hello")
}

// Problem 1
func outputDebts(debts []Debt, paymentPlans []PaymentPlan) error {
	m := make(map[int]int)
	for _, pp := range paymentPlans {
		m[pp.DebtID] = pp.ID
	}

	for i, d := range debts {
		if _, ok := m[d.ID]; ok {
			debts[i].IsInPaymentPlan = true
		}
	}
	return nil
}

// TODO
func addRemainingAmount(debts []Debt, paymentPlans []PaymentPlan, payments []Payment) error {

	return nil
}

//Assumptions

//Debt id = 2 -- Next Due Date
//1/1 + 8 weeks = 2/26
//Since no payment since 8/8,
//We add 2 weeks to the last payment date
func addNextPaymentDue(debts []Debt, paymentPlans []PaymentPlan, payments []Payment) error {
	return nil
}

//Debt represents a struct from the debtsURL endpoint
type Debt struct {
	Amount             float64   `json:"amount"`
	ID                 int       `json:"id"`
	IsInPaymentPlan    bool      `json:"is_in_payment_plan"`
	RemainingAmount    int       `json:"remaining_amount"`      //Problem 3
	NextPaymentDueDate time.Time `json:"next_payment_due_date"` //Problem 4
}

//Payment represents a struct from the paymentsURL endpoint
type Payment struct {
	Amount        float64 `json:"amount"`
	Date          string  `json:"date"`
	PaymentPlanID int     `json:"payment_plan_id"`
}

//PaymentPlan represents a struct from the paymentPlansURL endpoint
type PaymentPlan struct {
	AmountToPay          float64 `json:"amount_to_pay"`
	DebtID               int     `json:"debt_id"`
	ID                   int     `json:"id"`
	InstallmentAmount    float64 `json:"installment_amount"`
	InstallmentFrequency string  `json:"installment_frequency"`
	StartDate            string  `json:"start_date"`
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

func getPaymentPlans() []PaymentPlan {
	resp, err := http.Get(paymentPlansURL)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	paymentPlans := make([]PaymentPlan, 0)
	if err = json.Unmarshal(body, &paymentPlans); err != nil {
		log.Fatalln(err)
	}
	return paymentPlans
}
