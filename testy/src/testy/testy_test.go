package testy

import (
	"fmt"
	"os/exec"
	"testing"

	// "github.com/cbot918/liby/cmdy"

	u "github.com/cbot918/liby/util"
	"github.com/stretchr/testify/require"
)

type TestTesty struct {}

func NewTestTesty(){
	
}


const (
	image="cbot918/ubugo"
	container="unit"
	bin="tt"
)
var testy *Testy

func TestMain(t *testing.T){
	testy = New(image,container,bin)
	fmt.Println(testy)
}
 
func TestCiDockCreate(t *testing.T) {
	err := New(image,container,bin).CiDockCreate()
	require.NoError(t, err)
}

func TestCiDockCopyBin(t *testing.T){
	testy = New(image,container,bin)
	err := testy.CiDockCopyBin()
	require.NoError(t, err)
}

func TestCiDockExecCode(t *testing.T){
	testy = New(image,container,bin)
	// err := testy.CiDockExecCode("ls")
	cmd := exec.Command("/usr/bin/sh -c ls")
	err := cmd.Run()
	u.Checke(err, "run error")
	// cmdy.New().Run([]string{fmt.Sprintf("docker run -dit --name %s %s",container,image),})
	// cmdy.New().Run([]string{fmt.Sprintf("ls"),})
	// require.NoError(t, err)
}

func TestCiDockExecCodeAndGet(t *testing.T){
	testy = New(image,container,bin)
	result, err := testy.CiDockExecCodeAndGet("ls")
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestCiDockDelete(t *testing.T){
	testy = New(image,container,bin)
	err := testy.CiDockDelete()
	require.NoError(t, err)
}
