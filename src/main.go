package main

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	_ "mcrwayfun.com/red/packet/src/app"
	"mcrwayfun.com/red/packet/src/infra"
)

func main() {
	// 获取程序运行文件所在的路径
	file := kvs.GetCurrentFilePath("/brun/config.ini", 1)
	// 加载配置文件
	conf := ini.NewIniFileCompositeConfigSource(file)

	app := infra.New(conf)
	app.Start()
}
