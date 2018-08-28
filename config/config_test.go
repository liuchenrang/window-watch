package config

import (
	"fmt"
	"testing"
)

func Test_NewConfig(t *testing.T) {
	watch := NewConfig("/usr/local/Cellar/go/gopath/src/github.com/window-watch/watch.yaml")
	fmt.Printf("%+v xx", watch)
}
