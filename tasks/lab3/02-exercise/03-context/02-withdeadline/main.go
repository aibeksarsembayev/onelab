package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func main() {

	// TODO: set deadline for goroutine to return computational result.

	compute := func(ctx context.Context) <-chan data {
		ch := make(chan data)
		go func() {
			defer close(ch)
			for {
				select {
				case <-ctx.Done():
					return
				default:
					// Simulate work.
					time.Sleep(50 * time.Millisecond)
					// Report result.
					ch <- data{"123"}
				}
			}
		}()
		return ch
	}

	// Wait for the work to finish. If it takes too long move on.
	deadline := time.Now().Add(200 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	ch := compute(ctx)
	// d := <-ch
	for d := range ch {
		// fmt.Println(d)
		fmt.Printf("work complete: %s\n", d)
	}

}
