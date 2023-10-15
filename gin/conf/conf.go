// author gmfan
// date 2023/07/01

package conf

import (
	"github.com/tkgfan/got/core/env"
)

var (
	Port    = "8888"
	Timeout = 60 * 1000
)

func init() {
	must := env.CurModel != env.DevModel

	env.LoadStr(&Port, "PORT", must)
}
