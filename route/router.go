package route

import (
	"github.com/Gin-Web-demo/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 测试
	router.GET("/ping", api.Test)
	return router
}
