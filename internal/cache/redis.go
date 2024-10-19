package cache

import "time"

type WeatherCache interface {
	SetCache(key string, value string, expiration time.Duration) error
	GetCache(key string) (string, error)
}