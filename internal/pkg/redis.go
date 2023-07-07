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
	data   any
	prefix string
	unique any
}

func (c *CacheConfig) GetFullKey() string {
	return fmt.Sprintf("%s-%s-%v", consts.BasicPrefix, c.prefix, c.unique)
}

func (c *CacheConfig) GetString(ctx context.Context) error {
	var (
		err    error
		result string
	)

	if result, err = db.RDB(ctx).Get(c.GetFullKey()).Result(); err != nil {
		return err
	}
	if err = json.Unmarshal([]byte(result), c.data); err != nil {
		return err
	}

	return nil
}

func (c *CacheConfig) SetString(ctx context.Context) {
	var (
		err      error
		byteData []byte
	)

	if byteData, err = json.Marshal(c.data); err != nil {
		log.Println(err)
		return
	}

	if err = db.RDB(ctx).Set(c.GetFullKey(), string(byteData), consts.DefaultSampleRedisTtl).Err(); err != nil {
		log.Println(err)
		return
	}
	return
}

func (c *CacheConfig) GetZRevRangeWithScoresWithMin(ctx context.Context, min int64) error {
	var (
		result []redis.Z
		err    error
		opt    = redis.ZRangeBy{
			Max:    "+inf",
			Min:    strconv.FormatInt(min, 10),
			Offset: 0,
			Count:  consts.DefaultPageSize,
		}
	)
	if result, err = db.RDB(ctx).ZRevRangeByScoreWithScores(c.GetFullKey(), opt).Result(); err != nil {
		return err
	}
	c.data = result
	return nil
}

func (c *CacheConfig) SetZSet(ctx context.Context, data []redis.Z) {
	var err error
	if err = db.RDB(ctx).ZAdd(c.GetFullKey(), data...).Err(); err != nil {
		log.Println(err)
	}
	return
}
