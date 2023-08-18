package main

import "fmt"

// pointers are used to share memory, use copy whenever possible
// go is often used in highly concurrent scenarios such as web services
// pointers makes less likely to have data race, because pointers do not share memory
// pointer means address of memory, or just memory
func main() {
	s := "Hello, world!"

	// assign s's address to p
	// so, p is pointing to s's memory address
	p := &s

	// we print out the memory address of s
	fmt.Println(p)

	// *p is to de-reference p,
	// i.e. to de-reference &s, so we get s
	// and s is "Hello, world!"
	fmt.Println(*p)

	// *p is de-reference of &s, i.e. s
	// so we are doing: s = "Hello, Gophers!"
	*p = "Hello, Gophers!"

	fmt.Println("s -> ", s, "*p ->", *p)

	// "new" function creates pointer to anonymous variable
	// "string" is the type of the pointer, and p is the string type in the previous declaration.
	// here we just re-assign a new pointer to p
	p = new(string)

	fmt.Println("p -> ", p, "*p -> ", *p, "is empty, because there is no content in that address")

}
