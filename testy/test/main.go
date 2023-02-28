package main

import (
	// "github.com/cbot918/liby/configy"
	// u "github.com/cbot918/liby/util"
	"github.com/cbot918/testy/src/testy"
)

type Config struct {
	App struct {
		Vagrant string
		Other string
	}
}

const (
	image="cbot918/ubugo"
	container="ci-dock-testy"
	bin="tt"
)

func main(){
	t := testy.New(image,container,bin)
	t.Run()

	// cfg := &Config{}
	// configy.New().Set("./","app","yaml").Get(cfg)
	// u.Loggj(cfg)

}