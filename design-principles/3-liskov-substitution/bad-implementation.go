package main

import (
	"errors"
	"fmt"
	"time"
)

type BadImplementation struct{}

func (u *BadImplementation) Execute() {

}

type BadNotificationHandler interface {
	Send(to string, message string) error
	Schedule(time.Time) error
}

type BadEmailNotification struct{}

func (e BadEmailNotification) Send(to, message string) error {
	// send email
	fmt.Println("Email sent to", to)
	return nil
}
func (e BadEmailNotification) Schedule(t time.Time) error { return nil }

/*
🚨 LSP violation:

# The type technically implements the interface

# But substituting it breaks behavior expectations

Caller must now “know” which implementation it’s dealing with
*/
type BadSMSNotification struct{}

func (s BadSMSNotification) Send(to, message string) error {
	// send SMS
	fmt.Println("SMS sent to", to)
	return nil
}
func (s BadSMSNotification) Schedule(t time.Time) error {
	return errors.New("SMS scheduling not supported")
}

type BadWhatsAppNotification struct{}

func (w BadWhatsAppNotification) Send(to, message string) error {
	// send WhatsApp message
	fmt.Println("WhatsApp message sent to", to)
	return nil
}
func (e BadWhatsAppNotification) Schedule(t time.Time) error { return nil }

/*
Bad LSP (The Violation)
Imagine we create a SilentNotifier. It satisfies the interface, but it has a
"side effect" or a requirement that breaks the caller's expectations.
*/

type SilentNotifier struct{}

func (s SilentNotifier) Send(to, msg string) error {
	if msg == "" {
		// This might be fine, but what if it panics or
		// requires a global variable to be set first?
		return errors.New("cannot send empty")
	}
	// It doesn't actually send anything, it just logs it.
	// While technically valid code, if the caller EXPECTS an
	// external notification to go out, this "fake" behavior
	// can be seen as an LSP violation.
	return nil
}

/*
✅ Fix: smaller interfaces (Go’s strength)

type Sender interface {
	Send(to, message string) error
}

type Scheduler interface {
	Schedule(time.Time) error
}

*/
