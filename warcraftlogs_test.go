package warcraftlogs

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWarcraftLogs_Call(t *testing.T) {

	// Test Data
	testApiKey := "fake-token"
	testRequestPath := "some/endpoint"
	testResponseJSON := `["one", "two"]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		apiKeyParam := r.URL.Query().Get("api_key")
		if apiKeyParam == "" {
			t.Errorf("API Key must be sent to the server")
		}

		if apiKeyParam != testApiKey {
			t.Errorf("Invalid API Key sent. Expecting '%s', got '%s'.", testApiKey, apiKeyParam)
		}

		// Write Response
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, testResponseJSON)
	}))

	api := New(testApiKey)
	api.SetEndpoint(ts.URL)

	data := &[]*string{}
	api.get(testRequestPath, data)
}
