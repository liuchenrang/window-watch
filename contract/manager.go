package contract

import "github.com/window-watch/config"

type IManager interface {
	Init(config config.CWatch)
	Register()
	Has() bool
	Next()
	Reset()
	Get() IProcess
}
