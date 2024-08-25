package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type todo struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const baseurl = "https://jsonplaceholder.typicode.com/todos/"

func main() {

	var item todo
	resp, _ := http.Get(baseurl + "1")
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	err := json.Unmarshal(body, &item)
	if err != nil {
		fmt.Println("Error is", err)
		return
	}
	fmt.Println("Data is", item)

}
