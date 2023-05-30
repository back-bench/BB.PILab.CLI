package fileutility

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/back-bench/BB.PILab.CLI/pi/config"
	"github.com/back-bench/BB.PILab.CLI/pi/utility/httputility"
)

// CreateTextFile will take path, and array of string, one index will reprsent one line of file
func CreateTextFile(path string, content []string) (bool, error) {
	f, err := os.Create(path)
	var er error
	if err != nil {
		er = errors.New("error in creatring file")
		return false, er
	} else {
		for _, v := range content {
			fmt.Fprintln(f, v)
			if err != nil {
				er = errors.New("error in writing into file")
				return false, er
			}
		}
		err = f.Close()
		if err != nil {
			er = errors.New("error in closing file")
			return false, er
		}
	}
	return true, er
}

// CreateEnvJSONFile of type env content
func CreateEnvJSONFile(path string, content []config.Env) (bool, error) {
	jsonFile, err := os.Create(path)
	var er error
	if err != nil {
		er = errors.New("error in creating JSON file")
		return false, er
	}
	jsonWriter := io.Writer(jsonFile)
	encoder := json.NewEncoder(jsonWriter)
	err = encoder.Encode(&content)
	if err != nil {
		er = errors.New("error in encoding JSON to file")
		return false, er
	}
	return true, er
}

// DownloadFileFromGit for getting the file from EnvDock and create into give location
func DownloadFileFromGit(destination string, url string) (bool, error) {

	// Get the data
	resp, err := httputility.GetCall(url)
	var er error

	if err != nil {
		er = errors.New("something went wrong during connection establishment")
		return false, er
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {

		// Create the file
		out, err := os.Create(destination)
		if err != nil {
			er = errors.New("cant not create file in given destination")
			return false, er
		}
		defer out.Close()

		// Writer the body to file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			er = errors.New("erron in copying data from repositery")
			return false, er
		}

	} else {
		er = errors.New("connection failed or file not found")
		return false, er
	}
	return true, er
}
