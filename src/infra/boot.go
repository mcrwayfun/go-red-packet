package infra

import "github.com/tietang/props/kvs"

type BootApplication struct {
	conf           kvs.ConfigSource
	starterContext StarterContext
}

func New(conf kvs.ConfigSource) *BootApplication {
	b := &BootApplication{conf: conf, starterContext: StarterContext{}}
	b.starterContext[KeyProps] = conf
	return b
}

func (b *BootApplication) Start() {
	//1:初始化所有的starter
	b.init()
	//2:安装starter
	b.setUp()
	//3:启动starter
	b.start()
}

func (b *BootApplication) init() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(b.starterContext)
	}
}

func (b *BootApplication) setUp() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(b.starterContext)
	}
}

func (b *BootApplication) start() {
	for index, starter := range StarterRegister.AllStarters() {
		if starter.StartBlocking() {
			if index+1 == len(StarterRegister.AllStarters()) {
				// 最后一个，阻塞
				starter.Start(b.starterContext)
			} else {
				go starter.Start(b.starterContext)
			}
		} else {
			starter.Start(b.starterContext)
		}
	}
}
