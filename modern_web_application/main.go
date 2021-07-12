package main

import (
	"fmt"
	"net/http"
)

const (
	portNumber = 8080
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/operation", operationHandler)

	fmt.Printf("starting apllication on port %d\n", portNumber)

	// $ curl localhost:portNumber
	address := fmt.Sprintf(":%d", portNumber)
	_ = http.ListenAndServe(address, nil)
}
