package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/liuchenrang/window-watch/config"
	"github.com/liuchenrang/window-watch/process"
	"github.com/liuchenrang/window-watch/watch"
)

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

	config := config.NewWatch(configPath)
	dog := watch.NewDog()
	mgr := process.NewMgr()
	dog.SetProMgr(mgr)
	dog.Start(config.Interval, config)
}
