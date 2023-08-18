package main

import (
	"errors"
	"fmt"
)

// avoid panic if possible
// Error is treated as a value, so it is easier to recover e.g. if error, provide a default value
// when Panic, it changes the control flows (e.g. the panic method stops executing), and you use "defer" and recover()
// Panic, indicates, the programme is not stable
// Error just implies the things did not go as plan
func main() {

	// method 1

	//result, err := divide1(10, 5)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("result: ", result)

	// method 2

	//defer func() {
	//	if msg := recover(); msg != nil {
	//		fmt.Printf("something went wrong: %v\n", msg)
	//	}
	//}()
	// is handled by the defer function
	//divide2(10, 0)

	// method 3
	result, err := divide3(10, 0)
	fmt.Println(result, err)

}

// how to handle exception?
// method 1: return int and error, and if error is not nil, it means there is something wrong

func divide1(l, r int) (int, error) {
	if r == 0 {
		return 0, errors.New("invalid divisor: must not be zero")
	}
	return l / r, nil

}

// method 2: rely on GO runtime panic, and you handle that in your parent method
func divide2(l, r int) int {
	return l / r
}

// method 3: use error and panic (defer and recover)

func divide3(l, r int) (result int, err error) {
	defer func() {
		if msg := recover(); msg != nil {
			result = -1
			err = fmt.Errorf("%v", msg)
			// naked return - you do not need the return
		}
	}()
	return l / r, nil
}
