package warcraftlogs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWarcraftLogs_Classes(t *testing.T) {

	ts, api := testServerAndClient(t, "classes.json")
	defer ts.Close()

	classes := api.Classes()

	// Validations

	assert.Equal(t, 12, len(classes), "Unexpected number of classes returned")

	druidFound := false
	for _, class := range classes {
		if *class.Id == 2 {
			druidFound = true

			assert.Equal(t, "Druid", *class.Name, "Parsing of class name")

		}
	}

	assert.True(t, druidFound, "Druid class not found in response")

}
