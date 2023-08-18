package main

import (
	"fmt"
	"sync" // for goroutine
)

// a channel is a type i.e. chan
// channel must be created with the built-in "make" function
func main() {
	var wg sync.WaitGroup
	// 1. create an unbuffered channel of string
	// 2. you can only create a channel of one type
	ch := make(chan string)

	wg.Add(1)
	go func() {
		// 1. send a message to the channel
		// 2. this line will block, until someone reads from the channel ( does not need to finish the reads )
		ch <- "the message"
	}()

	go func() {
		// 1. send the message out from the channel
		// 2. this will block, if there is no message in the channel, until someone sends messages to the channel
		fmt.Println(<-ch)
		wg.Done()
	}()

	// preventing the main() exit before the goroutines finish
	wg.Wait()

}
