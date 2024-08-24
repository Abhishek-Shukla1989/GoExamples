package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

var raw = `
	<!DOCTYPE html>
	<html>
	<body>
	<h1>Hello, World!</h1>
	<image src="https://www.w3schools.com/images/w3schools_green.jpg" alt = "hello"/>
	<p>This is a paragraph.</p>
	<h3>Program to count images and text.</h3>
	 <script>
	 document.getElementById("demo").innerHTML = "Hello JavaScript!";
	 </script>
	</body>
	</html>
	`

func visit(c *html.Node, words, images *int) {

	if c.Type == html.TextNode {
		*words += len(strings.Fields(c.Data))
	} else if c.Type == html.ElementNode && c.Data == "img" {
		*images++
	} else if c.Type == html.ElementNode && c.Data == "script" {
		return // this will ignore all javascript
	}
	for n := c.FirstChild; n != nil; n = n.NextSibling {
		visit(n, words, images)
	}
}
func countWordsAndImages(doc *html.Node) (int, int) {
	var countWords, countImages int
	visit(doc, &countWords, &countImages)
	return countWords, countImages
}

func main() {

	doc, err := html.Parse(bytes.NewReader([]byte(raw)))
	if err != nil {
		fmt.Printf("Error parsing html %v", err)
	}

	word, pics := countWordsAndImages(doc)
	fmt.Printf("Words are %d and images are %d", word, pics)

}
