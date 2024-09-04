package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex // Create a mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()   // Lock the mutex before accessing the counter
			counter++   // Safely increment the counter
			mu.Unlock() // Unlock the mutex after the operation
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
