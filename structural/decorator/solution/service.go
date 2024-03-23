package main

import "fmt"

type Notifier interface {
	Send(msg string)
}

type EmailNotify struct {
	Receiver string
	Sender   string
}

func (n EmailNotify) Send(msg string) {
	fmt.Printf("*** Send Email:::: Sender: %s :::: Receiver: %s::::: Content: %s:::: \n",
		n.Sender, n.Receiver, msg)
}

type SMSNotify struct {
	PhoneNumber string
}

func (n SMSNotify) Send(msg string) {
	fmt.Printf("*** Send SMS:::: Phone number: %s ::::Content: %s:::: \n",
		n.PhoneNumber, msg)
}

type TelegramNotify struct {
	AccountName string
	GroupName   string
}

func (n TelegramNotify) Send(msg string) {
	fmt.Printf("*** Send Telegram:::: Group: %s :::: Account: %s::::: Content: %s:::: \n",
		n.GroupName, n.AccountName, msg)
}

type NotifyDecorator struct {
	notifier Notifier
	next     *NotifyDecorator
}

func (n NotifyDecorator) Send(msg string) {
	n.notifier.Send(msg)
	if n.next != nil {
		n.next.Send(msg)
	}
}

func (n NotifyDecorator) Decorate(notifier Notifier) NotifyDecorator {
	return NotifyDecorator{
		notifier: notifier,
		next:     &n,
	}
}

func NewNotifyDecorator(notifier Notifier) NotifyDecorator {
	return NotifyDecorator{
		notifier: notifier,
	}
}

type Service struct {
	notifier Notifier
}

func (s Service) SendNotification(msg string) {
	s.notifier.Send(msg)
}

func main() {
	notifier := NewNotifyDecorator(EmailNotify{
		Sender:   "fbnc@gmail.com",
		Receiver: "bbc@gmail.com",
	}).
		Decorate(SMSNotify{
			PhoneNumber: "0981643135",
		}).
		Decorate(TelegramNotify{
			GroupName:   "Dev",
			AccountName: "engineer",
		})

	s := Service{
		notifier: notifier,
	}

	msg := "Hello World !"
	s.SendNotification(msg)
}
