package main

import (
	"errors"
	"fmt"
	"time"
)

type BadImplementation struct{}

func (u *BadImplementation) Execute() {

	b := BankAccountA{
		Balance: 50000,
	}

	b.Deposit(1000)
	b.SaveToDB()
	b.Withdraw(20000)
	b.SaveToDB()
	b.SendEmailNotification()

}

type BankAccountA struct {
	Balance float32
}

/*
❌ Why this violates SRP
 -> Business logic
 -> Persistence
 -> Notification
All mixed together → multiple reasons to change
*/

func (b *BankAccountA) Deposit(amount float32) {
	fmt.Println("Depositing Funds")
	b.Balance += amount
}

func (b *BankAccountA) Withdraw(amount float32) error {
	fmt.Println("Withdrawing Funds")
	if amount > b.Balance {
		return errors.New("insufficient funds")
	}
	b.Balance -= amount
	return nil
}

func (b *BankAccountA) SendEmailNotification() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Sending email notification")
}

func (b *BankAccountA) SaveToDB() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Saving account to database")
}
