package main

import "fmt"
import "sync"

func main()  {
	var data int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		data++
		defer wg.Done()
	}() 
	wg.Wait()
	if data != 0 {
		fmt.Println("hello, world!")
	}
}