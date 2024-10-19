package db

import (
	"context"

	redis "github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func InitDB(dataSourceName string) (*redis.Client, error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: dataSourceName,
	})

	// Verify connection
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return rdb, nil
}
