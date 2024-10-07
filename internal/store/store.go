package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type store struct {
	rdb *redis.Client
}

type Store interface {
	SaveShortenedURL(ctx context.Context, _url string) (string, error)
	GetFullURL(ctx context.Context, code string) (string, error)
}

func NewStore(rdb *redis.Client) Store {
	return store{rdb: rdb}
}

func (s store) SaveShortenedURL(ctx context.Context, _url string) (string, error) {
	var code string
	for range 5 {
		code = genCode()
		if err := s.rdb.HGet(ctx, "shortener", code).Err(); err != nil {
			if errors.Is(err, redis.Nil) {
				break
			}
			return "", fmt.Errorf("failed to get code from shortener hashmap: %w", err)
		}

	}

	if err := s.rdb.HSet(ctx, "shortener", code, _url).Err(); err != nil {
		return "", fmt.Errorf("failed to set code on shortener hashmap: %w", err)
	}

	return code, nil
}

func (s store) GetFullURL(ctx context.Context, code string) (string, error) {
	fullURL, err := s.rdb.HGet(ctx, "shortener", code).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get code from shortener hashmap: %w", err)
	}
	return fullURL, nil
}
