// author gmfan
// date 2023/07/01

package main

import (
	"acsupport/common/middleware"
	"acsupport/conf"
	"acsupport/v1/router"
	"github.com/tkgfan/got/core/env"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	if env.CurModel == env.DevModel {
		gin.SetMode("debug")
	} else {
		gin.SetMode("release")
	}

	// 加载全局中间件
	loadMiddleware(r)

	// 注册路由
	router.InitRouter(r)

	// 运行服务器
	r.Run(":" + conf.Port)
}

func loadMiddleware(r *gin.Engine) {
	// 链路日志
	r.Use(middleware.TraceLog())
	r.Use(middleware.Timeout(time.Second * time.Duration(conf.Timeout)))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}