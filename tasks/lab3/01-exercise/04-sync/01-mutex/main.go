package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(4)

	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex
	deposit := func(amount int) {
		balance += amount
	}

	withdrawal := func(amount int) {
		balance -= amount
	}

	// make 100 deposits of $1
	// and 100 withdrawal of $1 concurrently.
	// run the program and check result.

	// TODO: fix the issue for consistent output.

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			deposit(1)
			mu.Unlock()
		}()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			withdrawal(1)
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(balance)
}
