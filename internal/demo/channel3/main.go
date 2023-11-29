package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	balance float64
	sema    = make(chan int, 1)
)

func Deposit(amount float64) {
	Lock(sema)
	defer Unlock(sema)
	currentBalance := balance
	time.Sleep(1 * time.Second)
	balance = currentBalance + amount
}
func Withdraw(amount float64) {
	Deposit(-amount)
}
func Balance() float64 {
	Lock(sema)
	defer Unlock(sema)
	return balance
}
func Lock(sema chan int) {
	sema <- 1
}

func Unlock(sema chan int) {
	<-sema
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
