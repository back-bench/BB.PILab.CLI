package config

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// EnvConfig for env information
type EnvConfig struct {
	Env        bool
	EnvName    string
	EnvAvl     string
	EnvVersion string
}

// AuthConfig for authentication imformation
type AuthConfig struct {
	Auth bool
	LKey string
	GKey string
}

// DirConfig work dir information
type DirConfig struct {
	Dir     bool
	WorkDir string
}

// GitIntLvl git intergration level information
type GitIntLvl struct {
	Commit bool
	Pull   bool
	Push   bool
}

// GitCred git information
type GitCred struct {
	Git         bool
	GitUserName string
	GitPass     string
	GitAccPerm  bool
	GitIntLevel GitIntLvl
}

// Config config information
type Config struct {
	FileName string
	Envr     EnvConfig
	Auth     AuthConfig
	Dir      DirConfig
	Git      GitCred
}

func (config *Config) generateContent(auth bool, git bool, Dir bool, Env bool, path string) Config {
	key := strings.Replace(path, "/", "", -1)
	//path level information
	config.FileName = "config"
	if auth {
		config.Auth.Auth = true
		config.Auth.LKey = key
		config.Auth.GKey = "nil" // need to check authentication and geraate
	} else if !auth {
		config.Auth.Auth = false
		config.Auth.LKey = path
		config.Auth.GKey = "nil"
	}

	//as of now it'll always false, so assigning everthing false
	config.Git.Git = git
	config.Git.GitIntLevel.Commit = false
	config.Git.GitIntLevel.Pull = false
	config.Git.GitIntLevel.Push = false
	config.Git.GitUserName = "nil"
	config.Git.GitPass = "nil"
	config.Git.GitAccPerm = false

	//env level information
	if Env {
		//alway false in first time
		// need to get env information

	} else if !Env {
		config.Envr.Env = false
		config.Envr.EnvName = "nil"
		config.Envr.EnvAvl = "nil"
		config.Envr.EnvVersion = "nil"
	}

	if Dir {
		config.Dir.Dir = true
		config.Dir.WorkDir = path
	}

	return *config
}

// WritableConfigContent return writeble content
func WritableConfigContent(conf Config) []string {
	var content []string
	if conf.Auth.Auth {
		content = append(content, "[Auth]")
		content = append(content, "LKey="+conf.Auth.LKey)
		content = append(content, "GKey="+conf.Auth.GKey)
	}
	if conf.Dir.Dir {
		content = append(content, "[Dir]")
		content = append(content, "WDir="+conf.Dir.WorkDir)
	}
	if conf.Git.Git {
		content = append(content, "[Git]")
		content = append(content, "GUserName="+conf.Git.GitUserName)
		content = append(content, "GPass="+conf.Git.GitPass)
		content = append(content, "GAccPerm="+strconv.FormatBool(conf.Git.GitAccPerm))
		content = append(content, "GCommit="+strconv.FormatBool(conf.Git.GitIntLevel.Commit))
		content = append(content, "GPull="+strconv.FormatBool(conf.Git.GitIntLevel.Pull))
		content = append(content, "GPush="+strconv.FormatBool(conf.Git.GitIntLevel.Push))
	}
	if conf.Envr.Env {
		content = append(content, "[Env]")
		content = append(content, "EName="+conf.Envr.EnvName)
		content = append(content, "EAvl="+conf.Envr.EnvAvl)
		content = append(content, "EVersion="+conf.Envr.EnvVersion)
	}
	return content
}

// ReadConfigFile function will return readble format.ReadConfigFile
func ReadConfigFile(path string) (*Config, error) {
	var conf Config
	var er error
	lines, err := readLines(path)
	var hashMap = make(map[string]string)
	//fmt.Println("it's getting call")
	//fmt.Println("Path", path)
	//fmt.Println("Line count", len(lines))
	if err == nil {
		//process the each line and map to confi abject
		for _, line := range lines {
			if line == "[Env]" {
				conf.Envr.Env = true
				continue
			} else if line == "[Dir]" {
				conf.Dir.Dir = true
				continue
			} else if line == "[Auth]" {
				conf.Auth.Auth = true
				continue
			} else if line == "[Git]" {
				conf.Git.Git = true
			} else {
				var s []string
				s = readValue(line)
				if len(s) == 2 {
					hashMap[s[0]] = s[1]
				} else {
					er = errors.New("config File is broken")
					break
				}
			}
		}
	}

	//go through the object and if we have key store into object.
	if conf.Auth.Auth {
		conf.Auth.GKey = hashMap["GKey"]
		conf.Auth.LKey = hashMap["LKey"]
	}
	if conf.Dir.Dir {
		conf.Dir.WorkDir = hashMap["WDir"]
	}
	if conf.Envr.Env {
		conf.Envr.EnvName = hashMap["EName"]
		conf.Envr.EnvAvl = hashMap["EAvl"]
		conf.Envr.EnvVersion = hashMap["EVersion"]
	}
	if conf.Git.Git {
		conf.Git.GitUserName = hashMap["UserName"]
		conf.Git.GitPass = hashMap["GPass"]
		conf.Git.GitAccPerm = false
		conf.Git.GitIntLevel.Commit = false
		conf.Git.GitIntLevel.Pull = false
		conf.Git.GitIntLevel.Push = false
	}

	return &conf, er
}

func readValue(line string) []string {
	return strings.Split(line, "=")
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	return lines, scanner.Err()
}

// CreateConfig return created config file
func CreateConfig(path string, intial bool) Config {
	var conf Config
	if intial {
		conf = conf.generateContent(true, false, true, false, path)
	}

	return conf
}

// UpdateAndInsertEnv into config file
func UpdateAndInsertEnv(conf Config, envName string, version string) []string {
	//update fields
	conf.Envr.Env = true
	conf.Envr.EnvName = envName
	conf.Envr.EnvVersion = version
	return WritableConfigContent(conf)
}
