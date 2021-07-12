package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTmpl renders templates using html/template
func RenderTmpl(w http.ResponseWriter, tmpl string) {
	parsedTmpl, _ := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	err := parsedTmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template")
		return
	}
}
