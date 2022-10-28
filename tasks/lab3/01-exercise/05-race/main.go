package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//TODO: identify the data race
// fix the issue.

func main() {
	start := time.Now()
	var t *time.Timer
	mu := sync.Mutex{}
	mu.Lock()
	t = time.AfterFunc(randomDuration(), func() {

		fmt.Println(time.Now().Sub(start))
		mu.Lock()
		t.Reset(randomDuration())
		mu.Unlock()
	})
	mu.Unlock()

	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

//----------------------------------------------------
// (main goroutine) -> t <- (time.AfterFunc goroutine)
//----------------------------------------------------
// (working condition)
// main goroutine..
// t = time.AfterFunc()  // returns a timer..

// AfterFunc goroutine
// t.Reset()        // timer reset
//----------------------------------------------------
// (race condition- random duration is very small)
// AfterFunc goroutine
// t.Reset() // t = nil

// main goroutine..
// t = time.AfterFunc()
//----------------------------------------------------
