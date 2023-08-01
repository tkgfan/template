// author gmfan
// date 2023/7/24
package middleware

import (
	"acsupport/common/errs"
	"acsupport/common/result"
	"acsupport/common/tlog"
	"github.com/gin-gonic/gin"
)

// TraceLog 处理链路日志
func TraceLog() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, err := tlog.UpdateTraceLog(c.Request.Context(), c.Request.Header.Get(tlog.TraceLogKey), c.Request.URL.Path)
		if err != nil {
			tlog.PrintlnError(err)
			result.HttpResult(c, nil, errs.NewCodeErrMgs(errs.ParamErr, "未设置链路日志"))
			c.Abort()
			return
		}
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
