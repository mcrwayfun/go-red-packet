package base

import (
	"fmt"
	"github.com/tietang/props/kvs"
	"mcrwayfun.com/red/packet/src/infra"
)

var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	return props
}

type PropsStarter struct {
	infra.BaseStarter
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {
	props = ctx.Props()
	fmt.Println("PropsStarter 初始化配置成功")
}
