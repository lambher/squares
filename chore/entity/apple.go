package entity

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Apple struct {
	ID       string
	Position *pixel.Vec
	Size     float64

	Imd *imdraw.IMDraw
}

func NewApple(id string) *Apple {
	return &Apple{
		ID:       id,
		Position: &pixel.Vec{},
		Size:     20,

		Imd: imdraw.New(nil),
	}
}

func (a *Apple) Set(other *Apple) {
	a.Position = other.Position
	a.Size = other.Size
}
