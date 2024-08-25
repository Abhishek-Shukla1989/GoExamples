package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type xkcd struct {
	Num        int    `json:"num"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "No file given")
		os.Exit(-1)
	}

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "No search text provided")
		os.Exit(0)
	}

	var (
		items []xkcd
		// terms []string
		input io.ReadCloser
		count int
		err   error
	)

	fmt.Fprintf(os.Stderr, "Arguments of reading file %v", os.Args[2:])

	//	Code to read small files in from disk line by line
	input, err = os.Open("comic.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file", err)
		return
	}

	decoder := json.NewDecoder(input) // Use a JSON decoder for the ReadCloser
	err = decoder.Decode(&items)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file", err)
		return
	}

	matchedData := matchTextWithData(os.Args[2:], items, &count)
	//Code to read small files in from disk
	// data, err := os.ReadFile("comic.json")
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "Error reading file", err)
	// 	return
	// }
	// err = json.Unmarshal(data, &items)
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "Error marshalig file", err)
	// 	return
	// }

	fmt.Fprintf(os.Stderr, "found data is %#v and count is %d", matchedData, count)

	fmt.Fprintf(os.Stderr, "found %d comic", count)
}

func matchTextWithData(text []string, data []xkcd, count *int) []xkcd {

	var matched []xkcd

	for _, element := range text {

		for _, comic := range data {
			if strings.Contains(comic.Title, element) {
				matched = append(matched, comic)
				*count++
			}
		}

	}

	return matched
}
