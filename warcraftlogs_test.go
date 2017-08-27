package warcraftlogs

import (
	"io/ioutil"
	"path/filepath"
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
)

func loadTestData(t *testing.T, name string) []byte {
	path := filepath.Join("testData", name) // relative path
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}

func loadTestDataString(t *testing.T, name string) string {
	return string(loadTestData(t, name))
}

func testServerAndClient(t *testing.T, testFile string) (*httptest.Server, *WarcraftLogs) {

	testApiKey := "fake-token"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, loadTestDataString(t, testFile))
	}))

	api := New(testApiKey)
	api.SetEndpoint(ts.URL)

	return ts, api
}
