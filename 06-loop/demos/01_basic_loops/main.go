package main

import (
	"fmt"
)

// all loops in GO is for loop
func main() {
	i := 99

	// infinite loop - just use "for"
	for {
		fmt.Println(i)
		i += 1
		break
	}

	i = 5
	// for condition { }
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// counter-based loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
