package main

import "fmt"

/*
✅ LSP holds because:

Any implementation can replace another

No special casing

No runtime surprises
*/

type GoodImplementation struct{}

func Notify(handler NotificationHandler, to, message string) error {
	return handler.Send(to, message)
}

func (u *GoodImplementation) Execute() {

	Notify(EmailNotification{}, "a@b.com", "Hello")
	Notify(SMSNotification{}, "1234567890", "Hello")
	Notify(WhatsAppNotification{}, "9876543210", "Hello")

}

type NotificationHandler interface {
	Send(to string, message string) error
}

type EmailNotification struct{}

func (e EmailNotification) Send(to, message string) error {
	// send email
	fmt.Println("Email sent to", to)
	return nil
}

type SMSNotification struct{}

func (s SMSNotification) Send(to, message string) error {
	// send SMS
	fmt.Println("SMS sent to", to)
	return nil
}

type WhatsAppNotification struct{}

func (w WhatsAppNotification) Send(to, message string) error {
	// send WhatsApp message
	fmt.Println("WhatsApp message sent to", to)
	return nil
}
