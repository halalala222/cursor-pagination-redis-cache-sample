package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var rdb *redis.Client

func init() {
	var err error
	redisAddr := fmt.Sprintf("%s:%s", "localhost", "6379")
	rdb = redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   0,
	})

	if _, err = rdb.Ping().Result(); err != nil {
		log.Fatal(err)
	}
}

func RDB(ctx context.Context) *redis.Client {
	return rdb.WithContext(ctx)
}
