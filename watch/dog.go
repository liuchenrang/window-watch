package watch

import (
	"time"

	"github.com/liuchenrang/window-watch/config"
	"github.com/liuchenrang/window-watch/contract"
	. "github.com/liuchenrang/window-watch/logger"
)

type Dog struct {
	mgr contract.IManager
}

// func (d *Dog) SetProMgr(mgr *contract.Manager) {
// 	d.mgr = mgr
// }
func (d *Dog) SetProMgr(mgr contract.IManager) {
	d.mgr = mgr
}
func (d *Dog) Start(interval int, watch config.CWatch) {
	d.mgr.Init(watch)
	timer := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		select {
		case <-timer.C:
			Logger.Infof("has %s \r\n", d.mgr.Has())
			for d.mgr.Has() {
				process := d.mgr.Get()

				if !process.Alive() {
					Logger.Info("start program %+v\r\n", process)
					go process.Start()
				} else {
					Logger.Info("already runing %+v", process)
				}
				d.mgr.Next()
			}
			d.mgr.Reset()
		}
	}

}
func NewDog() Dog {
	return Dog{}
}
