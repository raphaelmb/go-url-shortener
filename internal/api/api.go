package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type apiResponse struct {
	Error string `json:"error,omitempty`
	Data  any    `json:"data,omitempty`
}

func NewHandler() http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer, middleware.RequestID, middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Route("/url", func(r chi.Router) {
			r.Post("/shorten", handleShortenURL())
			r.Get("/{code}", handleGetShortenedURL())
		})
	})

	return r
}
