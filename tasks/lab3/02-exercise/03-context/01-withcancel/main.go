package main

import (
	"context"
	"fmt"
)

func main() {

	// TODO: generator -  generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the goroutine once
	// they consume 5th integer value
	// so that internal goroutine
	// started by gen is not leaked.
	generator := func(ctx context.Context) <-chan int {
		_, cancel := context.WithCancel(ctx)
		defer cancel()
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < 10; i++ {
				select {
				case <-ctx.Done():
					return
				default:
					ch <- i
				}
			}
		}()

		return ch
	}

	// Create a context that is cancellable.
	ctx, cancel := context.WithCancel(context.Background())

	ch := generator(ctx)

	for n := range ch {
		if n == 5 {
			cancel()
		} else {
			fmt.Println(n)
		}
	}
}
