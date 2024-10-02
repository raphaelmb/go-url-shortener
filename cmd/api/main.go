package main

import (
	"log"
	"net/http"
	"time"

	"github.com/raphaelmb/go-url-shortener/internal/api"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	handler := api.NewHandler()

	s := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
