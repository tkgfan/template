// Description: 网关校验
// Author: gmfan
// Date: 2024/8/13

package middleware

import (
	"anifun/common/cerr"
	"anifun/common/result"
	"github.com/gin-gonic/gin"
	"github.com/tkgfan/got/core/env"
	"github.com/tkgfan/got/core/tlog"
)

const (
	GatewayTokenKey = "Gateway-Token"
)

var (
	// GatewayToken 网关 token
	GatewayToken = "123456"
)

func init() {
	// 加载签名秘钥
	env.LoadStr(&GatewayToken, "GATEWAY_TOKEN", env.CurModel != env.DevModel)
}

// VerifyGatewayToken 校验网关 token
func VerifyGatewayToken() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		reqToken := c.Request.Header.Get(GatewayTokenKey)
		if reqToken != GatewayToken {
			tlog.CtxErrorf(ctx, "网关 Token 错误,reqToken: %s", reqToken)
			result.HttpResult(c, nil, cerr.NewCodeErr(cerr.UnknownErr))
			c.Abort()
			return
		}

		c.Next()
	}
}
