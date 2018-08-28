package process

import (
	"bytes"
	"os/exec"
	"strings"
	"time"

	"github.com/liuchenrang/window-watch/config"
	"github.com/liuchenrang/window-watch/contract"
	. "github.com/liuchenrang/window-watch/logger"
)

type ProMgr struct {
	programs     []Program
	config       config.CWatch
	cursor       int
	processTotal int
}

func (m *ProMgr) Init(config config.CWatch) {
	m.config = config
	m.Register()
	m.cursor = 0
}
func (m *ProMgr) Register() {
	for k, v := range m.config.Programs {
		Logger.Info("k %s v %+v \r\n", k, v)
		pp := Program{}
		pp.mode = m.config.Mode
		pp.info = v
		pp.timer = time.NewTicker(1 * time.Second)
		pp.Alive()
		m.programs = append(m.programs, pp)
	}
	m.processTotal = len(m.programs)
	Logger.Infof("register %d %+v", m.processTotal, m.programs)

}
func (m *ProMgr) Has() bool {
	if m.cursor >= m.processTotal {
		return false
	}
	return true
}
func processHas(mode string, name string) bool {
	var cmd *exec.Cmd
	var cli string
	if mode == "linux" {
		cli = "ps aux|grep " + name + "|grep -v grep"
		cmd = exec.Command("/bin/bash", "-c", cli)
	} else {
		cli = "tasklist|find " + name + ""
		cmd = exec.Command("cmd.exe", "/C", cli)

	}
	var out bytes.Buffer
	var err bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &err
	ee1 := cmd.Run()
	Logger.Infof("process run cli  %s \r\n", cli)

	Logger.Infof("process check stdOut \r\n %s \r\n", out.String())
	Logger.Infof("process check stdErr \r\n  %s \r\n", err.String())
	if ee1 != nil {
		Logger.Errorf("process run ee1  %s \r\n", ee1)
	}
	return strings.Contains(out.String(), name)
}
func (m *ProMgr) Next() {
	m.cursor++
}
func (m *ProMgr) Reset() {
	m.cursor = 0
}
func (m *ProMgr) Get() contract.IProcess {
	pr := m.programs[m.cursor]
	return &pr
}
func NewMgr() *ProMgr {
	mgr := ProMgr{programs: make([]Program, 0), processTotal: 0, cursor: 0}
	return &mgr
}
