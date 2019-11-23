package app

import (
	"mcrwayfun.com/red/packet/src/infra"
	"mcrwayfun.com/red/packet/src/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
}
