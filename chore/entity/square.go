package entity

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Square struct {
	ID string

	Velocity *pixel.Vec
	Position *pixel.Vec
	Speed    float64
	Friction float64
	Size     float64

	imd *imdraw.IMDraw
}

func NewSquare(id string) Square {
	return Square{
		ID:       id,
		Position: &pixel.Vec{},
		Velocity: &pixel.Vec{},
		Speed:    100,
		Friction: 5,
		Size:     20,

		imd: imdraw.New(nil),
	}
}

func (s *Square) Set(other *Square) {
	s.Position = other.Position
	s.Velocity = other.Velocity
}

func (s *Square) MoveUp() {
	s.Velocity.Y = s.Speed
}

func (s *Square) MoveDown() {
	s.Velocity.Y = -s.Speed
}

func (s *Square) MoveLeft() {
	s.Velocity.X = -s.Speed
}

func (s *Square) MoveRight() {
	s.Velocity.X = s.Speed
}

func (s *Square) Update(deltaTime float64) {
	s.updatePosition(deltaTime)
	s.updateFriction(deltaTime)
}

func (s *Square) updatePosition(deltaTime float64) {
	s.Position.X += s.Velocity.X * deltaTime
	s.Position.Y += s.Velocity.Y * deltaTime
}

func (s *Square) updateFriction(deltaTime float64) {
	if s.Velocity.X > 0 {
		s.Velocity.X -= s.Friction * deltaTime * 100
		if s.Velocity.X < 0 {
			s.Velocity.X = 0
		}
	}
	if s.Velocity.X < 0 {
		s.Velocity.X += s.Friction * deltaTime * 100
		if s.Velocity.X > 0 {
			s.Velocity.X = 0
		}
	}

	if s.Velocity.Y > 0 {
		s.Velocity.Y -= s.Friction * deltaTime * 100
		if s.Velocity.Y < 0 {
			s.Velocity.Y = 0
		}
	}
	if s.Velocity.Y < 0 {
		s.Velocity.Y += s.Friction * deltaTime * 100
		if s.Velocity.Y > 0 {
			s.Velocity.Y = 0
		}
	}
}

func (s *Square) Draw(win *pixelgl.Window) {
	s.imd.Clear()
	s.imd.Color = colornames.Darkgray
	s.imd.Push(pixel.V(s.Position.X-s.Size/2, s.Position.Y-s.Size/2), pixel.V(s.Position.X+s.Size/2, s.Position.Y+s.Size/2))
	s.imd.Rectangle(0)
	s.imd.Draw(win)
}
