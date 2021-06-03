// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Page 19.

// Example:
// go run main.go & (use & only for Mac OS X or Linux to start the server in background)
// go build 09fetch/main.go (the previous fetch program example)
// ./main http://localhost:8000/ (the fetch compiled program)
// Output: URL.Path = "//"
// ./main http://localhost:8000/help
// Output: URL.Path = "/help"

// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
