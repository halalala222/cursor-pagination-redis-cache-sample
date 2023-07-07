package sample

import (
	"context"
	"github.com/halalala222/cursor-pagination-redis-cache-sample/consts"
	"github.com/halalala222/cursor-pagination-redis-cache-sample/internal/db"
)

type Sample struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (s *Sample) getCursor(ctx context.Context, cursorId int64) ([]Sample, error) {
	var cursorData = make([]Sample, 0)
	err := db.DB(ctx).Model(&Sample{}).Where("id > ?", cursorId).Limit(consts.DefaultPageSize).Find(&cursorData).Error
	return cursorData, err
}
