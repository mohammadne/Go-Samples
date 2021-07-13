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
	var appConfig config.AppConfig

	tmplCache, err := render.CreateTmplCache()
	if err != nil {
		log.Fatal("Error create template cache", err)
	}

	appConfig.TemplateCache = tmplCache
	appConfig.UseCache = false

	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)

	render.NewTmpl(&appConfig)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("starting apllication on port %d\n", portNumber)

	// $ curl localhost:portNumber
	address := fmt.Sprintf(":%d", portNumber)
	_ = http.ListenAndServe(address, nil)
}
