package limiters

import (
	"fmt"
	"time"
)

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
