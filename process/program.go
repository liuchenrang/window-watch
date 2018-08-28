package process

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/window-watch/contract"

	"github.com/window-watch/config"
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
		fmt.Println("run path %s", p.info.Path)
		if err != nil {
			fmt.Errorf("error %s", err)
		} else {
			fmt.Printf("result %s", out.String())
		}
		errLen := len(stdErr.String())
		if errLen > 0 {
			fmt.Println("result stdErr %s", stdErr.String())
		} else {
			fmt.Println("result no error ")
		}
	} else {
		fmt.Println("wait start %+v", p.info)
	}

}
func (p *Program) CanTryStart() bool {
	return true
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
