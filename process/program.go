package process

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/liuchenrang/window-watch/contract"

	"github.com/liuchenrang/window-watch/config"
	. "github.com/liuchenrang/window-watch/logger"
)

type Program struct {
	info       config.CProgram
	status     int
	startTimes int
	timer      *time.Ticker
	mode       string
}

func (p *Program) Stop() {

}
func (p *Program) Start() {
	p.startTimes++
	if p.startTimes <= 1 || p.CanTryStart() {
		env := os.Environ()
		procAttr := &os.ProcAttr{
			Env: env,
			Files: []*os.File{
				os.Stdin,
				os.Stdout,
				os.Stderr,
			},
		}
		pid, err := os.StartProcess(p.info.Path, []string{}, procAttr)
		if err != nil {
			fmt.Printf("Error %v starting process!", err) //
		} else {
			fmt.Printf("The process id is %v", pid)
		}
	} else {
		Logger.Infof("wait start %+v", p.info)
	}

}
func (p *Program) CanTryStart() bool {
	return p.startTimes%3 == 0
}
func (p *Program) Alive() bool {
	//checkAlive
	if processHas(runtime.GOOS, p.info.Name) {
		p.status = contract.STARTING
	} else {
		p.status = contract.STOPING
	}
	return p.status != contract.STOPING
}
func (p *Program) Status() int {
	return p.status
}
