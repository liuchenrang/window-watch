package main

import (
	"fmt"
	"strconv"

	"github.com/window-watch/config"
	"github.com/window-watch/contract"
	"github.com/window-watch/process"
	"github.com/window-watch/watch"
)

// type Read interface {
// 	read(i int) string
// }
type IntRead struct {
	t int
}

func (r IntRead) read(i int) string {
	return strconv.Itoa(i * r.t)
}
func (r *IntRead) Start() {
	fmt.Printf("start ")
}

// func Output(r Read) {
// 	fmt.Printf("read %s", r.read(3))
// }
func Output2(r contract.IProcess) {
	r.Start()
}

func main() {
	// rd := IntRead{t: 4}
	// Output(rd)

	// var lt contract.IProcess
	// lt = &IntRead{}
	// Output2(lt)
	file := "/usr/local/Cellar/go/gopath/src/github.com/window-watch/watch.yaml"

	config := config.NewWatch(file)
	dog := watch.NewDog()
	mgr := process.NewMgr()
	dog.SetProMgr(mgr)
	dog.Start(3, config)
}
