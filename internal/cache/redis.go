package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type WeatherCache interface {
	SetCache(key string, value string, expiration time.Duration) error
	GetCache(key string) (string, error)
}

type cache struct {
	rdb *redis.Client
	ctx context.Context
}

func NewCache(rdb *redis.Client, ctx context.Context) WeatherCache {
	return &cache{rdb, ctx}
}

func (c *cache) GetCache(key string) (string, error) {
	return c.rdb.Get(c.ctx, key).Result()
}

func (c *cache) SetCache(key string, value string, expiration time.Duration) error {
	return c.rdb.Set(c.ctx, key, value, expiration).Err()
}
