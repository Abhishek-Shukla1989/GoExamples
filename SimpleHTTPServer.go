
package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

type todoN struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const baseurlN = "https://jsonplaceholder.typicode.com/"

var form = `
<h1>Todo #{{.Id}}</h1>
<div>{{printf "User %d" .UserId}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>`

func handler(w http.ResponseWriter, r *http.Request) {

	var item todoN
	resp, err := http.Get(baseurlN + r.URL.Path[1:])
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &item)
	if err != nil {
		http.Error(w, "Failed to unmarshal JSON", http.StatusInternalServerError)
		return
	}

	tmpl := template.New("mine")
	tmpl, err = tmpl.Parse(form)
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, item)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
