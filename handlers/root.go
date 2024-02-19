package handlers

import (
	"net/http"
	"text/template"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}
	templ := template.Must(template.ParseFiles("templates/pages/home.html"))
	templ.Execute(w, nil)
}
