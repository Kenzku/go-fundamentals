package main

import "fmt"

// In Go, panic is a built-in function that stops the ordinary flow of control and begins panicking.
// can be called by us, or by GO
// when is panic, the function panics will be terminated,
// and it is a chain all the way up, until there is "defer recover" -> like "catch(Exception e)"
// panic is like "throw new Exception()"
func main() {
	dividend, divisor := 10, 5
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))

	// this panic is called by the go runtime
	dividend, divisor = 10, 0
	// The %v verb gives a general, default representation of the value.
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide2(dividend, divisor))
}

func divide(dividend, divisor int) int {
	// 1. this is the "defer recover"
	// with the "defer recover", the caller of divide() will not panic
	// in other words, the chain stops
	// when the panic occurs, it returns the default value of the type int (i.e. 0)
	// 2. this is an anonymous function
	defer func() {
		// recover() returns what ever is passed into the "panic(arg)"
		if recover() != nil {
			fmt.Println("Panic detected!")
		}
	}() // <---- call the anonymous function immediately
	return dividend / divisor
}

// you can customise return value when panic
func divide2(dividend, divisor int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic detected!")
			result = -1
		}
	}()
	result = dividend / divisor
	// Naked Return: This automatically returns the current value of result
	// (because the return type of int with var result)
	return
}
