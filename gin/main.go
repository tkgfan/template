// author gmfan
// date 2023/07/01

package main

import (
	"acsupport/conf"
	"acsupport/v1/router"
	"github.com/tkgfan/got/core/env"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	if env.CurModel == env.DevModel {
		gin.SetMode("debug")
	} else {
		gin.SetMode("release")
	}

	// 注册路由
	router.InitRouter(r)

	// 运行服务器
	r.Run(":" + conf.Port)
}
