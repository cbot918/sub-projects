package main

import (
	"fmt"

	"github.com/cbot918/liby/cmdy"
	"github.com/cbot918/liby/configy"
	u "github.com/cbot918/liby/util"
)

type Config struct {
	Test struct {
		// Vagrant string
		Bin string
		ContainerName string
		Port int32
	}
}



func main(){
	c := cmdy.New()
	
	cfg := &Config{}
	configy.New().Set("./","app","yaml").Get(cfg)
	// u.Loggj(cfg)
	// u.Logg(cfg.Test.ContainerName)


	copyBin := []string{
		// build binary process
		// fmt.Sprintf("docker run -dit --name %s cbot918/ubugo",cfg.Test.ContainerName),
		fmt.Sprintf("docker cp %s %s:.",cfg.Test.Bin, cfg.Test.ContainerName),
	}
	c.Run(copyBin)

	execBin := []string{fmt.Sprintf("docker exec %s bash -c \"./%s init .\"", cfg.Test.ContainerName, cfg.Test.Bin),}
	u.Logg(c.RunAndGet(execBin),)

	rmBin := []string{fmt.Sprintf("docker exec %s bash -c \"rm %s\"", cfg.Test.ContainerName, cfg.Test.Bin),}
	c.Run(rmBin)


	
	// docker exec ci-dock-testy bash -c "./ya init ."
}

// ci-doc-run: $(BNAME)
// 	docker run -dit --name $(CINAME)  cbot918/ubugo
// 	docker cp $(BNAME) $(CINAME):.
// 	rm $(BNAME)
// 	docker exec $(CINAME) bash -c "./$(BNAME) install docker"
// 	docker exec -it $(CINAME) bash

// ci-doc-clean:
// 	docker stop $(CINAME)
// 	docker rm $(CINAME)