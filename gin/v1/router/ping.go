// author gmfan
// date 2023/07/01

package router

import "github.com/gin-gonic/gin"

// PingApi 注册测试路由
func PingApi(r *gin.Engine) {
	r.GET("/ping", pong)
}

func pong(c *gin.Context) {
	c.Writer.Write([]byte("pong"))
}
