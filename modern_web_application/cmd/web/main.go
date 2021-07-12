package main

import (
	"fmt"
	"net/http"

	"github.com/mohammadne/go_samples/modern_web_application/pkg/handlers"
)

const (
	portNumber = 8080
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("starting apllication on port %d\n", portNumber)

	// $ curl localhost:portNumber
	address := fmt.Sprintf(":%d", portNumber)
	_ = http.ListenAndServe(address, nil)
}
