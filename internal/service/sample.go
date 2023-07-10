package service

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/halalala222/cursor-pagination-redis-cache-sample/consts"
	"github.com/halalala222/cursor-pagination-redis-cache-sample/internal/model/sample"
	"github.com/halalala222/cursor-pagination-redis-cache-sample/internal/pkg"
	"github.com/samber/lo"
	"log"
	"strconv"
)

type SampleService struct{}

func (s *SampleService) GetOneSample(ctx context.Context, id int64) *sample.Sample {
	var (
		cache  = s.GetStringRedisConfig(id)
		result = &sample.Sample{}
	)
	if err := cache.GetString(ctx); err != nil {
		log.Println(err)
		var (
			newError error
			data     *sample.Sample
		)
		if data, newError = result.GetOne(ctx, id); newError != nil {
			log.Println(err)
			return data
		}
		cache.Data = data
		cache.SetString(ctx)
		return data
	}
	return cache.Data.(*sample.Sample)
}

func (s *SampleService) GetCursor(ctx context.Context, inputSample sample.Sample) []sample.Sample {
	var (
		cacheConfig = s.GetZSetRedisConfig()
		result      = make([]redis.Z, 0)
		err         error
		data        = make([]sample.Sample, 0)
		sampleData  = &sample.Sample{
			Id: inputSample.Id,
		}
	)

	if result, err = cacheConfig.GetZRevRangeWithScoresWithMax(ctx, inputSample.Id); err != nil {
		log.Println(err)
	}

	var (
		isGetLessData        = len(result) < consts.DefaultPageSize
		isGetFromRedisFailed = err != nil
		isKeyExist           = cacheConfig.KeyIsExit(ctx)
	)

	if isGetFromRedisFailed || !isKeyExist {
		if data, err = sampleData.GetCursor(ctx, sampleData.Id); err != nil {
			log.Println(err)
			return make([]sample.Sample, 0)
		}
		cacheConfig.SetZSet(ctx, lo.Map(data, func(item sample.Sample, index int) redis.Z {
			return redis.Z{
				Score:  float64(item.Id),
				Member: item.Id,
			}
		}))
		return data
	}

	data = lo.Map(result, func(item redis.Z, index int) sample.Sample {
		var parseInt int64
		if parseInt, err = strconv.ParseInt(item.Member.(string), 10, 64); err != nil {
			log.Println(err)
			return sample.Sample{}
		}
		return *s.GetOneSample(ctx, parseInt)
	})

	if isGetLessData {
		var (
			lastRecordId int64
			newData      []sample.Sample
		)

		if len(newData) != 0 {
			if lastRecordId, err = strconv.ParseInt(result[len(result)-1].Member.(string), 10, 64); err != nil {
				log.Println(err)
				return make([]sample.Sample, 0)
			}
			if newData, err = sampleData.GetCursor(ctx, lastRecordId); err != nil {
				log.Println(err)
				return make([]sample.Sample, 0)
			}
			cacheConfig.SetZSet(ctx, lo.Map(data, func(item sample.Sample, index int) redis.Z {
				return redis.Z{
					Score:  float64(item.Id),
					Member: item.Id,
				}
			}))
			data = append(data, newData...)
		}
	}

	return data
}

func (s *SampleService) GetZSetRedisConfig() *pkg.CacheConfig {
	cacheConfig := &pkg.CacheConfig{
		Prefix: consts.SampleZSetPrefix,
		Unique: nil,
		Data:   &[]sample.Sample{},
	}
	return cacheConfig
}

func (s *SampleService) GetStringRedisConfig(unique int64) *pkg.CacheConfig {
	cacheConfig := &pkg.CacheConfig{
		Prefix: consts.SamplePrefix,
		Unique: unique,
		Data:   &sample.Sample{},
	}
	return cacheConfig
}
