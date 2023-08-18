package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// A struct is the ONLY heterogeneous aggregate type in GO
// while array and slice are homogeneous types, meaning they have only one type per instance.
// Aggregate Type: an aggregate type is a data type that groups multiple values or variables under a single identifier
// Heterogeneous: the elements or components within the structure can be of different data types
// the struct itself is value type
// so when you assign a struct to another, it just copies the value, like array
func main() {
	fmt.Println("Please select an option")
	fmt.Println("1) Print menu")
	in := bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice) // we don't know what to do with this yet!

	// declare a custom type, using struct
	// now you have a new type "menuItem" to represent the struct,
	// in other words, you simplify the code later when using it
	// (not recommended) if you use "var" instead of "type", but then you create a variable with the "struct" type, and the variable name would be "menuItem"
	type menuItem struct {
		name string
		// map[key]value
		prices map[string]float64
	}

	// menuItem is now a type
	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
		{name: "Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25, "triple": 2.55}},
	}

	fmt.Println(menu)

	// use the "dot" syntax to update struct fields, sometimes called "selector"
	// here, menu is an array type
	menu[0].name = "Coffee+"
	fmt.Println("------------")
	fmt.Println(menu)

	fmt.Println("------------")
	// Declaring and initializing an anonymous struct
	// here s is NOT a type
	var s struct {
		Name    string
		Age     int
		Address string
	}

	// Assigning values to the struct fields
	s.Name = "Alice"
	s.Age = 30
	s.Address = "123 Go St."

	// Accessing and printing the struct fields
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
	fmt.Println("Address:", s.Address)
}
