package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	balance float64
	mu      sync.Mutex
)

func Deposit(amount float64) {
	mu.Lock()
	defer mu.Unlock()
	currentBalance := balance
	time.Sleep(1 * time.Second)
	balance = currentBalance + amount
}
func Withdraw(amount float64) {
	Deposit(-amount)
}
func Balance() float64 {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit(200) // A1
	}()

	// Bob:
	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit(100)
	}()

	wg.Wait()
	fmt.Printf("balance:%f\n", Balance())
}
