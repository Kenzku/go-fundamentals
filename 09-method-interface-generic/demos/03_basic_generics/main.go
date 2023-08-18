package main

import (
	"fmt"
)

func main() {
	// a slice of float64
	testScores := []float64{
		87.3,
		105,
		63.5,
		27,
	}
	// create an independent copy (none-referenced)
	c := clone(testScores)

	// &testScores[0] memory address of testScores[0]
	// &c[0] memory address of c[0]
	fmt.Println(c, &testScores[0] == &c[0])
}

// [T any] indicates T is the generic type,
// [T any, S any] for multi-generic types per function
// The keyword "any" is a type constraint that indicates T can be any type
// the letter T does not matter, in other words, you can use F, V etc
func clone[T any](s []T) []T {
	// make is used to init slice / map / chan
	result := make([]T, len(s))
	for i, v := range s {
		result[i] = v
	}

	return result
}
