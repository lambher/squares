package window

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/lambher/multiplayer/chore/entity"
	"github.com/lambher/multiplayer/chore/messages"
)

func listenEvent(s *entity.Square, win *pixelgl.Window, c chan messages.Message) {
	if win.Pressed(keyMap[Up]) {
		s.MoveUp()
		c <- messages.Up
	}
	if win.Pressed(keyMap[Down]) {
		s.MoveDown()
		c <- messages.Down
	}
	if win.Pressed(keyMap[Left]) {
		s.MoveLeft()
		c <- messages.Left
	}
	if win.Pressed(keyMap[Right]) {
		s.MoveRight()
		c <- messages.Right
	}
}
