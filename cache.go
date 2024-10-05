package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // our password
		DB:       0,  // deualt db
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

// save urls in cache
func cacheURL(shortURL, longURL string) {
	err := rdb.Set(ctx, shortURL, longURL, 24*time.Hour).Err() // for 24 hours
	if err != nil {
		log.Printf("Could not cache URL: %v", err)
	}
}

// get from cache
func getURLFromCache(shortURL string) (string, bool) {
	longURL, err := rdb.Get(ctx, shortURL).Result()
	if err == redis.Nil {
		return "", false
	} else if err != nil {
		log.Printf("Could not get URL from cache: %v", err)
		return "", false
	}

	return longURL, true
}
