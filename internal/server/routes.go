package server

import (
	"net/http"
)

func registerRoutes(mux *http.ServeMux) error {

	// serve htmx dependency
	mux.HandleFunc("/web/assets/*", assetHandler)

	// serve pages
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/hello", helloHandler)

	return nil
}
