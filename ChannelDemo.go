package main

import (
	"fmt"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
		resp.Body.Close()
	} else {

		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func main() {

	results := make(chan result)

	list := []string{
		"https://amazon.com",
		"https://apple.com",
		"https://google.com",
		"https://wsm.com",
		//"http://localhost.com",
	}

	for _, url := range list {
		go get(url, results)
	}

	for range list {
		r := <-results

		if r.err != nil {

			fmt.Printf("Erros is %s and url is %v\n", r.err, r.url)

		} else {
			fmt.Printf("Data received in %v and url is %s\n", r.latency, r.url)

		}
	}
}
