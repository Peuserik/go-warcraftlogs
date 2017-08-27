package warcraftlogs

import (
	"github.com/AlexejK/go-warcraftlogs/types/warcraft"
)

// Classes provides information about all available classes in game and their specializations.
func (w *WarcraftLogs) Classes() []*warcraft.Class {
	classes := []*warcraft.Class{}
	w.get("classes", &classes)

	return classes
}
