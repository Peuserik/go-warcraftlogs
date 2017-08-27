package warcraftlogs

import (
	"github.com/stretchr/testify/assert"

	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWarcraftLogs_BaseCall(t *testing.T) {

	// Test Data
	testApiKey := "fake-token"
	testRequestPath := "some/endpoint"
	testResponseJSON := `["one", "two"]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		apiKeyParam := r.URL.Query().Get("api_key")

		assert.NotEmpty(t, apiKeyParam, "API Key must be sent to the server")
		assert.Equal(t, testApiKey, apiKeyParam, "Invalid API key sent")

		// Write Response
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, testResponseJSON)
	}))

	api := New(testApiKey)
	api.SetEndpoint(ts.URL)

	data := []*string{}
	api.get(testRequestPath, &data)

	assert.Equal(t, 2, len(data), "Parsing of response data failed")

	first := *data[0]
	assert.Equal(t, "one", first, "Expected first element of response to be 'one'")
}
