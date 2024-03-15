package internal

import "fmt"

// receiver ======================

type Profile struct {
	name    string
	balance float64
}

func (p Profile) Receive(amount float64) {
	p.balance += amount
	fmt.Printf("%s current balance is %f \n", p.name, p.balance)
}

// sender =========================

func Deposit(iterator TransferIterator, amount float64) {
	for iterator.HasNext() {
		iterator.Next().Receive(amount)
	}
}

// implement iterator ==============

type TransferInterface interface {
	Receive(amount float64)
}

type TransferIterator interface {
	Next() TransferInterface
	HasNext() bool
}
