package main

import (
	"fmt"
	"sync" // for WaitGroup
)

// 1. concurrency is not the same as parallelism
// concurrency means can do more than one thing at a time (imagine a thread have tiny windows allow to run other tasks and back to the previous tasks)
// parallelism means can run more than one tasks at the same time.
// 2. concurrency in go: worker -> channel -> worker -> channel
// where -> means sending message / data to the next worker via a channel
// Fan-in model: multi-workers -> worker
// Fan-out model: worker -> multi-workers
// Fan-in/out is used to balance the messages throughput
// 3. the "worker" is the "goroutine" in terminology in GO
// a goroutine is a function
// a channel is a type
// 4. a WaitGroup is a counter (you can increase or decrease it), having special behaviour when the value is 0 - like CountDownLatch in Java
func main() {
	// declare a WaitGroup variable
	var wg sync.WaitGroup

	// you need to call this Add() before you call the goroutine is invoked
	wg.Add(1) // Add(delta): add the counter by the delta, here the delta is 1, so wg = wg + 1

	// the keyword "go" turns a func into a "goroutine"
	go func() {
		fmt.Println("This happens asynchronously!")
		wg.Done() // Done(): decrease the counter by 1, i.e. wg = wg - 1
	}()

	fmt.Println("This is synchronous")

	// 1. wait is blocking, until the WaitGroup value is 0
	// if we do not wait, the main() will ends, before the async task (the goroutine) finishes
	// just like future.join()
	// 2. Wait() is called after the goroutine is invoked
	wg.Wait()

}
