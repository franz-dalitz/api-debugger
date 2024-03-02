package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/franz-dalitz/api-debugger/web/assets"
	"github.com/franz-dalitz/api-debugger/web/templates"
)

func assetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/web/assets/output.css" {
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprint(w, assets.OutputCss)
	} else if r.URL.Path == "/web/assets/htmx.min.js" {
		fmt.Fprint(w, assets.HtmxJs)
	} else {
		errorHandler(w, r, 404)
		return
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, 404)
		return
	}

	templHome := template.Must(template.ParseFS(templates.Get(), "pages/home.tmpl", "layouts/base.tmpl"))
	templHome.ExecuteTemplate(w, "base", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		errorHandler(w, r, 404)
		return
	}
	fmt.Fprint(w, "hello world")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
