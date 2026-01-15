package main

import "fmt"

type GoodImplementation struct{}

func (u *GoodImplementation) Execute() {
	amount := 5000.0
	payment := &CreditCardPayment{}
	err := payment.Process(amount)
	if err != nil {
		fmt.Printf("Payment Failed due to: %s", err.Error())
	}
}

type PaymentProcessor interface {
	Process(amount float64) error
}

type CreditCardPayment struct{}

func (c *CreditCardPayment) Process(amount float64) error {
	fmt.Println("Processing credit card payment:", amount)
	return nil
}

type PayPalPayment struct{}

func (p *PayPalPayment) Process(amount float64) error {
	fmt.Println("Processing PayPal payment:", amount)
	return nil
}

/*
 To add support for new Payment method simply add another struct like UpiPaymentProcessor and implement Process() method
 THis doesn't involve modifying existing methods but requires to Extend our solution.
*/
