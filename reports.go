package warcraftlogs

import (
	"fmt"
	"github.com/peuserik/go-warcraftlogs/types"
	"github.com/peuserik/go-warcraftlogs/types/warcraft"
)

// ReportsForGuild returns a list of all reports available for provided guild.
// Warcraftlogs retention time of reports differs depending on reports tier.
func (w *WarcraftLogs) ReportsForGuild(guildName string, realm warcraft.Realm, region warcraft.Region) []*types.Report {

	url := fmt.Sprintf("reports/guild/%s/%s/%s", guildName, realm, region)

	return w.reportsFromUrl(url)
}

// ReportsForUser returns a list of all reports available for provided user.
// Warcraftlogs retention time of reports differs depending on reports tier.
func (w *WarcraftLogs) ReportsForUser(userName string) []*types.Report {
	url := fmt.Sprintf("reports/user/%s", userName)

	return w.reportsFromUrl(url)
}

// Support function to get Reports from URL
func (w *WarcraftLogs) reportsFromUrl(url string) []*types.Report {

	reports := []*types.Report{}
	w.get(url, &reports)

	return reports
}
