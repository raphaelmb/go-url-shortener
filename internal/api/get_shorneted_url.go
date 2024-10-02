package api

import "net/http"

type getShortenedURLResponse struct {
	FullURL string `json:"full_url"`
}

func handleGetShortenedURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
