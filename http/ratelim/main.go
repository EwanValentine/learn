package main

import (
	"fmt"
	"net/http"
	"time"

	rredis "github.com/go-redis/redis/v9"
	// "ratelim/internal/redis"
)

func RateLimiter(client *rredis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Do something
	}
}

type InMemoryFixedWindowCounter struct {
	storage map[string]int
	rps     int
}

func NewInMemoryFixedWindowCounter(rps int) *InMemoryFixedWindowCounter {
	return &InMemoryFixedWindowCounter{
		storage: make(map[string]int),
		rps:     rps,
	}
}

func (c *InMemoryFixedWindowCounter) Inc(key string) {
	c.storage[key]++
}

func (c *InMemoryFixedWindowCounter) IsAllowed(clientID string) bool {
	currentTime := time.Now().Unix()
	key := fmt.Sprintf("%s:%d", clientID, currentTime)

	requestCount, ok := c.storage[key]
	if !ok {
		c.storage[key] = 1
		return true
	}

	if requestCount > c.rps {
		return false
	}

	c.storage[key]++

	return true
}

func main() {
	// rclient, err := redis.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":8080", mux)
}
