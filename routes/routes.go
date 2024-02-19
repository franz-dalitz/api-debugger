package routes

import (
	"net/http"

	"github.com/franz-dalitz/api-debugger/handlers"
)

func RegisterRoutes(mux *http.ServeMux) error {
	// serve static directory
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// serve node module dependencies
	mux.HandleFunc("/node_modules/htmx.org/dist/htmx.min.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "node_modules/htmx.org/dist/htmx.min.js")
	})

	// serve home template
	mux.HandleFunc("/", handlers.RootHandler)

	return nil
}
