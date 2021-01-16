package bank

import ()

var (
	sema = make(chan struct{})
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{}
	balance += amount
	<- sema
}

func Balances() int {
	sema <- struct{}{}
	b := balance
	<- sema
	return b
}