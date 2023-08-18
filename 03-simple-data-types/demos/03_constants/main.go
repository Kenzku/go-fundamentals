package main

import "fmt"

func main() {
	const a = 42
	const b float32 = 3

	var i int = a
	fmt.Println(i)
	var f32 float32 = a
	fmt.Println(f32)
	f32 = b
	fmt.Println(f32)

	// var i2 int = b
	// fmt.Println(i2)

	const c = iota
	fmt.Println(c)

	// create a group of constants, saving typing "const"
	const (
		d = 2 * 5
		// if there is no value provided in a const group,
		// it simply follows the values above
		e
		// iota is 0-index int within the constant group, so "f" is 2
		// no related to the usage, so when we assign `c = iota`, it does not impact on it.
		f = iota
		g
		// constant expression: assigning the result of something
		h = 10 * iota
	)

	fmt.Println(d, e, f, g, h)

}
