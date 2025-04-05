# Sloki

A slog handler (for GoLang) which sends logs to Loki.

You can find a docker-compose file to run Loki and Grafana in the [loki_docker](loki_docker) directory. There is also an example dashboard for Grafana.

## Installation

```bash
go get github.com/OliverSchlueter/sloki
```

## Usage

```go
package main

import (
	"github.com/OliverSchlueter/sloki/sloki"
	"log/slog"
)

func main() {
	logger := sloki.NewService(sloki.Configuration{
		URL:          "http://localhost:3100/loki/api/v1/push",
		ConsoleLevel: slog.LevelDebug,
		LokiLevel:    slog.LevelInfo,
	})
	slog.SetDefault(slog.New(logger))
}
```