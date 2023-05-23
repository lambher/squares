package game

import (
	"github.com/lambher/multiplayer/chore/entity"
)

type Game struct {
	mapSquares map[string]*entity.Square
	Squares    []*entity.Square

	Apples map[string]*entity.Apple
}

func NewGame() *Game {
	return &Game{
		Squares:    make([]*entity.Square, 0),
		mapSquares: make(map[string]*entity.Square),
		Apples:     make(map[string]*entity.Apple),
	}
}

func (g *Game) AddSquare(square *entity.Square) {
	g.Squares = append(g.Squares, square)
	g.mapSquares[square.ID] = square
}

func (g *Game) AddApple(apple *entity.Apple) {
	g.Apples[apple.ID] = apple
}

func (g *Game) GetSquare(id string) *entity.Square {
	return g.mapSquares[id]
}

func (g *Game) GetApple(id string) *entity.Apple {
	return g.Apples[id]
}

func (g *Game) Update(deltaTime float64) {
	for _, square := range g.Squares {
		square.Update(deltaTime)
	}
}
