// author gmfan
// date 2023/07/01

package main

import (
	"acsupport/conf"
	"acsupport/router"

	"github.com/gin-gonic/gin"
)

func main(){
	r:=gin.New()
	gin.SetMode(conf.Mode)
	
	// 注册路由
	router.Register(r)

	// 运行服务器
	r.Run(":"+conf.Port)
}