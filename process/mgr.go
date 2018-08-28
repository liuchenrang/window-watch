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
		fmt.Printf("k %s v %+v \r\n", k, v)
		pp := Program{}
		pp.info = v
		pp.timer = time.NewTicker(1 * time.Second)
		pp.Alive()
		m.programs = append(m.programs, pp)
	}
	fmt.Println("register %+v", m.programs)
	m.processTotal = len(m.programs)
}
func (m *ProMgr) Has() bool {
	if m.cursor >= m.processTotal {
		return false
	}
	return true
}
func processHas(name string) bool {
	cli := "ps aux|grep " + name + "|grep -v grep"
	cmd := exec.Command("/bin/bash", "-c", cli)
	var out bytes.Buffer
	var err bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &err
	ee1 := cmd.Run()
	fmt.Printf("process run cli  %s \r\n", cli)

	fmt.Printf("process check stdOut %s \r\n", out.String())
	fmt.Printf("process check stdErr %s \r\n", err.String())
	fmt.Printf("process run ee1  %s \r\n", ee1)
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
	mgr := ProMgr{programs: make([]Program, 0)}
	return &mgr
}
