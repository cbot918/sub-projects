package main

import (
	"os"

	u "github.com/cbot918/liby/util"
	// "github.com/cbot918/liby/cmdy"
)

func main(){
	args := os.Args
	
	switch args[1]{
	case "init " :{
		if args[2] == "." {
			u.Logg("....")
		}
	}
	}
}