package warcraftlogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// WarcraftLogs Client
type WarcraftLogs struct {
	client   http.Client
	apiToken string

	endpoint string
}

// New WarcraftLogs Client, with required API Token.
// Currently using v1 API endpoint as default
func New(apiToken string) *WarcraftLogs {

	api := &WarcraftLogs{
		endpoint: "https://www.warcraftlogs.com/v1/",
		client: http.Client{
			Timeout: time.Second * 2,
		},
		apiToken: apiToken,
	}

	return api
}

// SetEndpoint allows customizing API Endpoint that client will use.
// If provided endoint does not end with a trailing slash, it will be added
func (w *WarcraftLogs) SetEndpoint(endpoint string) {

	w.endpoint = endpoint

	if !strings.HasSuffix(w.endpoint, "/") {
		w.endpoint = w.endpoint + "/"
	}
}

// Support function to make an authenticated GET call and parse response JSON to a responseHolder.
func (w *WarcraftLogs) get(path string, responseHolder interface{}) error {

	url := fmt.Sprintf("%s%s?api_key=%s", w.endpoint, path, w.apiToken)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	res, getErr := w.client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
		return getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
		return readErr
	}

	jsonErr := json.Unmarshal(body, responseHolder)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		return jsonErr
	}

	return nil

}
