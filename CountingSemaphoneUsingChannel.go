package main

import (
	"fmt"
	"sync"
)

func do() int {

	var n int64
	var w sync.WaitGroup
	m := make(chan bool, 1) // Created channel and resolve this issue by counting semaphore. It has capacity of
	// of only one which means it limited the number of active go routines. Not it is safe and n will be used by one go routine at a time.
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func() {
			m <- true
			n++ // Since counter++ is not atomic (it involves a read-modify-write cycle),
			//multiple goroutines can interfere with each other, leading to incorrect results.

			<-m
			w.Done()
		}()
	}
	w.Wait()
	return int(n)
}
func main() {
	fmt.Println(do())
}
