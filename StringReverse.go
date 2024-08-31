package main

import (
	"fmt"
	"slices"
)

func main() {
  
	str := "abhishek"
	newbyte := []byte(str)
	slices.Reverse(newbyte)
	str = string(newbyte)

}
