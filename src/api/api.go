package api

import (
	"net/http"

	"github.com/DevKayoS/goMovies/src/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(apiKey string) http.Handler {
	router := chi.NewMux()

	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Get("/search-movie", services.HandleSearchMovie(apiKey))

	return router
}
