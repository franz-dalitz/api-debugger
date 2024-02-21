package routes

import (
	"fmt"
	"net/http"

	"github.com/franz-dalitz/api-debugger/internal/handlers"
	"github.com/franz-dalitz/api-debugger/web/dependencies"
)

func RegisterRoutes(mux *http.ServeMux) error {

	// serve htmx dependency
	mux.HandleFunc("/web/dependencies/*", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/web/dependencies/output.css" {
			w.Header().Add("Content-Type", "text/css")
			fmt.Fprint(w, dependencies.OutputCss)
		} else if r.URL.Path == "/web/dependencies/node_modules/htmx.org/dist/htmx.min.js" {
			fmt.Fprint(w, dependencies.HtmxJs)
		}
	})

	// serve pages
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/hello", handlers.HelloHandler)

	return nil
}
