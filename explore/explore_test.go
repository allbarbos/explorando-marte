package explore_test

import (
	"explorando-marte/explore"
	"explorando-marte/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	route, cardinal, name string
	probes, y, x          int
}{
	{
		name:     "probe 1",
		probes:   2,
		route:    "LMLMLMLMM",
		x:        1,
		y:        3,
		cardinal: "N",
	},
	{
		name:     "probe 2",
		probes:   2,
		route:    "MMRMMRMRRM",
		x:        5,
		y:        1,
		cardinal: "E",
	},
}

func TestInit(t *testing.T) {
	t.Parallel()

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			file := utils.OpenFile("../cmd/input.csv")
			defer file.Close()
			probes := explore.Init(file)

			route := probes[i].GetRoute()
			cardinal := probes[i].GetCardinal()
			x, y := probes[i].GetAxis()

			assert.Equal(t, tt.probes, len(probes))
			assert.Equal(t, tt.route, route)
			assert.Equal(t, tt.x, x)
			assert.Equal(t, tt.y, y)
			assert.Equal(t, tt.cardinal, cardinal)
		})
	}
}
