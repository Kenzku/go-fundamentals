package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("this is an error")

	fmt.Println(err)

	// %w is used with fmt.Errorf to wrap an error, like add more details
	// you can use errors.Unwrap to unwrap it later
	err2 := fmt.Errorf("this error wraps the first one: %w", err)
	fmt.Println("wrapped: ", err2)

	fmt.Println("unwrapped: ", errors.Unwrap(err2))
}
