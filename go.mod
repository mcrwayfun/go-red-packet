module mcrwayfun.com/red/packet

go 1.12

replace (
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190322080309-f49334f85ddc
)

require (
	github.com/go-ini/ini v1.51.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/tietang/go-utils v0.1.3 // indirect
	github.com/tietang/props v2.2.0+incompatible
	github.com/valyala/fasttemplate v1.1.0 // indirect
)
