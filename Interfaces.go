package main

import "fmt"

// Define an interface
type Speaker interface {
    Speak() string
}

// Define a type that implements the interface
type Dog struct {
    Name string
}

// Implement the Speak method for Dog
func (d Dog) Speak() string {
    return "Woof!"
}

func main() {
    var s Speaker
    d := Dog{Name: "Buddy"}
    s = d // Dog implements Speaker, so it can be assigned

    fmt.Println(s.Speak()) // Outputs: Woof!
}
