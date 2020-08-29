package probe

import "fmt"

// Probe contract
type Probe interface {
	GetAxis() (int, int)
	GetCardinal() string
	GetPosition() string
	GetRoute() string
	Move(cardinal string, x, y int)
}

type probe struct {
	position position
	route    string
}

type position struct {
	x        int
	y        int
	cardinal string
}

// NewProbe returns new instance
func NewProbe(cardinal, route string, x, y int) *probe {
	return &probe{
		route: route,
		position: position{
			x:        x,
			y:        y,
			cardinal: cardinal,
		},
	}
}

func (p probe) GetPosition() string {
	return fmt.Sprintf("%v %v %v", p.position.x, p.position.y, p.position.cardinal)
}

func (p probe) GetRoute() string {
	return p.route
}

func (p probe) GetCardinal() string {
	return p.position.cardinal
}

// GetAxis returns x,y
func (p probe) GetAxis() (int, int) {
	return p.position.x, p.position.y
}

func (p *probe) Move(cardinal string, x, y int) {
	p.position.cardinal = cardinal
	p.position.x += x
	p.position.y += y
}
