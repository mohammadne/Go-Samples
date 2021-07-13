package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mohammadne/go_samples/modern_web_application/pkg/config"
	"github.com/mohammadne/go_samples/modern_web_application/pkg/handlers"
	"github.com/mohammadne/go_samples/modern_web_application/pkg/render"
)

const (
	portNumber = 8080
)

func main() {
	tmplCache, err := render.CreateTmplCache()
	if err != nil {
		log.Fatal("Error create template cache", err)
	}

	var app config.AppConfig
	app.TemplateCache = tmplCache

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("starting apllication on port %d\n", portNumber)

	// $ curl localhost:portNumber
	address := fmt.Sprintf(":%d", portNumber)
	_ = http.ListenAndServe(address, nil)
}
