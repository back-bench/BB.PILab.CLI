package httputility

import (
	"errors"
	"net/http"
)

// GetCall do get call
func GetCall(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	resp, _ := http.DefaultClient.Do(req)
	var er error
	if err != nil {
		er = errors.New("Error in fetching data from given url")
	}
	return resp, er
}
