package init

import (
	"fmt"
	"os"
	"strings"

	"github.com/back-bench/BB.PILab.CLI/pi/command"
	"github.com/back-bench/BB.PILab.CLI/pi/config"
	"github.com/back-bench/BB.PILab.CLI/pi/utility/crypt"
	"github.com/back-bench/BB.PILab.CLI/pi/utility/fileutility"
	"github.com/spf13/cobra"
)

// AMZInitCommand '.' sub comaand required.
func PIInitCommand(cli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   ".",
		Short: "initialize current and sub directories for PI Lab.",
		//Args:  cli.RequiresMaxArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				fmt.Println("no argument required")
				//options.optionName = args[0]
			}
			return Init(cli)
		},
	}
	return cmd
}

// Init initialize the amz
func Init(cli command.Cli) error {
	//copied code from prev version
	//need to do changes
	//get current directory
	path := cli.GetWDir()
	if path == "" {
		fmt.Println("\033[0;31m[PI]\033[0m Having trouble to get current directory.") // print error msg
	} else {
		path = path + "/.amz" //this shouldn't be there, if it is there project is already intitated for this directory.
		if checkSetup(path) {
			startProjectSetup(path)
		} else {
			fmt.Println("\033[0;31m[PI]\033[0m Looks like PI Lab is already assigned to this directory.")
		}
	}

	return nil
}

// very low level validation
func checkSetup(path string) bool {
	//currently we are just checking .amz directory
	dirInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true
	}
	if dirInfo != nil {
		return false
	}
	return false
}

// modification required
func startProjectSetup(path string) {
	fmt.Println("\033[0;32m[PI]\033[0m Setup is Started for current directory.....")
	//create .amz directory
	os.MkdirAll(path, 0777)
	//create config file content
	key := []byte(strings.Replace(path, "/", "", -1))

	//get config file

	var configContent = config.WritableConfigContent(config.CreateConfig(path, true))
	file, err := fileutility.CreateTextFile(path+"/config", configContent)
	if err == nil && file {
		fmt.Println("\033[0;32m[PI]\033[0m CONFIG file has been created successfully.....")
	} else {
		fmt.Println("\033[0;33m[PI]\033[0m Having trauble in creating config file.....")
		haltProjectSetup()
	}
	//create env folder
	fmt.Println("\033[0;32m[PI]\033[0m Environment Setup process has been started....")
	os.MkdirAll(path+"/envs", 0777)
	//get the json data from factory
	jsonData := getEnvJSON(path, key)
	// create main file inside env folder
	jsn, err := fileutility.CreateEnvJSONFile(path+"/envs/main.json", jsonData)
	// insert the global env to main file
	if err == nil && jsn {
		fmt.Println("\033[0;32m[PI]\033[0m Environment setup completed successfully.....")
		fmt.Println("\033[0;32m[PI]\033[0m\033[0;36m Quantam provide these environments as a sample environment, you can get more public global envoirnment using \033[0;94menv\033[0m\033[0;36m command.\033[0m")
		fmt.Println("\033[0;32m[PI]\033[0m\033[0;36m Failed environment setup can be recovered, this dosen't empact on overall AMZ setup process.\033[0m")
	} else {
		fmt.Println("\033[0;33m[PI]\033[0m Having trauble in creating Environments.....")
		haltProjectSetup()
	}

	fmt.Println("\033[0;32m[PI]\033[0m Setup has been completed.")
}

// get Env json via http call
func getEnvJSON(path string, key []byte) []config.Env {
	//var data []envconf.Env
	//get all the by making http call
	fmt.Println("\033[0;32m[PI]\033[0m Fetching Quantam default environment....")

	fmt.Println("\033[0;32m[PI]\033[0m Downloading environments from Quantam global repository.")
	var lurl = "http://localhost:8888/api/amz/environments/getQuantamPublicEnv/"
	data, err := config.GetEnvFileContentByHTTPCall(lurl)

	fmt.Println("\033[0;32m[PI]\033[0m Extracting environments information.")
	if err != nil {
		fmt.Println("\033[0;32m[PI]\033[0m", err.Error(), "try using \033[0;94menv\033[0m command.")
	} else {
		fmt.Println("\033[0;32m[PI]\033[0m Environments has been downloaded.")
	}

	fmt.Println("\033[0;32m[PI]\033[0m Environments has been extracted.")

	for i := 0; i < len(data); i++ {
		fmt.Println("\033[0;32m[PI]\033[0m\033[0;33m", data[i].Name, "\033[0m setup has been started.")
		ciphertext, err := crypt.Encrypt(key, []byte(data[i].GlobalPath))
		//text, err:= decrypt(key,ciphertext) for normal text
		if err != nil {
			fmt.Println(err) //need to handle
		} else {
			data[i].GlobalPath = string(ciphertext)
		}
		data[i].LocalPath = path + "/envs/" + strings.Replace(data[i].Name, " ", "", -1)
		if createEnvSetup(data[i]) {
			data[i].Status = "up-to-date"
			data[i].Global = true
			data[i].Local = true
		} else {
			data[i].Status = "broken"
			data[i].Global = true
			data[i].Local = false
		}
		data[i].Version = "latest"
	}

	return data
}

// setup env for each envs
func createEnvSetup(data config.Env) bool {

	//create folder
	os.MkdirAll(data.LocalPath, 0777)
	a, _ := fileutility.DownloadFileFromGit(data.LocalPath+"/Dockerfile", string("https://raw.githubusercontent.com/QuantamDataLab/Env_dock/master/"+data.Name+"/Dockerfile"))
	b, _ := fileutility.DownloadFileFromGit(data.LocalPath+"/requirements.txt", string("https://raw.githubusercontent.com/QuantamDataLab/Env_dock/master/"+data.Name+"/requirements.txt"))

	if !(a && b) {
		fmt.Println("\033[0;32m[PI]\033[0m\033[0;91m", data.Name, "\033[0m setup is not completed yet. try using \033[0;94menv\033[0m command.")
		return false
	}

	fmt.Println("\033[0;32m[PI]\033[0m\033[0;33m", data.Name, "\033[0m is up-to-date.")
	return true
}

// nned to write in better way
func haltProjectSetup() {
	fmt.Println("\033[0;33m[PI]\033[0m Setup has not been completed due to internal errors.")
}
