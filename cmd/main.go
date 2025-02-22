package main

import (
	"log/slog"

	"github.com/oktavarium/doit-bot/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		slog.Error("error running server", slog.Any("Error", err))
		return
	}
}
