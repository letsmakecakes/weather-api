package middleware

import (
	"sync"
	"weatherapi/pkg/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimiter tracks rate limiters per client IP
type RateLimiter struct {
	clients   map[string]*rate.Limiter
	mu        sync.Mutex
	rateLimit int
	burst     int
}

// NewRateLimiter creates a new rate limiter instance
func NewRateLimiter(rateLimit int, burst int) *RateLimiter {
	return &RateLimiter{
		clients:   make(map[string]*rate.Limiter),
		rateLimit: rateLimit,
		burst:     burst,
	}
}

// GetLimiter returns the rate limiter for a given IP, or creates one if it doesn't exist
func (r *RateLimiter) GetLimiter(ip string) *rate.Limiter {
	r.mu.Lock()
	defer r.mu.Unlock()

	limiter, exists := r.clients[ip]
	if !exists {
		limiter = rate.NewLimiter(rate.Limit(r.rateLimit), r.burst)
		r.clients[ip] = limiter
	}

	return limiter
}

// RateLimitMiddleware is a Gin middleware for rate limiting
func (r *RateLimiter) RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		limiter := r.GetLimiter(ip)
		if !limiter.Allow() {
			utils.RespondWithError(c, 429, "Too many request")
			c.Abort()
			return
		}

		c.Next()
	}
}
