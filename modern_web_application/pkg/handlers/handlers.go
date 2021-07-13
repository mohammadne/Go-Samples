package handlers

import (
	"net/http"

	"github.com/mohammadne/go_samples/modern_web_application/pkg/config"
	"github.com/mohammadne/go_samples/modern_web_application/pkg/models"
	"github.com/mohammadne/go_samples/modern_web_application/pkg/render"
)

// Repo the repository used by the handlers (repository pattern)
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is home page handler
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTmpl(w, "home.page.tmpl", &models.TmplData{})
	// fmt.Fprintln(w, "welcome to simple home page of my web server")
}

// About is about handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	// do some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// sent that into template

	render.RenderTmpl(w, "about.page.tmpl", &models.TmplData{
		StringMap: stringMap,
	})
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
