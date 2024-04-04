package main

import (
	"net/http"

	rredis "github.com/go-redis/redis/v9"
	// "ratelim/internal/redis"
)

func RateLimiter(client *rredis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Do something
	}
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
