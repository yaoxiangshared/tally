package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {

	args := os.Args
	a, _ := strconv.ParseInt(args[1], 10, 32)
	switch a {
	case 0:
		{
			demo0()
			break
		}
	case 1:
		{
			demo1()
			break
		}
	case 2:
		{
			demo2()
			break
		}
	case 3:
		{
			demo3()
			break
		}
	}
	time.Sleep(2 * time.Second)
}

func demo0() {
	salutation := "hello"
	go func() {
		salutation = "welcome"
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(salutation)
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
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		go func() {
			fmt.Println(salutation)
		}()
	}
}

func demo3() {
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		go func(salutation string) {
			fmt.Println(salutation)
		}(salutation)
	}
}
