package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/window-watch/config"
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
// func Output2(r contract.IProcess) {
// 	r.Start()
// }
var (
	configFlag = flag.String("c", "", "config file path")
)

func main() {
	flag.Parse()

	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	defaultConfigPath := path + "/watch.yaml"
	configPath := *configFlag
	if *configFlag == "" {
		configPath = defaultConfigPath
	}
	// rd := IntRead{t: 4}
	// Output(rd)

	// var lt contract.IProcess
	// lt = &IntRead{}
	// Output2(lt)

	config := config.NewWatch(configPath)
	dog := watch.NewDog()
	mgr := process.NewMgr()
	dog.SetProMgr(mgr)
	dog.Start(config.Interval, config)
}
