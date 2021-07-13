package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/mohammadne/go_samples/modern_web_application/pkg/config"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	return mux
}
