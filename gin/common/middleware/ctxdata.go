// author gmfan
// date 2023/8/12

package middleware

import (
	"anifun/common/cerr"
	"anifun/common/ctxdata"
	"anifun/common/result"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tkgfan/got/core/tlog"
	"strconv"
)

// SetCtxData 解析 Header 数据并设置到 Request 的上下文中
func SetCtxData() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		// 获取 UID
		uidStr := c.Request.Header.Get(ctxdata.UIDKey)
		if uidStr != "" {
			uid, err := strconv.ParseInt(uidStr, 10, 64)
			if err != nil {
				tlog.CtxErrorf(ctx, "uidStr: %s,err: %+v", uidStr, err)
				result.HttpResult(c, nil, cerr.NewCodeErr(cerr.NotLogin))
				c.Abort()
				return
			}
			ctx = context.WithValue(ctx, ctxdata.UIDKey, uid)
		}

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
