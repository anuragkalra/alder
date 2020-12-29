package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	debtsURL        = "https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/debts"
	paymentPlansURL = "https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payment_plans"
	paymentsURL     = "https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payments"

	layout = "2006-01-02"

	week     = 7 * 24 * time.Hour
	twoWeeks = 2 * week
)

func main() {
	Run()
}

//Run is the method for executing the program
func Run() {
	debts, err := getDebts()
	if err != nil {
		log.Fatalln(err)
	}

	paymentPlans, err := getPaymentPlans()
	if err != nil {
		log.Fatalln(err)
	}

	payments, err := getPayments()
	if err != nil {
		log.Fatalln(err)
	}

	ComputeDebtInfo(debts, paymentPlans, payments)

	for _, d := range debts {
		fmt.Println(d)
	}
}

// ComputeDebtInfo takes as parameter the debts, paymentPlans, and payments and computes the following fields on the debts slice:
// is_in_payment_plan, remaining_amount, next_payment_due_date
func ComputeDebtInfo(debts []Debt, paymentPlans []PaymentPlan, payments []Payment) error {
	//is_in_payment_plan
	if err := updateIsInPaymentPlan(debts, paymentPlans); err != nil {
		return err
	}

	//remaining_amount
	if err := updateRemainingAmount(debts, paymentPlans, payments); err != nil {
		return err
	}

	//next_payment_due_date
	if err := updateNextPaymentDueDate(debts, paymentPlans, payments); err != nil {
		return err
	}

	return nil
}

func updateIsInPaymentPlan(debts []Debt, paymentPlans []PaymentPlan) error {
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

func updateRemainingAmount(debts []Debt, paymentPlans []PaymentPlan, payments []Payment) error {
	for i, d := range debts {
		if d.IsInPaymentPlan {
			atp := paymentPlans[i].AmountToPay
			tp := totalPaid(i, payments)
			debts[i].RemainingAmount = atp - tp
		} else {
			debts[i].RemainingAmount = d.Amount
		}
	}
	return nil
}

//totalPaid is a helper to calculate total paid on a payment plan
func totalPaid(paymentPlanID int, payments []Payment) float64 {
	sum := float64(0)
	for _, p := range payments {
		if p.PaymentPlanID == paymentPlanID {
			sum += p.Amount
		}
	}
	return sum
}

func updateNextPaymentDueDate(debts []Debt, paymentPlans []PaymentPlan, payments []Payment) error {
	for i, d := range debts {
		if d.RemainingAmount == 0 || !d.IsInPaymentPlan {
			debts[i].NextPaymentDueDate = time.Time{} //zero value "null"
		} else {
			lpd := lastPaymentDate(i, payments)
			insfreq := paymentPlans[i].InstallmentFrequency
			if lastPaymentBeforeStartDate(lpd, paymentPlans[i].StartDate) {
				t, _ := time.Parse(layout, paymentPlans[i].StartDate)
				debts[i].NextPaymentDueDate = t
				continue
			}
			if insfreq == "WEEKLY" {
				debts[i].NextPaymentDueDate = lpd.Add(week)
			} else if insfreq == "BI_WEEKLY" {
				debts[i].NextPaymentDueDate = lpd.Add(twoWeeks)
			} else {
				//Create an error: frequency out of range
				return errors.New("Unknown Payment Plan Frequency")
			}
		}
	}

	return nil
}

func lastPaymentDate(debtID int, payments []Payment) time.Time {
	ps := make([]Payment, 0)
	for i, p := range payments {
		if p.PaymentPlanID == debtID {
			ps = append(ps, payments[i])
		}
	}
	ts := time.Time{}
	for _, p := range ps {
		t, _ := time.Parse(layout, p.Date)
		if t.After(ts) {
			ts = t
		}
	}
	return ts
}

func lastPaymentBeforeStartDate(lpd time.Time, paymentPlanStartDate string) bool {
	t, _ := time.Parse(layout, paymentPlanStartDate)
	return t.After(lpd)
}
