package testy

import (
	"fmt"
	"os"

	"github.com/cbot918/liby/cmdy"
	u "github.com/cbot918/liby/util"
)

type Testy struct {
	cmdy 			cmdy.Cmdy
	image 		string
	container string
	bin 			string
}

func New(image string, container string, bin string) *Testy{
	t := new(Testy)
	t.cmdy = *cmdy.New()
	t.container= container
	t.image = image
	t.bin = bin
	
	return t
}

func (t *Testy) Run(){
	args := os.Args
	if len(args) == 1 {
		fmt.Println("welcome to Testy !!")
	}else{
		switch args[1] {
		
			case "init":{
				fmt.Println("create config file: testy.conf.yaml")
				t.cmdy.Run([]string{"touch testy.conf.yaml",})
			}

			case "ci-dock": {
				if len(args) == 2{
					fmt.Println("./tt ci-dock")
				} else{
					switch args[2]{
					case "create":{ t.CiDockCreate() }
					case "setup":{ t.cmdy.Run([]string{fmt.Sprintf("docker cp %s %s:.",t.bin,t.container),}) }
					case "run":{ result,_ := t.CiDockExecCodeAndGet("test");fmt.Println(result) }
					case "teardown": { t.cmdy.Run([]string{fmt.Sprintf("docker exec %s bash -c \"rm %s\"",t.container,t.bin),}) }
					case "delete":{ t.CiDockDelete()}
					case "test":{ fmt.Println("test success !!") }
					default :fmt.Println("./tt cidock unknown")
					}				
				}
			}
			
			default: fmt.Printf("./tt %s\n", args[1])
		}
	}
}


func (t *Testy) CiDockCreate() error {
	err := t.cmdy.Run([]string{fmt.Sprintf("docker run -dit --name %s %s",t.container,t.image),})
	u.Checke(err, "err in testy.go: fn CiDockCreate")
	return err
}

func (t *Testy) CiDockCopyBin() error {
	err := t.cmdy.Run([]string{fmt.Sprintf("docker cp %s %s:.",t.image,t.container),})
	u.Checke(err, "err in testy.go: fn CiDockCopyBin")
	return err
}

func (t *Testy) CiDockExecCode(code string) error {
	err := t.cmdy.Run([]string{fmt.Sprintf("docker exec %s bash -c \"./%s %s\"",t.container, t.bin, code)})
	u.Checke(err, "err in testy.go: fn CiDockExecCode")
	return err
}

func (t *Testy) CiDockExecCodeAndGet(code string) (string,error){
	result, err := t.cmdy.RunAndGet([]string{fmt.Sprintf("docker exec %s bash -c \"./%s %s\"",t.container, t.bin, code)})
	u.Checke(err, "err in testy.go: fn CiDockExecCodeAndGet")
	return result, err
}

func (t *Testy) CiDockDelete() error {
	err := t.cmdy.Run([]string{
		fmt.Sprintf("docker container stop %s",t.container),
		fmt.Sprintf("docker container rm %s",t.container),
	})
	u.Checke(err, "err in testy.go: fn CiDockDelete")
	return err
}
