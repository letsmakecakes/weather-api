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

// GetLimiter returns the rate limiter for a given IP, or creates one if it doesn't exist
func (r *RateLimiter) GetLimiter(ip string, rateLimit int, burst int) *rate.Limiter {
	r.mu.Lock()
	defer r.mu.Unlock()

	limiter, exists := r.clients[ip]
	if !exists {
		limiter = rate.NewLimiter(rate.Limit(rateLimit), burst)
		r.clients[ip] = limiter
	}

	return limiter
}
