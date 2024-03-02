package server

import (
	"log/slog"
	"net/http"
	"os"
)

var Mux = http.NewServeMux()

func StartServer() {
	// set up routes and start server
	err := registerRoutes(Mux)
	if err != nil {
		slog.Error("failed to register routes", "err", err)
		os.Exit(1)
	}
	err = http.ListenAndServe(":8080", Mux)
	if err != nil {
		slog.Error("http server failure", "err", err)
		os.Exit(1)
	}
}
