package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go fun("go function call")
	// goroutine with anonymous function
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("go anonymous function")
		}
	}()
	// goroutine with function value call
	go func(limiter int) {
		for i := 0; i < limiter; i++ {
			fmt.Println("go anonymous function with value")
		}
	}(3)
	// wait for goroutines to end
	time.Sleep(time.Second)
	
	fmt.Println("done..")
}
