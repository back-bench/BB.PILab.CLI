package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/back-bench/BB.PILab.CLI/pi/utility/httputility"
)

// Env type variabe
type Env struct {
	Name       string `json:"env_name"`
	GlobalPath string `json:"env_dockerfile_url"`
	LocalPath  string `json:"local_path"`
	Local      bool
	Global     bool
	ParentEnv  []string `json:"parent_env"` // An unexported field is not encoded.
	Created    string   `json:"created"`
	CreatedBy  string   `json:"created_by_name"`
	Version    string   //need to add in api``
	Status     string
}

// ReadEnvFile Read env file from json
func ReadEnvFile(path string) ([]Env, error) {
	jsonFile, _ := os.Open(path)

	byteValue, err := ioutil.ReadAll(jsonFile)
	var er error

	var data []Env
	jsonErr := json.Unmarshal(byteValue, &data)

	if jsonErr != nil {
		er = errors.New("Unable to unmarshal environment file")
	}
	if err != nil {
		er = errors.New("Unable to read environment file")
	}
	return data, er
}

// GetEnvFileContentByHTTPCall get the env file from json response
func GetEnvFileContentByHTTPCall(url string) ([]Env, error) {

	resp, err := httputility.GetCall(url)
	var data []Env
	var er error

	if err != nil {
		er = errors.New("Unable to establish connection")
	} else {

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			er = errors.New("Can't read http response")
		}

		jsonErr := json.Unmarshal(body, &data)
		if jsonErr != nil {
			er = errors.New("Unable unmarshal http contenet into Env object")
		}
	}
	return data, er
}

// GetEnvInfoByName return Env information as per name
func GetEnvInfoByName(path string, name string, version string) (*Env, error) {
	data, _ := ReadEnvFile(path)
	var result Env
	var er error
	for i := 0; i < len(data); i++ {
		if (data[i].Name == name) && (data[i].Version == version) {
			result = data[i]
			er = nil
		}
	}
	if result.Name == "" {
		er = errors.New("Environment is not avilable in local")
	}
	return &result, er
}
