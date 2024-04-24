// author gmfan
// date 2023/7/24
package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tkgfan/got/core/logx"
)

// TraceLog 处理链路日志
func TraceLog() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 设置链路日志
		tid := c.Request.Header.Get(logx.TraceCtxKey)
		tc := logx.NewTraceCtx(tid, c.ClientIP(), c.Request.URL.Path)
		ctx := context.WithValue(c.Request.Context(), logx.TraceCtxKey, tc)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
