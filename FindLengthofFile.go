package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	for _, fname := range os.Args[1:] {
		//fmt.Println("Reading file:", fname)

		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file:", err)
			continue
		}

		// first count all lines, text and characters
		var nlines, nchars, nwords int

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {

			s := scanner.Text()
			nchars += len(s)
			nlines++
			nwords += len(strings.Fields(s))

		}
		fmt.Printf("File is %s, %d lines, %d words, %d characters\n", fname, nlines, nwords, nchars)

		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprintln(os.Stderr, "Error copying file content to stdout:", err)
			continue
		}

		// fmt.Println("File size is", len(data), "bytes")
		// Ensure file is closed after reading
		defer file.Close()

		// Explicitly close the file (note: defer also ensures closure)
		if err := file.Close(); err != nil {
			fmt.Fprintln(os.Stderr, "Error closing file:", err)
			continue
		}
	}
}
