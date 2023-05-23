package entity

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/lambher/multiplayer/chore/game"
	"golang.org/x/image/colornames"
	"image/color"
)

type Square struct {
	ID string

	Velocity *pixel.Vec
	Position *pixel.Vec
	Speed    float64
	Friction float64
	Size     float64
	Color    color.RGBA

	Imd *imdraw.IMDraw
}

func NewSquare(id string) Square {
	return Square{
		ID:       id,
		Position: &pixel.Vec{},
		Velocity: &pixel.Vec{},
		Speed:    100,
		Friction: 5,
		Size:     20,
		Color:    colornames.Darkgray,

		Imd: imdraw.New(nil),
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

	if s.Position.X < 0 {
		s.Position.X = game.ScreenWidth
	}
	if s.Position.X > game.ScreenWidth {
		s.Position.X = 0
	}
	if s.Position.Y < 0 {
		s.Position.Y = game.ScreenHeight
	}
	if s.Position.Y > game.ScreenHeight {
		s.Position.Y = 0
	}
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
