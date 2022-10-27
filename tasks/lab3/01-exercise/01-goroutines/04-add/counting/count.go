package counting

import (
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateNumbers - random number generation
func GenerateNumbers(max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

// Add - sequential code to add numbers
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

//TODO: complete the concurrent version of add function.

// AddConcurrent - concurrent code to add numbers
func AddConcurrent(numbers []int) int64 {
	var sum int64
	// Utilize all cores on machine
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	// Divide the input into parts
	for _, value := range numbers {
		// Run computation for each part in seperate goroutine.
		wg.Add(1)
		go func(i int) {
			// Add part sum to cummulative sum
			sum += int64(i)
			wg.Done()
		}(value)
		wg.Wait()
	}
	return sum
}
