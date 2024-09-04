package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	rwMu    sync.RWMutex
	wg      sync.WaitGroup
)

func reader(id int) {
	defer wg.Done()
	rwMu.RLock() // Acquire read lock
	fmt.Printf("Reader %d reading counter: %d\n", id, counter)
	time.Sleep(1 * time.Second) // Simulate read delay
	rwMu.RUnlock()              // Release read lock
}

func writer(id int) {
	defer wg.Done()
	rwMu.Lock() // Acquire write lock
	counter++
	fmt.Printf("Writer %d incremented counter to: %d\n", id, counter)
	time.Sleep(1 * time.Second) // Simulate write delay
	rwMu.Unlock()               // Release write lock
}

func main() {
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go reader(i)
	}
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go writer(i)
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
