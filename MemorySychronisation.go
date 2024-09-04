package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedVar int
var mu sync.Mutex

func writer() {
	mu.Lock()
	sharedVar = 42 // Write to shared variable
	mu.Unlock()
}

func reader() {
	mu.Lock()
	fmt.Println("Shared Variable:", sharedVar) // Read shared variable
	mu.Unlock()
}

func main() {
	go writer()
	time.Sleep(1 * time.Second)
	go reader()
	time.Sleep(1 * time.Second)
}
