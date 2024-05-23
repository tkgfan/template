// author gmfan
// date 2023/7/24
package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tkgfan/got/core/tlog"
)

// TraceLog 处理链路日志
func TraceLog() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 设置链路日志
		tid := c.Request.Header.Get(tlog.CtxTidKey)
		if tid == "" {
			tid = tlog.NewTid()
		}
		ctx := context.WithValue(c.Request.Context(), tlog.CtxTidKey, tid)
		tlog.CtxInfof(ctx, "ip=%s path=%s", c.ClientIP(), c.Request.URL.Path)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		if c.Writer.Status() != 200 {
			tlog.CtxErrorf(ctx, "ip=%s path=%s status=%d", c.ClientIP(), c.Request.URL.Path, c.Writer.Status())
		}
	}
}
