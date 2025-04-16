package main

import (
	"log/slog"
	"os"
)

func main() {
	if err := App(); err != nil {
		slog.Error("failed to execute code", "error", err)
		os.Exit(1)
	}

	slog.Info("all system  offline")
}
