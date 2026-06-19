package main

import "sync/atomic"

var balance atomic.Int64

func Deposit(amount int64) {
	balance.Add(amount)
}

func Balance() int64 {
	return balance.Load()
}

func main() {

}
