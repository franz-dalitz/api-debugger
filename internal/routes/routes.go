package routes

import (
	"net/http"

	"github.com/franz-dalitz/api-debugger/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux) error {

	// serve htmx dependency
	mux.HandleFunc("/ui-deps", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/ui-deps/node_modules/htmx.org/dist/htmx.min.js" {
			http.ServeFile(w, r, "/ui-deps/node_modules/htmx.org/dist/htmx.min.js")
		} else if r.URL.Path == "/ui-deps/output.css" {
			http.ServeFile(w, r, "/ui-deps/output.css")
		}
	})

	// serve pages
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/hello", handlers.HelloHandler)

	return nil
}
