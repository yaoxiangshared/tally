package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//demo1()
	//demo2()
	demo3()
	time.Sleep(2 * time.Second)
}

func demo1() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
}

func demo2() {
	//var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		//wg.Add(1)
		go func() {
			//defer wg.Done()
			fmt.Println(salutation)
		}()
		//wg.Wait()
	}
}

func demo3() {
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		go func(salutation string) {
			fmt.Println(salutation)
		}(salutation)
	}
}
