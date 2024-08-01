// author gmfan
// date 2023/7/8
package result

import (
	"anifun/common/cerr"
	"github.com/gin-gonic/gin"
	"github.com/tkgfan/got/core/errs"
	"github.com/tkgfan/got/core/model"
	"github.com/tkgfan/got/core/structs"
	"github.com/tkgfan/got/core/tlog"
	"net/http"
)

const (
	OK uint32 = 0
)

const (
	OkMsg = "OK"
)

// HttpResult 统一处理返回结果
func HttpResult(c *gin.Context, data any, err error) {
	if structs.IsNil(err) {
		tlog.CtxInfo(c.Request.Context(), "OK")
		c.JSON(http.StatusOK, model.NewSuccessResp(OK, OkMsg, data))
		return
	}

	tlog.CtxErrorf(c.Request.Context(), "%+v", err)

	// 处理错误
	cause := errs.Cause(err)
	if e, ok := cerr.IsCodeErr(cause); ok {
		c.JSON(http.StatusOK, model.NewFailResp(e.Code, e.Msg))
	} else {
		c.JSON(http.StatusOK, model.NewFailResp(cerr.UnknownErr, cerr.UnknownErrMsg))
	}
}
