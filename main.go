package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/franz-dalitz/api-debugger/internal/config"
	"github.com/franz-dalitz/api-debugger/internal/routes"
)

func main() {
	// customize slog to use logfmt
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	// define and parse flags
	cfgpath := flag.String("config", "/etc/api-debugger/config.yaml", "set path of config file")
	flag.Parse()
	slog.Debug(fmt.Sprintf("using '%v' as config file path", *cfgpath))

	// load and log configuration
	cfg, err := config.LoadConfig(*cfgpath)
	if err != nil {
		slog.Error("failed to read config file", "err", err)
		os.Exit(1)
	}
	slog.Debug(fmt.Sprintf("successfully loaded the following config: %v", cfg))

	// set up routes and start server
	mux := http.NewServeMux()
	err = routes.RegisterRoutes(mux)
	if err != nil {
		slog.Error("failed to register routes", "err", err)
		os.Exit(1)
	}
	slog.Debug("routes have been registered")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		slog.Error("http server failure", "err", err)
		os.Exit(1)
	}
}
