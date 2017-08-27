package warcraftlogs

import "github.com/AlexejK/go-warcraftlogs/types"

// Zones provides information about all available zones and their encounters.
func (w *WarcraftLogs) Zones() []*types.Zone {
	zones := []*types.Zone{}
	w.get("zones", &zones)

	return zones
}
