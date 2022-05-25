package main

import (
	"bookings/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the main applicaiton function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
