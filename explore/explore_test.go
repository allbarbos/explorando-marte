package explore_test

import (
	"explorando-marte/explore"
	"explorando-marte/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
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

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := utils.OpenFile("./testdata/input.csv")
			defer file.Close()
			probes, _ := explore.Init(file)

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

func TestInit_InvalidAxis(t *testing.T) {
	var tests = []struct {
		name, file, want string
	}{
		{
			name: "when x-axis is invalid",
			file: "./testdata/invalid-axis-x.csv",
			want: `strconv.Atoi: parsing "a": invalid syntax`,
		},
		{
			name: "when y-axis is invalid",
			file: "./testdata/invalid-axis-y.csv",
			want: `strconv.Atoi: parsing "a": invalid syntax`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := utils.OpenFile(tt.file)
			defer file.Close()
			_, err := explore.Init(file)

			assert.Equal(t, tt.want, err.Error())
		})
	}
}

func TestInit_InvalidLimits(t *testing.T) {
	var tests = []struct {
		name, file, want string
	}{
		{
			name: "when x limit is invalid",
			file: "./testdata/invalid-limit-x.csv",
			want: `strconv.Atoi: parsing "a": invalid syntax`,
		},
		{
			name: "when y limit is invalid",
			file: "./testdata/invalid-limit-y.csv",
			want: `strconv.Atoi: parsing "a": invalid syntax`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := utils.OpenFile(tt.file)
			defer file.Close()

			_, err := explore.Init(file)

			assert.Equal(t, tt.want, err.Error())
		})
	}
}

func TestInit_ValidateLimitError(t *testing.T) {
	t.Run("when the probe goes out of bounds", func(t *testing.T) {
		file := utils.OpenFile("./testdata/invalid-plateau-limit.csv")
		defer file.Close()

		_, err := explore.Init(file)

		assert.NotNil(t, err)
		assert.Equal(t, "outside the exploration area: x-axis", err.Error())
	})
}
