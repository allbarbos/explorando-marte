package plateau

import (
	"errors"
	p "explorando-marte/probe"
)

// Limit contract
type Limit interface {
	Validate(probe p.Probe)
	SetUpper(upperX, upperY int)
}

type limit struct {
	lowerX int
	lowerY int
	upperX int
	upperY int
}

// NewLimit returns new instance
func NewLimit() *limit {
	return &limit{}
}

func (l *limit) SetUpper(upperX, upperY int) {
	l.upperX = upperX
	l.upperY = upperY
}

// Validate plateau boundaries
func (l limit) Validate(probe p.Probe) error {
	x, y := probe.GetAxis()

	if x < l.lowerX || x > l.upperX {
		return errors.New("outside the exploration area: x-axis")
	}

	if y < l.lowerY || y > l.upperY {
		return errors.New("outside the exploration area: y-axis")
	}

	return nil
}
