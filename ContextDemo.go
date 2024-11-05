package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type resultN1 struct{
url string
err error
latency time.Duration

}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(),250*time.Millisecond)
    defer cancel()
	resultT := make(chan resultN1)

	list := []string{
		"http://www.google.com",
		"http://www.apple.com",
		"http://www.amazon.com",
	}

	for _, n:= range list{
		go getData(ctx,n, resultT)
	}

	for range list {
		select{
		case dt:= <-resultT:
		 if dt.err != nil{

			fmt.Println("Error getting data from", dt.url)
		 }else{
			fmt.Println("Getting data from" ,dt.url," in",dt.latency,"seconds" )
		 }
		case <-ctx.Done():
			fmt.Println("request timed out")
		
		}
	}
}

func getData(ctx context.Context,url string, ch chan<- resultN1){
	
	start:= time.Now()

	req, _:= http.NewRequestWithContext(ctx, "GET",url,nil)

	if resp, err := http.DefaultClient.Do(req); err != nil{

		ch <- resultN1{
			url,
			nil,
			0,
		}
		resp.Body.Close()

	}else{
		t:= time.Since(start).Round(time.Millisecond)
		ch<- resultN1{url, nil, t}
		resp.Body.Close()

	}

}


