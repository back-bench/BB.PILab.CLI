package main

import (
	"fmt"
	"os"

	"github.com/back-bench/BB.PILab.CLI/pi/command"
)

func main() {
	//updated code
	//create new command object
	piCli, err := command.NewPiCli()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	//file := amzCli.ConfigFile()
	//filee := amzCli.EnvConfigFile()

	if err := runPi(piCli); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
