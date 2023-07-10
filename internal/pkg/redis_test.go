package pkg

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/halalala222/cursor-pagination-redis-cache-sample/internal/model/sample"
	"testing"
)

var setData = []redis.Z{
	{
		Score:  1,
		Member: 1,
	},
	{
		Score:  2,
		Member: 2,
	},
	{
		Score:  3,
		Member: 3,
	},
	{
		Score:  4,
		Member: 4,
	},
	{
		Score:  5,
		Member: 5,
	},
	{
		Score:  6,
		Member: 6,
	},
	{
		Score:  7,
		Member: 7,
	},
	{
		Score:  8,
		Member: 8,
	},
	{
		Score:  9,
		Member: 9,
	},
	{
		Score:  10,
		Member: 10,
	},
	{
		Score:  11,
		Member: 11,
	},
	{
		Score:  12,
		Member: 12,
	},
	{
		Score:  13,
		Member: 13,
	},
	{
		Score:  14,
		Member: 14,
	},
	{
		Score:  15,
		Member: 15,
	},
}

var testConfig = []CacheConfig{
	{
		Data:   nil,
		Prefix: "test-get-full-key",
		Unique: 1,
	},
	{
		Data: []sample.Sample{
			{
				Id:   1,
				Name: "test-1",
			},
			{
				Id:   2,
				Name: "test-2",
			},
		},
		Prefix: "test-get-full-key",
		Unique: 1,
	},
	{
		Data:   &[]sample.Sample{},
		Prefix: "test-get-full-key",
		Unique: 1,
	},
	{
		Data:   &[]sample.Sample{},
		Prefix: "test-ZSet",
		Unique: 2,
	},
	{
		Data:   &[]sample.Sample{},
		Prefix: "test",
		Unique: nil,
	},
}

func TestCacheConfig_GetFullKey(t *testing.T) {
	t.Log(testConfig[0].getFullKey())
}

func TestCacheConfig_SetString(t *testing.T) {
	testConfig[1].SetString(context.Background())
}

func TestCacheConfig_GetString(t *testing.T) {
	err := testConfig[2].GetString(context.Background())
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(testConfig[2].Data)
}

func TestCacheConfig_SetZSet(t *testing.T) {
	testConfig[3].SetZSet(context.Background(), setData)
}

func TestCacheConfig_GetZRevRangeWithScoresWithMin(t *testing.T) {
	var (
		data []redis.Z
		err  error
	)
	if data, err = testConfig[3].GetZRevRangeWithScoresWithMin(context.Background(), 10); err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
}

func TestCacheConfig_GetZRevRangeWithScoresWithMax(t *testing.T) {
	var (
		data []redis.Z
		err  error
	)
	if data, err = testConfig[3].GetZRevRangeWithScoresWithMax(context.Background(), 14); err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
}

func TestCacheConfig_KeyIsExit(t *testing.T) {
	isExit := testConfig[0].KeyIsExit(context.Background())
	if isExit {
		t.Error("expect false by get true")
	}
	isExit = testConfig[4].KeyIsExit(context.Background())
	if !isExit {
		t.Error("expect true by get false")
	}
}
