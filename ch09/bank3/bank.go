package bank

import "sync"

var (
	mu sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balances() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}