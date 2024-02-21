package handlers

import (
	"net/http"
	"text/template"

	"github.com/franz-dalitz/api-debugger/web/templates"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}
	templ := template.Must(template.ParseFS(templates.Get(), "pages/home.tmpl"))
	templ.Execute(w, nil)
}
