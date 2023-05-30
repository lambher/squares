package game

import (
	"github.com/lambher/multiplayer/chore/entity"
)

type Game struct {
	mapSquares map[string]*entity.Square
	Squares    []*entity.Square

	Apples       map[string]*entity.Apple
	eventChannel chan Event
}

func NewGame(eventChannel chan Event) *Game {
	return &Game{
		Squares:      make([]*entity.Square, 0),
		mapSquares:   make(map[string]*entity.Square),
		Apples:       make(map[string]*entity.Apple),
		eventChannel: eventChannel,
	}
}

func (g *Game) AddSquare(square *entity.Square) {
	g.Squares = append(g.Squares, square)
	g.mapSquares[square.ID] = square
}

func (g *Game) AddApple(apple *entity.Apple) {
	g.Apples[apple.ID] = apple
}

func (g *Game) RemoveApple(apple *entity.Apple) {
	delete(g.Apples, apple.ID)
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
		g.checkAppleCollision(square)
	}
}

func (g *Game) checkAppleCollision(square *entity.Square) {
	for _, apple := range g.Apples {
		if apple.Collide(square) {
			g.eventChannel <- Event{
				Type:   AppleCollision,
				Square: square,
				Apple:  apple,
			}
		}
	}
}
