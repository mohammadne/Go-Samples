package handlers

import (
	"net/http"

	"github.com/mohammadne/go_samples/modern_web_application/pkg/render"
)

// Home is home handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTmpl(w, "home")
	// fmt.Fprintln(w, "welcome to simple home page of my web server")
}

// About is about handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTmpl(w, "about")
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
