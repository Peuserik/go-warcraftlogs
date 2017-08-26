package warcraftlogs

import (
	"fmt"
	"github.com/AlexejK/go-warcraftlogs/types/warcraft"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWarcraftLogs_ReportsForGuild(t *testing.T) {

	// Test Data
	testApiKey := "fake-token"
	testResponseJSON := `[
   {"id":"wK8qcnc49PakFvNJ","title":"Highmaul Heroic","owner":"user2","start":1418679807631,"end":1418680839199,"zone":6},
   {"id":"kCpxrcaAbaXDP3RV","title":"ToS Heroic","owner":"user1","start":1502992911089,"end":1503003469088,"zone":13}
]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Write Response
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, testResponseJSON)
	}))
	defer ts.Close()

	api := New(testApiKey)
	api.SetEndpoint(ts.URL)

	reports := api.ReportsForGuild("GuildName", warcraft.Realm_Silvermoon, warcraft.Region_EU)

	for _, r := range reports {
		fmt.Printf("Report (%s): %s by %s \n", *r.Id, *r.Title, *r.Owner)
	}

}
