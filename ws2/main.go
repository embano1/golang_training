// A simple webserver with primitive use of channels to demonstrate goroutines
package main

// Import some packages we need
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// These variables can´t be modified (immutable)
const (
	version = "1.0"
	port    = ":8000"
)

// Entry point for our program (we don´t use init() in this example)
func main() {

	// Used to eventually terminate the program
	finish := make(chan bool)

	//create a notification channel to shutdown
	sigChan := make(chan os.Signal, 1)

	// Define the terminate commands we´re interested in
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Log info
	log.Println("Starting web server on port", port)

	// Register URL path with a http handler function
	http.HandleFunc("/", defaulthandler)
	http.HandleFunc("/v1/info", infohandler)

	// ---!!!--- Run in goroutine and panic if ListenAndServe throws an error
	go func() {
		log.Fatal(http.ListenAndServe(port, nil))
	}()

	// ---!!!--- Run in goroutine and wait for CTRL+C or SIGTERM
	go func() {
		s := <-sigChan
		log.Println("Got signal", s)
		log.Println("Shutting down...")
		finish <- true
	}()

	// Wait until someone sends to finish (see above) and thus terminate the program
	<-finish

}

// A function for our webserver
func defaulthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello surfer!")
}

// Another function for our webserver
func infohandler(w http.ResponseWriter, r *http.Request) {

	// Get the hostname using a call to os package
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hostname: %s \n", hostname)
	fmt.Fprintf(w, "Version: %s", version)

}
