package main

import (
	"errors"
	"fmt"
	"log"
)

type Product struct {
	Name  string
	Price float32
}

type Inventory struct {
	products []Product
}

func (i *Inventory) Lookup(name string) (*Product, error) {
	for _, product := range i.products {
		if product.Name == name {
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}

type Account struct {
	Name    string
	balance float32
}

func (a *Account) Deposit(amount float32) {
	a.balance += amount
}

func (a *Account) Withdraw(amount float32) {
	a.balance -= amount
}

func (a *Account) GetBalance() float32 {
	return a.balance
}

type AccountStorage struct {
	accounts []*Account
}

func (as AccountStorage) Lookup(name string) (*Account, error) {
	for _, acc := range as.accounts {
		if acc.Name == name {
			return acc, nil
		}
	}
	return nil, errors.New("account not found")
}

type FacadeSvc struct {
	inventory      Inventory
	accountStorage AccountStorage
}

func NewFacadeSvc() FacadeSvc {
	return FacadeSvc{
		inventory: Inventory{
			products: []Product{
				{Name: "iphone 10", Price: 25},
				{Name: "iphone 11", Price: 28},
				{Name: "iphone 12", Price: 30},
			},
		},
		accountStorage: AccountStorage{
			accounts: []*Account{
				{Name: "thanhvc", balance: 25000},
				{Name: "notail", balance: 32000},
			},
		},
	}
}

func (s *FacadeSvc) BuyProduct(productName, accountName string) error {
	product, err := s.inventory.Lookup(productName)
	if err != nil {
		return err
	}

	account, err := s.accountStorage.Lookup(accountName)
	if err != nil {
		return err
	}

	if account.GetBalance() < product.Price {
		return errors.New("not enough balance in account")
	}

	account.Withdraw(product.Price)

	return nil
}

func (s *FacadeSvc) Deposit(accountName string, amount float32) error {
	account, err := s.accountStorage.Lookup(accountName)
	if err != nil {
		return err
	}

	account.Deposit(amount)

	return nil
}

func (s *FacadeSvc) FetchBalance(accountName string) float32 {
	account, err := s.accountStorage.Lookup(accountName)
	if err != nil {
		return 0
	}

	return account.GetBalance()
}

func main() {
	service := NewFacadeSvc()

	// case 1: buy a product with an account
	productName := "iphone 10"
	accountName := "thanhvc"

	if err := service.BuyProduct(productName, accountName); err != nil {
		log.Fatal(err)
	}

	// check my balance
	fmt.Println("Account balance: ", service.FetchBalance(accountName))

	//case 2: Deposit 100 into account
	if err := service.Deposit(accountName, 100); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Account balance: ", service.FetchBalance(accountName))
}
