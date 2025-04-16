package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/DevKayoS/goMovies/src/api"
)

// rodar com OMDB_KEY=chave_api go run .
func App() error {
	apiKey := os.Getenv("OMDB_KEY")
	slog.Info("🚀Server running in port: 8080🚀")
	handler := api.NewHandler(apiKey)

	app := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := app.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
