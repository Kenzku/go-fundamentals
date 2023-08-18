package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
func funcName (paras) (return-var return-type) {
	func body
}
paras:
1. (para1 type1, para2 type2)
2. (para1, para2 type) e.g. name1, name2 string -> name1 and name2 both string
3. (para ...type) e.g. names ...string -> names is now a slice (has to be the last parameter)
4. (para1 *type) para is a pointer / reference, to use/update para's value: *para
   -> if para is not a reference, it only copies the value to the parameter,
      so change the parameter value does not change it outside the function

you only use the reference if you want to share memory.

return:
1. you can omit () if you have only 1 return type
2. you can have multiple return values: (return-type1, return type2)
   -> return 10, true
3. named return value: e.g. (result int, ok bool), then you can have a naked return
   -> result = 10
      ok = true
	  return
   in naked return, if the return-vars unset, they will be default values.
*/

// menuItem and menu are outside the main function like javascript
type menuItem struct {
	name   string
	prices map[string]float64
}

var menu = []menuItem{
	{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
	{name: "Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25, "triple": 2.55}},
}

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
			printMenu()
		case "2":
			addItem()
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}
}

func printMenu() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, cost := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, cost)
		}
	}
}
func addItem() {
	fmt.Println("Please enter the name of the new item")
	name, _ := in.ReadString('\n')
	menu = append(menu, menuItem{name: strings.TrimSpace(name), prices: make(map[string]float64)})
}
