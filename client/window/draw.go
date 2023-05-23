package window

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lambher/multiplayer/chore/entity"
	"github.com/lambher/multiplayer/chore/game"
	"image/color"
)

func drawGame(g *game.Game, win *pixelgl.Window) {
	for _, square := range g.Squares {
		drawSquare(square, win)
	}
	for _, apple := range g.Apples {
		drawApple(apple, win)
	}
}

func drawSquare(s *entity.Square, win *pixelgl.Window) {
	s.Imd.Clear()
	s.Imd.Color = s.Color
	s.Imd.Push(pixel.V(s.Position.X-s.Size/2, s.Position.Y-s.Size/2), pixel.V(s.Position.X+s.Size/2, s.Position.Y+s.Size/2))
	s.Imd.Rectangle(0)
	s.Imd.Draw(win)

}

func drawApple(a *entity.Apple, win *pixelgl.Window) {
	a.Imd.Clear()
	a.Imd.Color = color.RGBA{0, 255, 0, 255}
	a.Imd.Push(pixel.V(a.Position.X-a.Size/2, a.Position.Y-a.Size/2), pixel.V(a.Position.X+a.Size/2, a.Position.Y+a.Size/2))
	a.Imd.Rectangle(0)
	a.Imd.Draw(win)

}
