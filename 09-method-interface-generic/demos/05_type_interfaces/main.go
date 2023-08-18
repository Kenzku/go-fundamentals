package main

import (
	"fmt"
)

// Addable this is a type constraint, (even though they look similar) not a behaviour interface (with methods)
// int | float64 | string: This is a type list.
// It specifies that any type T that satisfies the Addable constraint
// must be one of int, float64, or string.
type Addable interface {
	// you can add "string" like in Javascript too
	int | float64 | string
}

func main() {
	a1 := []int{1, 2, 3}
	a2 := []float64{3.14, 6.02}
	a3 := []string{"foo", "bar", "baz"}

	s1 := add(a1) // s1 is an int
	s2 := add(a2) // s2 is a float64
	s3 := add(a3) // s3 is a string

	fmt.Printf("Sum of %v: %v\n", a1, s1)
	fmt.Printf("Sum of %v: %v\n", a2, s2)
	fmt.Printf("Sum of %v: %v\n", a3, s3)
}

// Instead of allowing any type,
// you can restrict the kind of types a type parameter can accept by setting a constraint.
// This is done using Type Interface. (this case, Addable)
func add[V Addable](s []V) V {
	var result V
	for _, v := range s {
		result += v
	}
	return result
}
