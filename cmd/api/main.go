package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/raphaelmb/go-url-shortener/internal/api"
	"github.com/raphaelmb/go-url-shortener/internal/store"
	"github.com/redis/go-redis/v9"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	store := store.NewStore(rdb)

	handler := api.NewHandler(store)

	s := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server running on port 8080...")
	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
