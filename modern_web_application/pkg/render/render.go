package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	functions = template.FuncMap{}
)

// RenderTmpl renders templates using html/template
func RenderTmpl(w http.ResponseWriter, name string) {
	tmplCache, err := createTmplCache()
	if err != nil {
		log.Fatal("Error create template cache", err)
	}

	tmpl, ok := tmplCache[name]
	if !ok {
		log.Fatal("Error not exists in cache", err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, nil)
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal("Error writing template to browser", err)
	}

}

// createTmplCache creates a template cache as a map
func createTmplCache() (map[string]*template.Template, error) {
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
