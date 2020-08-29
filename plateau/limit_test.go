package plateau_test

import (
	"explorando-marte/plateau"
	"explorando-marte/probe"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	var tests = []struct {
		name  string
		x, y  int
		probe probe.Probe
	}{
		{
			x:     5,
			y:     5,
			name:  "when probe 1 receive a valid coordinate",
			probe: probe.NewProbe("N", "LMLMLMLMM", 1, 2),
		},
		{
			x:     5,
			y:     5,
			name:  "when probe 2 receive a valid coordinate",
			probe: probe.NewProbe("E", "MMRMMRMRRM", 3, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := plateau.NewLimit()
			l.SetUpper(tt.x, tt.y)
			err := l.Validate(tt.probe)

			assert.Equal(t, nil, err)
		})
	}
}

func TestValidate_Error(t *testing.T) {
	var tests = []struct {
		name, want string
		x, y       int
		probe      probe.Probe
	}{
		{
			x:     5,
			y:     5,
			name:  "when probe 1 receive a invalid coordinate",
			probe: probe.NewProbe("N", "LMLMLMLMM", 1, 6),
			want:  "outside the exploration area: y-axis",
		},
		{
			x:     5,
			y:     5,
			name:  "when probe 1 receive a invalid coordinate",
			probe: probe.NewProbe("N", "LMLMLMLMM", 15, 3),
			want:  "outside the exploration area: x-axis",
		},
		{
			x:     5,
			y:     5,
			name:  "when probe 2 receive a invalid coordinate",
			probe: probe.NewProbe("E", "MMRMMRMRRM", 3, 6),
			want:  "outside the exploration area: y-axis",
		},
		{
			x:     5,
			y:     5,
			name:  "when probe 2 receive a invalid coordinate",
			probe: probe.NewProbe("E", "MMRMMRMRRM", 1000, 3),
			want:  "outside the exploration area: x-axis",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := plateau.NewLimit()
			l.SetUpper(tt.x, tt.y)
			err := l.Validate(tt.probe)

			assert.Equal(t, tt.want, err.Error())
		})
	}
}
