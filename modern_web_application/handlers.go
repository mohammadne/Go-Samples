package main

import (
	"net/http"
)

// homeHandler is home handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTmpl(w, "home")
	// fmt.Fprintln(w, "welcome to simple home page of my web server")
}

// operationHandler is about handler
func operationHandler(w http.ResponseWriter, r *http.Request) {
	renderTmpl(w, "operation")

}

// func addValues(x, y int) int {
// 	return x + y
// }

// func devideValues(x, y float32) (float32, error) {
// 	if y == 0 {
// 		return 0, errors.New("can't devide by zero")
// 	}

// 	return x / y, nil
// }
