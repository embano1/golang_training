// A simple webserver
package main

// Import some packages we need
import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// These variables can´t be modified (immutable)
const (
	version = "1.0"
	port    = ":8000"
)

// Entry point for our program (we don´t use init() in this example)
func main() {

	// Log info
	log.Println("Starting web server on port", port)

	// Register URL path with a http handler function
	http.HandleFunc("/", defaulthandler)
	http.HandleFunc("/v1/info", infohandler)

	// Panic if ListenAndServe throws an error
	log.Fatal(http.ListenAndServe(port, nil))

}

// A function for our webserver
func defaulthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello surfer!")
}

// Another function for out webserver
func infohandler(w http.ResponseWriter, r *http.Request) {

	// Get the hostname using a call to os package
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hostname: %s \n", hostname)
	fmt.Fprintf(w, "Version: %s", version)

}
