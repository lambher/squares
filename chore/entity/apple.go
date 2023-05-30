package entity

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Apple struct {
	ID       string
	Position *pixel.Vec
	Radius   float64

	Imd *imdraw.IMDraw
}

func NewApple(id string) *Apple {
	return &Apple{
		ID:       id,
		Position: &pixel.Vec{},
		Radius:   5,

		Imd: imdraw.New(nil),
	}
}

func (a *Apple) Set(other *Apple) {
	a.Position = other.Position
	a.Radius = other.Radius
}

func (a *Apple) Collide(square *Square) bool {
	box := pixel.Rect{Min: *square.Position, Max: square.Position.Add(pixel.Vec{X: square.Size, Y: square.Size})}
	return box.Contains(*a.Position)
}
