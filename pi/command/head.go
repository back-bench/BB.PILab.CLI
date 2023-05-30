package command

import (
	"fmt"
	"io"
	"os"

	"github.com/back-bench/BB.PILab.CLI/pi/config"
)

// Cli is for important functions
type Cli interface {
	ConfigFile() *config.Config
	EnvConfigFile() []config.Env
	EnvInfo(name string, version string) *config.Env
	GetWDir() string
	Err() io.Writer
}

// PiCli contains
type PiCli struct {
	DirPath       string
	configFile    *config.Config
	envConfigFile []config.Env
	envInfo       *config.Env
	err           io.Writer
}

// GetWDir return  directory
func (cli *PiCli) GetWDir() string {
	return cli.DirPath
}

// ConfigFile return config file
func (cli *PiCli) ConfigFile() *config.Config {
	if cli.configFile == nil {
		cli.getConfigFile()
	}
	return cli.configFile
}

// EnvConfigFile get envconfig file
func (cli *PiCli) EnvConfigFile() []config.Env {
	if cli.envConfigFile == nil {
		//fmt.Println("env is coming from kit") every time this is getting call that means cli is intiating on each command.
		cli.getEnvConfigFile()
	} else {
		//fmt.Println("already came just returning.")
	}
	return cli.envConfigFile
}

// EnvInfo get the env infotmation
func (cli *PiCli) EnvInfo(name string, version string) *config.Env {
	if cli.envInfo == nil {
		cli.getEnvInformation(name, version)
	}
	return cli.envInfo
}

func (cli *PiCli) getEnvInformation(name string, version string) {
	cli.envInfo, _ = config.GetEnvInfoByName(cli.GetWDir()+"/.amz/envs/main.json", name, version) //for testing
}

func (cli *PiCli) getConfigFile() {
	//path is needed
	cli.configFile, _ = config.ReadConfigFile(cli.GetWDir() + "/.amz/config") //for testing
}

func (cli *PiCli) getEnvConfigFile() {
	//path is needed
	cli.envConfigFile, _ = config.ReadEnvFile(cli.GetWDir() + "/.amz/envs/main.json") //for testing
}

// Err wd
func (cli *PiCli) Err() io.Writer {
	return cli.err
}

// NewPiCli intiate AMZ CLI command for one valid directory address
// proof of validty of directry yet to discovered
func NewPiCli() (*PiCli, error) {
	//each time you run some command this is getting call, it shouldn't be like this.
	//fmt.Println("this is getting call")
	cli := &PiCli{}
	path, err := os.Getwd()

	//whatever path we got from here, first we'll validate whether it's valid or not
	//validater package is dedicated to this only

	if err != nil {
		fmt.Println("\033[0;31m[AMZ]\033[0m Having trouble to get current directory.") // print error msg
	}
	cli.DirPath = path //for testing
	return cli, nil
}

// ShowHelp show help if argument is null
func ShowHelp() {
	fmt.Println("Need help")
}
