package main

import (
	"fmt"
	"maps"
)

// a map is dynamic size like a slice
// a map is also a reference type like a slice
// keys is not ordered
func main() {
	// [key-type]value-type
	// here we have key as string, and value as int
	var map1 map[string]int
	fmt.Println("map1", map1)
	map1 = map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map1", map1)

	// here, []string is the array type, so we create a map with key as string, and value as array
	var m map[string][]string
	fmt.Println("m", m) // return map[] nil

	println("-------------")

	// use multi-lines init syntax:
	// 1. []string is omitted, we use {} directly
	// 2. we need the last , in GO, so GO knows to continue next line
	m = map[string][]string{
		"coffee": {"Coffee", "Espresso", "Cappuccino"},
		"tea":    {"Hot Tea", "Chai Tea", "Chai Latte"},
	}

	fmt.Println(m)

	fmt.Println(m["coffee"])

	// delete key in map
	delete(m, "tea")
	fmt.Println(m)

	println("-------------")

	// we need the "[]string", this is the array syntax
	m["other"] = []string{"Hot Chocolate"}
	fmt.Println(m)

	fmt.Println(m["tea"]) // return []

	v, ok := m["tea"]

	fmt.Println(v, ok) // v -> default value of the type, ok -> false, if the map does not have such key

	// since map is a reference type, when we do "m2 := m", we actually assign the reference
	// in other words, if we change m2, m will be changed; and if we change m, m2 will also change
	m2 := m
	m2["coffee"] = []string{"Coffee"}
	m["tea"] = []string{"Hot Tea"}

	fmt.Println(m)
	fmt.Println(m2)

	println("-------------")

	// to create an independent clone, use Clone()
	m3 := maps.Clone(m)

	m3["new"] = []string{"unique from m"}
	fmt.Println("m", m)
	fmt.Println("m3", m3)

}
