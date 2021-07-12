package render

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	functions = template.FuncMap{}
)

// RenderTmpl renders templates using html/template
func RenderTmpl(w http.ResponseWriter, tmpl string) {
	_, err := RenderTmplTest(w)
	if err != nil {
		fmt.Println("Error getting template cache", err)
		return
	}

	parsedTmpl, _ := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	err = parsedTmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

func RenderTmplTest(w http.ResponseWriter) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("page is currently", page)
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
