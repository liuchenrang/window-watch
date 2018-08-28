package config

import (
	"io/ioutil"

	. "github.com/liuchenrang/window-watch/logger"
	"gopkg.in/yaml.v2"
)

func NewWatch(file string) CWatch {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	watch := CWatch{}
	err = yaml.Unmarshal(content, &watch)
	if err != nil {
		panic(err)
	}
	Logger.Infof("config %+v", watch)
	return watch
}
