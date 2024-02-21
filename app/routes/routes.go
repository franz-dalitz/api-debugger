package routes

import (
	"net/http"

	"github.com/franz-dalitz/api-debugger/app/handlers"
)

func RegisterRoutes(mux *http.ServeMux) error {
	// serve static directory
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// serve htmx dependency
	mux.HandleFunc("/node_modules/htmx.org/dist/htmx.min.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "node_modules/htmx.org/dist/htmx.min.js")
	})

	// serve pages
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/hello", handlers.HelloHandler)

	return nil
}
