package main

import (
	"log/slog"
	"os"

	"github.com/franz-dalitz/api-debugger/internal/config"
	"github.com/franz-dalitz/api-debugger/internal/server"
)

func main() {
	config.LoadSettings()
	err := config.LoadConfig()
	if err != nil {
		slog.Error("failed to read config file", "err", err)
		os.Exit(1)
	}
	server.StartServer()
}
