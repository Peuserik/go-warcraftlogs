package warcraftlogs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWarcraftLogs_Zones(t *testing.T) {

	ts, api := testServerAndClient(t, "zones.json")
	defer ts.Close()

	zones := api.Zones()

	assert.Equal(t, 3, len(zones))
	assert.Equal(t, int64(13), *zones[1].Id)

	/*
		for _, z := range zones {
			fmt.Printf("Zone (%d) %s\n", *z.Id, *z.Name)
			fmt.Println(" -- Encounters --")
			for _, e := range z.Encounters {
				fmt.Printf(" Encounter (%d) %s\n", *e.Id, *e.Name)
			}

			fmt.Println(" -- Brackets --")
			for _, b := range z.Brackets {
				fmt.Printf(" Bracket (%d) %s\n", *b.Id, *b.Name)
			}
		}*/

}
