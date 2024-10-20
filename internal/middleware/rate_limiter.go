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
