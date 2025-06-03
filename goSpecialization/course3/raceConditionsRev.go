package main

import (
	"fmt"
	"sync"
	"time"
)

// A race condition occurs when multiple concurrent processes (in this case, goroutines) access and manipulate shared data simultaneously, and the outcome depends on the relative timing of their execution.
// In this program:

// We have a shared variable counter that both goroutines are trying to increment.
// Each goroutine executes the same loop 1000 times: read the current value, wait a tiny bit, then increment and write back the value.
// If the program worked correctly, the final value of counter should be 2000 (1000 increments from each goroutine).

// How the Race Condition Occurs
// Here's the problematic sequence that can happen:

// Goroutine A reads counter (let's say it's 42)
// Goroutine B also reads counter (still 42)
// Goroutine A calculates tmp + 1 (43) and writes it to counter
// Goroutine B calculates tmp + 1 (also 43) and writes it to counter

// The end result is that counter is 43, even though two increments were performed.

func main() {
	// Shared variable
	counter := 0

	// WaitGroup to wait for both goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// First goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			// Read counter
			tmp := counter
			// Small delay to increase chance of race condition
			time.Sleep(time.Microsecond)
			// Write counter
			counter = tmp + 1
		}
	}()

	// Second goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			// Read counter
			tmp := counter
			// Small delay to increase chance of race condition
			time.Sleep(time.Microsecond)
			// Write counter
			counter = tmp + 1
		}
	}()

	// Wait for both goroutines to finish
	wg.Wait()

	// Print final counter value
	fmt.Println("Expected count:", 20)
	fmt.Println("Actual count:", counter)
}
