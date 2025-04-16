package main

import "log/slog"

func main() {
	if err := App(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}

	slog.Info("all system  offline")
}
