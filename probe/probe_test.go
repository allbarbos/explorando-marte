package probe_test

import (
	"explorando-marte/probe"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPosition(t *testing.T) {
	var tests = []struct {
		name, cardinal, route, want string
		x, y                        int
	}{
		{
			x:        1,
			y:        3,
			cardinal: "N",
			want:     "1 3 N",
			route:    "LMLMLMLMM",
			name:     "must return the probe 1 position",
		},
		{
			x:        5,
			y:        1,
			cardinal: "E",
			want:     "5 1 E",
			route:    "LMLMLMLMM",
			name:     "must return the probe 2 position",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := probe.NewProbe(tt.cardinal, tt.route, tt.x, tt.y)

			assert.Equal(t, tt.want, p.GetPosition())
		})
	}
}

func TestGetRoute(t *testing.T) {
	var tests = []struct {
		name, cardinal, route, want string
		x, y                        int
	}{
		{
			x:        1,
			y:        2,
			cardinal: "N",
			want:     "LMLMLMLMM",
			route:    "LMLMLMLMM",
			name:     "must return the probe 1 route",
		},
		{
			x:        3,
			y:        3,
			cardinal: "E",
			want:     "LMLMLMLMM",
			route:    "LMLMLMLMM",
			name:     "must return the probe 2 route",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := probe.NewProbe(tt.cardinal, tt.route, tt.x, tt.y)

			assert.Equal(t, tt.want, p.GetRoute())
		})
	}
}

func TestGetCardinal(t *testing.T) {
	var tests = []struct {
		name, cardinal, route, want string
		x, y                        int
	}{
		{
			x:        1,
			y:        2,
			cardinal: "N",
			want:     "N",
			route:    "LMLMLMLMM",
			name:     "must return the probe 1 cardinal",
		},
		{
			x:        3,
			y:        3,
			cardinal: "E",
			want:     "E",
			route:    "LMLMLMLMM",
			name:     "must return the probe 2 cardinal",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := probe.NewProbe(tt.cardinal, tt.route, tt.x, tt.y)

			assert.Equal(t, tt.want, p.GetCardinal())
		})
	}
}

func TestGetAxis(t *testing.T) {
	var tests = []struct {
		name, cardinal, route string
		x, wantX, y, wantY    int
	}{
		{
			x:        1,
			y:        2,
			cardinal: "N",
			wantX:    1,
			wantY:    2,
			route:    "LMLMLMLMM",
			name:     "must return the probe 1 axis",
		},
		{
			x:        3,
			y:        3,
			cardinal: "E",
			wantX:    3,
			wantY:    3,
			route:    "LMLMLMLMM",
			name:     "must return the probe 2 axis",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := probe.NewProbe(tt.cardinal, tt.route, tt.x, tt.y)

			x, y := p.GetAxis()

			assert.Equal(t, tt.wantX, x)
			assert.Equal(t, tt.wantY, y)
		})
	}
}

func TestMove(t *testing.T) {
	var tests = []struct {
		name, cardinal, mCardinal, route string
		x, mX, y, mY                     int
	}{
		{
			x:         1,
			y:         2,
			cardinal:  "N",
			mX:        3,
			mY:        4,
			mCardinal: "E",
			route:     "LMLMLMLMM",
			name:      "must move the probe 1",
		},
		{
			x:         3,
			y:         3,
			cardinal:  "E",
			mX:        1,
			mY:        2,
			mCardinal: "N",
			route:     "LMLMLMLMM",
			name:      "must move the probe 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := probe.NewProbe(tt.cardinal, tt.route, tt.x, tt.y)

			p.Move(tt.mCardinal, tt.mX, tt.mY)
			x, y := p.GetAxis()
			c := p.GetCardinal()

			assert.Equal(t, (tt.mX + tt.x), x)
			assert.Equal(t, (tt.mY + tt.y), y)
			assert.Equal(t, tt.mCardinal, c)
		})
	}
}
