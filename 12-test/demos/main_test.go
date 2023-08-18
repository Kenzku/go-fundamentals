package main

import "testing"

// 1. unit test: individual unit of the code is working, can be a function, some functions etc
// 2. component test: e.g. packages are working correctly
// 3. integration test: the entire application is working, e.g. build the app as a war, running it in a container
// 4. end-to-end test: the entire application in a system is working, e.g. in a network, work with other applications, services
//
//	e.g. TA UI works with TA Server and DB running is a dockers, stagings etc.

//  1. the file name, _test.go is significant
//  2. the function Name has to be TestXxx (X is upper case) for a test, FuzzXxx, etc
//  3. "t" is an object given by the TestRunner,
//     and we will use the t in the test, and communicate the t back to the TestRunner
//     this means we share data, that is why we pass in reference, *
//  4. run in CLI: go test .
//     go test ./... (3 dots) means to run all test in the current folders and all sub folders
func TestAdd(t *testing.T) {
	// arrangement
	l, r := 1, 2
	expect := 3

	// action
	got := Add(l, r)

	// assertion
	if expect != got {
		// report the error
		// here we use the t to report back to the TestRunner
		t.Errorf("Failed to Add %v and %v. Got %v, expected %v\n",
			l, r, got, expect)
	}
}
