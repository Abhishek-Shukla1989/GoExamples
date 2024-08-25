package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getOne(i int) []byte {

	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can'd read %s", err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Skipping %d %d", i, resp.StatusCode)

		return nil
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid body %s", err)

	}

	return body

}
func main() {

	var (
		output io.WriteCloser = os.Stdout
		err    error
		cnt    int
		fails  int
		data   []byte
	)

	if len(os.Args) > 1 {

		output, err = os.Create(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		defer output.Close()
	}

	fmt.Fprint(output, "[")
	defer fmt.Fprint(output, "]")

	for i := 1; i < 300; i++ {

		fmt.Fprintf(os.Stderr, "Loop running  : %d\n", i)

		if data = getOne(i); data == nil {
			fails++
			continue
		}

		if cnt > 0 {
			fmt.Fprint(output, ",")
		}

		_, err = io.Copy(output, bytes.NewBuffer(data))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Stopped : %s\n", err)
			os.Exit(-1)
		}
		cnt++
		fails = 0
	}

	fmt.Fprintf(os.Stderr, "Completed operation, Please check the file")
}
