package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/mohammadne/go_samples/modern_web_application/pkg/config"
	"github.com/mohammadne/go_samples/modern_web_application/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTmpl sets the config foe the template package
func NewTmpl(a *config.AppConfig) {
	app = a
}

func AddDefaultData(tmplData *models.TmplData) *models.TmplData {
	return tmplData
}

// RenderTmpl renders templates using html/template
func RenderTmpl(w http.ResponseWriter, name string, tmplData *models.TmplData) {
	var tmplSet map[string]*template.Template
	if app.UseCache {
		tmplSet = app.TemplateCache
	} else {
		tmplSet, _ = CreateTmplCache()
	}

	tmpl, ok := tmplSet[name]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	tmplData = AddDefaultData(tmplData)
	err := tmpl.Execute(buf, tmplData)
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal("Error writing template to browser", err)
	}

}

// CreateTmplCache creates a template cache as a map
func CreateTmplCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmplSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return nil, err
		}

		if len(matches) > 0 {
			tmplSet, err = tmplSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return nil, err
			}
		}

		cache[name] = tmplSet
	}

	return cache, nil
}
