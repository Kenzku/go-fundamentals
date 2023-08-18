package main

import (
	"fmt"
)

// Resource example of multi-return clean up with defer
type Resource struct {
	name string
}

func (r *Resource) acquire() error {
	// Simulate some process of acquiring a resource
	fmt.Printf("Acquired %s\n", r.name)
	return nil
}

func (r *Resource) release() {
	// Simulate some process of releasing a resource
	fmt.Printf("Released %s\n", r.name)
}

func acquireResource1() (*Resource, error) {
	res := &Resource{name: "Resource1"}
	return res, res.acquire()
}

func acquireResource2() (*Resource, error) {
	res := &Resource{name: "Resource2"}
	return res, res.acquire()
}

func acquireResource3() (*Resource, error) {
	res := &Resource{name: "Resource3"}
	return res, res.acquire()
}

/** without defer
func complexFunction() error {
    resource1, err := acquireResource1()
    if err != nil {
        return err
    }
    resource2, err := acquireResource2()
    if err != nil {
        resource1.release()
        return err
    }
    resource3, err := acquireResource3()
    if err != nil {
        resource2.release()
        resource1.release()
        return err
    }

    // some code that uses resources

    resource3.release()
    resource2.release()
    resource1.release()
    return nil
}
*/

func main() {
	resource1, err := acquireResource1()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resource1.release()

	resource2, err := acquireResource2()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resource2.release()

	resource3, err := acquireResource3()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resource3.release()

	// ... some operations on the resources ...
}
