package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/halalala222/cursor-pagination-redis-cache-sample/consts"
	"github.com/halalala222/cursor-pagination-redis-cache-sample/internal/db"
	"log"
	"strconv"
)

type CacheConfig struct {
	Data   any
	Prefix string
	Unique any
}

func (c *CacheConfig) getFullKey() string {
	if c.Unique != nil {
		return fmt.Sprintf("%s-%s-%v", consts.BasicPrefix, c.Prefix, c.Unique)
	}
	return fmt.Sprintf("%s-%s", consts.BasicPrefix, c.Prefix)
}

func (c *CacheConfig) GetString(ctx context.Context) error {
	var (
		err    error
		result string
	)

	if result, err = db.RDB(ctx).Get(c.getFullKey()).Result(); err != nil {
		return err
	}
	if err = json.Unmarshal([]byte(result), c.Data); err != nil {
		return err
	}

	return nil
}

func (c *CacheConfig) SetString(ctx context.Context) {
	var (
		err      error
		byteData []byte
	)

	if byteData, err = json.Marshal(c.Data); err != nil {
		log.Println(err)
		return
	}

	if err = db.RDB(ctx).Set(c.getFullKey(), string(byteData), consts.DefaultSampleRedisTTL).Err(); err != nil {
		log.Println(err)
		return
	}
	return
}

func (c *CacheConfig) GetZRevRangeWithScoresWithMin(ctx context.Context, min int64) ([]redis.Z, error) {
	var (
		opt = redis.ZRangeBy{
			Max:    "+inf",
			Min:    strconv.FormatInt(min, 10),
			Offset: 0,
			Count:  consts.DefaultPageSize,
		}
	)
	return db.RDB(ctx).ZRevRangeByScoreWithScores(c.getFullKey(), opt).Result()
}

func (c *CacheConfig) GetZRevRangeWithScoresWithMax(ctx context.Context, max int64) ([]redis.Z, error) {
	var (
		opt = redis.ZRangeBy{
			Max:    strconv.FormatInt(max, 10),
			Min:    "-inf",
			Offset: 0,
			Count:  consts.DefaultPageSize,
		}
	)
	return db.RDB(ctx).ZRevRangeByScoreWithScores(c.getFullKey(), opt).Result()
}

func (c *CacheConfig) SetZSet(ctx context.Context, data []redis.Z) {
	var err error
	if err = db.RDB(ctx).ZAdd(c.getFullKey(), data...).Err(); err != nil {
		log.Println(err)
	}
	return
}
