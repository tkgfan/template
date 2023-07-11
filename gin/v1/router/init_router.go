// author gmfan
// date 2023/07/01

package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	PingApi(r)
}
