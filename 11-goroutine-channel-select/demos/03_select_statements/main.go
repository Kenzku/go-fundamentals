package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		// the sleep, make the "select" run case 2
		// time.Sleep(5 * time.Millisecond)
		ch1 <- "message to channel 1"
	}()

	go func() {
		ch2 <- "message to channel 2"
	}()

	// it only pauses the main()
	// it gives time for both goroutines finish, and it demos the "select" randomly chooses a case to run
	time.Sleep(10 * time.Millisecond)

	// "select" statement is similar to "switch"
	// 1. but select is used for channels
	// 2. there is no var after "select"
	select {
	// if there are more than one "cases" can be run, than they are chosen randomly to run
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	// 1. optional, using "default" creates a non-blocking select
	// 2. without "default", the "select" is blocking, until one of the "case" can run
	default:
		fmt.Println("no messages available")
	}

}
