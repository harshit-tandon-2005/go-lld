package main

import "fmt"

type BadImplementation struct{}

func (u *BadImplementation) Execute() {

}

func ProcessPayment(method string, amount float64) {
	if method == "credit" {
		fmt.Println("Processing credit card payment")
	} else if method == "paypal" {
		fmt.Println("Processing PayPal payment")
	}
}

// To Add a new payment handle we need to modiy the method ProcessPayment
