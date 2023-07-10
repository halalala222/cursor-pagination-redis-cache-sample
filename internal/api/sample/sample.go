package sample

import (
	"github.com/gin-gonic/gin"
	"github.com/halalala222/cursor-pagination-redis-cache-sample/internal/api/controller"
)

var router *gin.RouterGroup

func Run() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	router = engine.Group("/api")
	router.GET("/sample", controller.GetCursor)
	return engine
}
