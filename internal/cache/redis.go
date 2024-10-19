package cache

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type WeatherCache interface {
	SetCache(key string, value string, expiration time.Duration) error
	GetCache(key string) (string, error)
}

type cache struct {
	rdb *redis.Client
}

func NewCache(rdb *redis.Client) WeatherCache {
	return &cache{rdb}
}
