package db

import (
	"context"

	redis "github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var CTX = context.Background()

func InitDB(dataSourceName string) (*redis.Client, error) {
	RDB = redis.NewClient(&redis.Options{
		Addr: dataSourceName,
	})

	// Verify connection
	if _, err := RDB.Ping(CTX).Result(); err != nil {
		return nil, err
	}

	return RDB, nil
}
