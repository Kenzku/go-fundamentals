package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 1. in the interface, you define methods
// 2. in GO, you do not implement the interface,
// but just create a method (not a function) with the same signature as in the interface

type printer interface {
	Print() string
}

type user struct {
	username string
	id       int
}

// Print has the same method signature, as the printer interface
// so GO automatically treats this method as an implementation of the interface.
func (u user) Print() string {
	return fmt.Sprintf("%v [%v]\n", u.username, u.id)
}

type menuItem struct {
	name   string
	prices map[string]float64
}

// implement another interface: just keep the same method signature.

func (mi menuItem) Print() string {
	var b bytes.Buffer
	b.WriteString(mi.name + "\n")
	b.WriteString(strings.Repeat("-", 10) + "\n")
	for size, cost := range mi.prices {
		fmt.Fprintf(&b, "\t%10s%10.2f\n", size, cost)
	}

	return b.String()
}

func main() {
	// this is how to use interface to call its method
	// 1. you define a var using the interface
	var p printer
	// 2. you can then assign the concrete user to the interface.
	// note: the user is the type bound to the implemented method of the interface
	// note: the interface cannot see members of user.
	p = user{username: "adent", id: 42}
	fmt.Println(p.Print())

	// menuItem is another type bound to the implemented method of the interface
	p = menuItem{name: "Coffee",
		prices: map[string]float64{"small": 1.65,
			"medium": 1.80,
			"large":  1.95,
		},
	}

	// 3. you can just call by using the interface.method
	fmt.Println(p.Print())

	// Type Assertion: interface_var.(concrete_type)
	// ok -> false, if the type is not the concrete_type
	u, ok := p.(user)
	fmt.Println(u, ok) // ok is false here, without "ok" (it does not matter if true or false), you get panic

	mi, ok := p.(menuItem)
	fmt.Println(mi, ok)

	// Type Assertion: from interface back to the concrete type.
	// this is called type switches
	switch v := p.(type) {
	// case type:
	case user:
		fmt.Println("Found a user!", v)
	case menuItem:
		// this case will be run, because the last assignment of p is from menuItem
		fmt.Println("Found a menuItem!", v)
	default:
		fmt.Println("I'm not sure what this is...")
	}

	// the main problem with interface is:
	// after you convert a type into an interface, it lost the type, unless you assert it back explicitly
	// in other words, GO runtime does not know what was the original concrete type (of course we know by reading the code)
	// e.g. once you assign a user to a printer p, GO runtime would only know p is a printer, it does not know p is from a user

	// in Generic Function, you get the original type back after
}
