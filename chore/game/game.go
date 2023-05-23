package game

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/lambher/multiplayer/chore/entity"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type Game struct {
	mapSquares map[string]*entity.Square
	Squares    []*entity.Square
}

func NewGame() *Game {
	return &Game{
		Squares:    make([]*entity.Square, 0),
		mapSquares: make(map[string]*entity.Square),
	}
}

func (g *Game) AddSquare(square *entity.Square) {
	g.Squares = append(g.Squares, square)
	g.mapSquares[square.ID] = square
}

func (g *Game) GetSquare(id string) *entity.Square {
	return g.mapSquares[id]
}

func (g *Game) Update(deltaTime float64) {
	for _, square := range g.Squares {
		square.Update(deltaTime)
	}
}

func (g *Game) Draw(win *pixelgl.Window) {
	for _, square := range g.Squares {
		square.Draw(win)
	}
}
