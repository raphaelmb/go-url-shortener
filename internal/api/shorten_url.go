package api

import "net/http"

type shortenURLRequest struct {
	URL string `json:"url"`
}

type shortenURLResponse struct {
	Code string `json:"code"`
}

func handleShortenURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
