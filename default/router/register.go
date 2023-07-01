// author gmfan
// date 2023/07/01

package router

import (
	v1 "acsupport/router/v1"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine){
	v1.PingApi(r)
}