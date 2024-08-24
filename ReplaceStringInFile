package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "not enough arguments")
		os.Exit(-1)
	}
	var finalS string
	new, old := os.Args[2], os.Args[1]

	// this code is to read complete file and replace the string, another way to read files using bufio package

	byte_data, err := os.ReadFile("./cmd/nums.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read file")
		os.Exit(-1)
	}
	new_str := strings.ReplaceAll(string(byte_data), old, new)
	os.WriteFile("./cmd/nums.txt", []byte(new_str), 0666)
	fmt.Println(finalS)
}
