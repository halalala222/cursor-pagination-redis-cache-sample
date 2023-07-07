package sample

import "github.com/gin-gonic/gin"

var router *gin.RouterGroup

func Run() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	router = engine.Group("/api")
	return engine
}

func RootRouter() *gin.RouterGroup {
	return router
}
