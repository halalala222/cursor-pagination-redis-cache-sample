package service

import (
	"context"
	"github.com/halalala222/cursor-pagination-redis-cache-sample/internal/model/sample"
	"testing"
)

func TestSampleService_GetCursor(t *testing.T) {
	var sampleService = &SampleService{}
	data := sampleService.GetCursor(context.Background(), sample.Sample{
		Id: 5,
	})
	t.Log(data)
}
