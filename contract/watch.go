package contract

import "github.com/liuchenrang/window-watch/config"

type IWatch interface {
	SetProMgr(mgr *IManager)
	Start(interval int, config config.CWatch)
}
