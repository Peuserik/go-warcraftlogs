package warcraftlogs

import (
	"github.com/AlexejK/go-warcraftlogs/types/warcraft"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWarcraftLogs_ReportsForGuild(t *testing.T) {

	ts, api := testServerAndClient(t, "reports.json")
	defer ts.Close()

	reports := api.ReportsForGuild("GuildName", warcraft.Realm_Silvermoon, warcraft.Region_EU)

	assert.Equal(t, 2, len(reports), "Unexpected amount of reports")

	first := reports[0]

	assert.Equal(t, "Highmaul Heroic", *first.Title)
	assert.Equal(t, "wK8qcnc49PakFvNJ", *first.Id)
	assert.Equal(t, "user2", *first.Owner)

}
