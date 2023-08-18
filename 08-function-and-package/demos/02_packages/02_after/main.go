package main

import (
	"bufio"
	// "module/package"
	"demo/menu"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func main() {

loop:
	for {
		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			// menu is the name of the package
			menu.Print( /* this is how to use comment*/ )
		case "2":
			menu.Add()
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}
}
