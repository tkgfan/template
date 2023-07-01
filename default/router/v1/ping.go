// author gmfan
// date 2023/07/01

package v1

import "github.com/gin-gonic/gin"

// PingApi 注册测试路由
func PingApi(r *gin.Engine){
	g:=r.Group("/v1")
	g.GET("/ping",pong)
}

func pong(c *gin.Context){
	c.Writer.Write([]byte("pong"))
}