package main

import (
	"fmt"
)

// function invoke -> execute statements -> exit function -> execute deferred statements -> return to caller
// defer is used for:
// 1. Resource Cleanup, like database, connections, files
// 2. Simplifying Complex Functions: instead of using multi-returns, you can use defer.
// 3. Ordered Execution: Last-In-First-Out (LIFO) -> just like the example below
// 4. Handling Panics: When combined with Go's panic and recover mechanism, (like in someFunction)
// defer can be used to handle exceptional situations gracefully, allowing you to recover from panics and handle the error
func main() {

	fmt.Println("main 1")
	defer fmt.Println("defer 1")
	fmt.Println("main 2")
	defer fmt.Println("defer 2")

	// import "database/sql"
	// db, _ := sql.Open("driverName", "connection string")
	// defer db.Close()

	// rows, _ := db.Query("some sql query here")
	// defer rows.Close()

	// then, after the main() exit, the defer functions running orders: rows.Close() -> db.Close()

}

func someFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()
	// ... some code that might panic ...
}
