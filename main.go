package main

import (
	"github.com/Gin-Web-demo/config"
	"github.com/Gin-Web-demo/model"
	"github.com/Gin-Web-demo/route"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 载入路由
	r := route.InitRouter()

	// 载入配置
	config.InitConfig()

	// 创建数据库连接
	model.InitDb()

	// 监听8000端口
	gin.SetMode(viper.GetString("server.run_mode"))
	r.Run(viper.GetString("server.addr"))
}
