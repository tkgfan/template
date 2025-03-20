// author gmfan
// date 2023/07/01

package conf

import (
	"github.com/tkgfan/got/core/env"
	"github.com/tkgfan/got/core/tlog"
)

var (
	Port      = "8888"
	Timeout   = 60 * 1000
	LogLevel  = tlog.InfoLevel
	ApiPrefix = ""
)

func init() {
	ginTemplateInit(false)
}

func ginTemplateInit(must bool) {
	env.LoadStr(&Port, "PORT", must)
	env.LoadInt(&Timeout, "TIMEOUT", must)
	env.LoadStr(&LogLevel, "LOG_LEVEL", must)
	tlog.SetLevel(LogLevel)
}
