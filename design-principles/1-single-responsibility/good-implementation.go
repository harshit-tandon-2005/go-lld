package main

import (
	"errors"
	"fmt"
	"time"
)

type (
	GoodImplementation struct{}

	BankAccountB struct {
		Balance float32
	}

	AccountRepository struct{}

	NotificationService struct{}
)

func (u *GoodImplementation) Execute() {
	b := BankAccountB{
		Balance: 50000,
	}

	b.Deposit(1000)
	b.Withdraw(20000)

	a := AccountRepository{}
	a.Save(&b)

	n := NotificationService{}
	n.Send("Data")
}

/*
🧠 Key Takeaway

One struct → one job
In Go, small structs + small methods are the norm.
*/

func (b *BankAccountB) Deposit(amount float32) {
	fmt.Println("Depositing Funds")
	b.Balance += amount
}

func (b *BankAccountB) Withdraw(amount float32) error {
	fmt.Println("Withdrawing Funds")
	if amount > b.Balance {
		return errors.New("insufficient funds")
	}
	b.Balance -= amount
	return nil
}

func (r *AccountRepository) Save(account *BankAccountB) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Saving account to database")
}

func (n *NotificationService) Send(message string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Sending email notification", message)
}
