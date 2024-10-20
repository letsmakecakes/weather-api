package middleware

import (
	"sync"

	"golang.org/x/time/rate"
)

// RateLimiter tracks rate limiters per client IP
type RateLimiter struct {
	clients map[string]*rate.Limiter
	mu      sync.Mutex
}

// NewRateLimiter creates a new rate limiter instance
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		clients: make(map[string]*rate.Limiter),
	}
}
