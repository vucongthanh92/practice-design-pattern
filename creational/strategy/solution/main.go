package main

import "fmt"

type Notifier interface {
	Send(msg string)
}

type EmailNotify struct{}

func (EmailNotify) Send(msg string) {
	fmt.Println("Email send: ", msg)
}

type SMSNotify struct{}

func (SMSNotify) Send(msg string) {
	fmt.Println("SMS send: ", msg)
}

type NotificationSvc struct {
	notifier Notifier
}

func (s NotificationSvc) SendNotification(msg string) {
	s.notifier.Send(msg)
}

func main() {
	s := NotificationSvc{
		notifier: SMSNotify{},
	}
	s.SendNotification("Hello World")
}
