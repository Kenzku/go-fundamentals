package main

import "fmt"

func main() {
	var arr [3]string
	fmt.Println(arr)

	arr = [3]string{"Coffee", "Espresso", "Cappuccino"}
	fmt.Println(arr)

	fmt.Println(arr[1])
	arr[1] = "Chai Tea"

	fmt.Println(arr)

	// declare arr2, and assign the array arr to it.
	// you only copy the value, not the pointer
	// so change elements in arr does not impact on arr2 later
	// this is a copy operation
	arr2 := arr

	arr2[2] = "Chai Latte"

	fmt.Println(arr, arr2)
}
