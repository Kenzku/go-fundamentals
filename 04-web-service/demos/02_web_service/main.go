package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	// HandleFunc allows us to register a function as a back controller
	http.HandleFunc("/", Handler)
	// default to localhost, if you do not provide an IP
	// the handler is the front controller, and use nil to use the default front controller
	http.ListenAndServe(":3000", nil)
}

// Handler back controller
func Handler(w http.ResponseWriter, r *http.Request) {
	// return a pair file and error
	f, _ := os.Open("menu.txt")

	// copy from the read source (f) to a write destination (w)
	io.Copy(w, f)
}
