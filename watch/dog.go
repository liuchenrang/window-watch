package watch

import (
	"fmt"
	"time"

	"github.com/window-watch/config"
	"github.com/window-watch/contract"
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
	timer := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Printf("has %s \r\n", d.mgr.Has())
			for d.mgr.Has() {
				process := d.mgr.Get()
				fmt.Printf("process status %t %+v  \r\n", process.Alive(), process)

				if !process.Alive() {
					fmt.Printf("start program %+v\r\n", process)
					process.Start()
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
