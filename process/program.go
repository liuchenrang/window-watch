package process

import (
	"bytes"
	"os/exec"
	"strings"
	"time"

	"github.com/window-watch/contract"

	"github.com/window-watch/config"
	. "github.com/window-watch/logger"
)

type Program struct {
	info       config.CProgram
	status     int
	startTimes int
	timer      *time.Ticker
}

func (p *Program) Stop() {

}
func (p *Program) Start() {
	p.startTimes++
	if p.startTimes <= 1 || p.CanTryStart() {
		cmd := exec.Command(p.info.Name)
		cmd.Stdin = strings.NewReader("/tmp")
		var out bytes.Buffer
		var stdErr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stdErr
		err := cmd.Run()
		Logger.Infof("run path %s", p.info.Path)
		if err != nil {
			Logger.Errorf("error %s", err)
		} else {
			Logger.Infof("result \r\n %s", out.String())
		}
		errLen := len(stdErr.String())
		if errLen > 0 {
			Logger.Errorf("result stdErr \r\n  %s", stdErr.String())
		} else {
			Logger.Infof("result no error ")
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
	if processHas(p.info.Name) {
		p.status = contract.STARTING
	} else {
		p.status = contract.STOPING
	}
	return p.status != contract.STOPING
}
func (p *Program) Status() int {
	return p.status
}
