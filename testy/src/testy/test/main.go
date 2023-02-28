package main

import (
	"github.com/cbot918/testy/src/testy"
)


const (
	image="cbot918/ubugo"
	container="unit"
	bin="tt"
)
func main(){
	t := testy.New(image,container,bin)
	t.CiDockCreate()
}