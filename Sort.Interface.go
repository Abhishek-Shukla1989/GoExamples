//Implement sort.Interface to make a type sortable: 
// Implementation of custom interface which is alternative of sort.Slicestring, and pass it in sort.Sort function.



package main

import (
	"fmt"
	"sort"
)

type Organ struct {
	Name   string
	Weight int
}

type Organs []Organ

func (s Organs) Len() int      { return len(s) }
func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByName struct{ Organs }

func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

type ByWeight struct{ Organs }

func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func main() {

	s := []Organ{
		{"brain", 1340}, {"heart", 290},
		{"liver", 1494}, {"pancreas", 131},
		{"spleen", 162},
	}
	sort.Sort(ByWeight{s}) // pancreas first
	fmt.Println(s)
	sort.Sort(ByName{s}) // brain first
	fmt.Println(s)

}



