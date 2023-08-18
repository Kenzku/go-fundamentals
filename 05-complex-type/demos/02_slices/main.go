package main

import (
	"fmt"
	// set the module GOPATH
	// enable the go module integration
	// run: go get golang.org/x/exp/slices is used to import dependencies into the file
	"golang.org/x/exp/slices"
)

// A slice is a data type like an array, with no fixed size
// A slice does not store data, and always refers to data stored by some array
// the first index (from 0) of the slice points to an element of an array
// like a pointer of the array
// so if you change the value in the array, it changes in the slides,
// and if you change the vavlue in the slices, it also changes the values in the array
func main() {
	// define a slice "s", because we do not provide the size in []
	var s []string
	fmt.Println(s) // print out [] (nil)

	s = []string{"Coffee", "Espresso", "Cappuccino"}
	fmt.Println(s)

	fmt.Println(s[1])
	s[1] = "Chai Tea"
	fmt.Println(s)

	s = append(s, "Hot Chocolate", "Hot Tea")
	fmt.Println("append", s)

	// remove indices from 1 (included) to 2 (excluded)
	s = slices.Delete(s, 1, 2)
	fmt.Println("delete", s)

	s2 := s

	s[2] = "Chai Latte"

	fmt.Println(s, s2)
}
