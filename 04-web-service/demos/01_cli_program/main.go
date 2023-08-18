package main

import (
	// buffer input output, allowing us to read lines at a time
	"bufio"
	"fmt" // scans -> allows us to read from various sources including stdin
	"os"  // stdin, stdout and stderr
	"strings"
)

func main() {
	fmt.Println("What would you like me to scream?")
	// Stdin is a file-like object that go uses to represent input from keyboard,
	// but it only reads one character at a time
	// bufio.NewReader allows we to read more than one character
	in := bufio.NewReader(os.Stdin)
	// the buffer stops at \n, so we only read one line here.
	// the second variable "_" is the error, but we do not use it now
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s + "!")
}
