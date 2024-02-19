package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/franz-dalitz/api-debugger/config"
	"github.com/franz-dalitz/api-debugger/routes"
)

func main() {
	// customize slog to use logfmt
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// load and log configuration
	cfg, err := config.LoadConfig("examples/config.yaml")
	if err != nil {
		slog.Error("failed to read config file")
		log.Fatal(err)
	}
	slog.Debug(fmt.Sprintf("successfully loaded the following config: %v", cfg))

	// set up routes and start server
	mux := http.NewServeMux()
	err = routes.RegisterRoutes(mux)
	if err != nil {
		slog.Error("failed to register routes")
		log.Fatal(err)
	}
	slog.Debug("routes have been registered")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		slog.Error("http handler failure")
		log.Fatal(err)
	}
}
