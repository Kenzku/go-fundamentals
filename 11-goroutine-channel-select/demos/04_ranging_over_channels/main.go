package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			// Itoa short for Integer to ASCII (but it is not limited to ASCII, but also unicode)
			ch <- "message " + strconv.Itoa(i)
		}
		// we need to use close(ch) to terminate the for loop below
		close(ch)
	}()

	// every iteration of the loop will pull one message from the channel
	// range ch is unbounded, so we need to use close(ch) to terminate the for loop
	for msg := range ch {
		fmt.Println(msg)
	}

}
