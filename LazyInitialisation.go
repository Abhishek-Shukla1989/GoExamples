package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var config string

func loadConfig() {
	config = "Loaded Configuration"
	fmt.Println("Configuration loaded")
}

func getConfig() string {
	once.Do(loadConfig) // Ensure loadConfig is called only once
	return config
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: %s\n", id, getConfig())
		}(i)
	}
	wg.Wait()
}
