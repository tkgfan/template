// author gmfan
// date 2023/07/01

package main

import (
	"anifun/common/middleware"
	"anifun/conf"
	"anifun/router"
	"github.com/tkgfan/got/core/env"
	"net/http"
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
	srv := &http.Server{
		Addr:         ":" + conf.Port,
		Handler:      r,
		ReadTimeout:  time.Duration(conf.Timeout) * time.Millisecond,
		WriteTimeout: time.Duration(conf.Timeout) * time.Millisecond,
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func loadMiddleware(r *gin.Engine) {
	// 链路日志
	r.Use(middleware.TraceLog())
	// 超时设置
	r.Use(middleware.Timeout(time.Millisecond * time.Duration(conf.Timeout)))
	// 解析 Request 中 Header 数据并设置到上下文中
	r.Use(middleware.SetCtxData())
	// 非开发环境
	if env.CurModel != env.DevModel {
		// 校验签名
		r.Use(middleware.VerifyGatewayToken())
	}
	r.Use(gin.Recovery())
}
