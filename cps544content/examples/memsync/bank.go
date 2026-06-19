package bank

import (
	"sync"
	"sync/atomic"
)

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

var balance2 int32

func Deposit2(amount int32) {
	atomic.AddInt32(&balance2, amount)
}
