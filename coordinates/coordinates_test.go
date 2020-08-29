package coordinates_test

import (
	co "explorando-marte/coordinates"
	"explorando-marte/plateau"

	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	cmd, cardinal, wantCardinal, name string
	wantX, wantY                      int
	limit                             plateau.Limit
}{
	{
		name:         "when command is L",
		cmd:          "L",
		cardinal:     "E",
		wantX:        0,
		wantY:        0,
		wantCardinal: "N",
	},
	{
		name:         "when command is M",
		cmd:          "M",
		cardinal:     "W",
		wantX:        -1,
		wantY:        0,
		wantCardinal: "W",
	},
	{
		name:         "when command is R",
		cmd:          "R",
		cardinal:     "S",
		wantX:        0,
		wantY:        0,
		wantCardinal: "W",
	},
}

func TestCoordinates(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, y, c := co.Coordinates(tt.cmd, tt.cardinal)

			assert.Equal(t, tt.wantX, x)
			assert.Equal(t, tt.wantY, y)
			assert.Equal(t, tt.wantCardinal, c)
		})
	}
}
