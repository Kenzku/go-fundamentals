package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

// define a menu type
// menu is a slices of menuItem
type menu []menuItem

var in = bufio.NewReader(os.Stdin)

// in GO, difference between function and method:
// 1. a method is like in object-oriented programming, it associates to a type
// e.g. a struct, here (m menu) is the type the method "print" associates to
// in other words, the method "print" can only be executed by "menu" type variables.
// (just like an instance.method in Java)
// 2. a function does not have this (m menu) association before the function name.
// 3. (m menu) is called receiver, m is a value receiver
// 4. this is basically a way to add method to type
func (m menu) print() {
	for _, item := range m {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, cost := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, cost)
		}
	}
}

// (m *menu) where m is a pointer receiver, a reference.
// in other words, m inside the function is a reference.
// if you want to update m's value, you need to use *m =

func (m *menu) add() error {
	fmt.Println("Please enter the name of the new item")
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name)

	for _, item := range data {
		if item.name == name {
			return errors.New("menu item already exists")
		}
	}
	// m is a reference, so you need to de-reference it to get or update the value
	*m = append(*m, menuItem{name: name, prices: make(map[string]float64)})
	return nil
}

func Print() {
	// data is a menu type, and print is also a menu type
	data.print()
}
func Add() error {
	return data.add()
}
