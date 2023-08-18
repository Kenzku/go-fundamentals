package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please select an option")
	fmt.Println("1) Print menu")
	in := bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice) // we don't know what to do with this yet!

	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
		{name: "Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25, "triple": 2.55}},
	}

	// syntax 1: for key, value := range collection_var { }
	// syntax 2: for key := range collection_var { } // ignore value
	// syntax 3: for _, value := range collection_var { } // ignore key
	// collections: array, slice and map
	// range is the keyword to loop through collections
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, cost := range item.prices {
			// \t - tab character
			// %10s - format the size
			// %s is used to print a string.
			// 10 specifies a width of 10 characters for the string.
			// This means the string will be right-aligned within a field of 10 characters.
			// If the string is shorter than 10 characters, it will be padded with spaces on the left.
			//
			// %10.2f: This is a format specifier for a floating-point number -> the cost
			// %f is used to print a floating-point number.
			// .2 2 decimal places
			fmt.Printf("\t%10s%10.2f\n", size, cost)
		}
	}

}
