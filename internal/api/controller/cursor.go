package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/halalala222/cursor-pagination-redis-cache-sample/internal/model/form"
	"github.com/halalala222/cursor-pagination-redis-cache-sample/internal/model/sample"
	service2 "github.com/halalala222/cursor-pagination-redis-cache-sample/internal/service"
	"net/http"
)

func GetCursor(ctx *gin.Context) {
	var (
		service service2.SampleService
		data    form.CursorSample
	)
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.AsciiJSON(http.StatusBadRequest, gin.H{
			"code":    2,
			"message": "参数错误",
		})
	}
	samples := service.GetCursor(ctx, sample.Sample{
		Id: data.Id,
	})
	ctx.AsciiJSON(http.StatusOK, gin.H{
		"code": 1,
		"data": samples,
	})
}
