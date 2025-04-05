package main

import (
	"github.com/OliverSchlueter/sloki/sloki"
	"log/slog"
)

func main() {
	logger := sloki.NewService(sloki.Configuration{
		URL:          "http://localhost:3100/loki/api/v1/push",
		Service:      "my-service",
		ConsoleLevel: slog.LevelDebug,
		LokiLevel:    slog.LevelInfo,
	})
	slog.SetDefault(slog.New(logger))
}
