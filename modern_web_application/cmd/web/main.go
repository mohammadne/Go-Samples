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

	fmt.Printf("starting apllication on port %d\n", portNumber)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", portNumber),
		Handler: routes(&appConfig),
	}

	// $ curl localhost:portNumber
	err = server.ListenAndServe()
	log.Fatal(err)
}
