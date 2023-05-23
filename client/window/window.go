package window

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lambher/multiplayer/chore/config"
	"github.com/lambher/multiplayer/chore/entity"
	"github.com/lambher/multiplayer/chore/game"
	"github.com/lambher/multiplayer/chore/messages"
	"golang.org/x/image/colornames"
	"time"
)

type KeyType string

const (
	Up    KeyType = "up"
	Down  KeyType = "down"
	Left  KeyType = "left"
	Right KeyType = "right"

	UpBis    KeyType = "up_bis"
	DownBis  KeyType = "down_bis"
	LeftBis  KeyType = "left_bis"
	RightBis KeyType = "right_bis"
)

var keyMap = map[KeyType]pixelgl.Button{
	Up:    pixelgl.KeyW,
	Down:  pixelgl.KeyS,
	Left:  pixelgl.KeyA,
	Right: pixelgl.KeyD,

	UpBis:    pixelgl.KeyUp,
	DownBis:  pixelgl.KeyDown,
	LeftBis:  pixelgl.KeyLeft,
	RightBis: pixelgl.KeyRight,
}

func Start(s *entity.Square, g *game.Game, c chan messages.Message) {
	pixelgl.Run(func() {
		cfg := pixelgl.WindowConfig{
			Title:  "Red Square Game",
			Bounds: pixel.R(0, 0, config.ScreenWidth, config.ScreenHeight),
		}
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		var previousTime = time.Now()

		for !win.Closed() {
			currentTime := time.Now()
			deltaTime := currentTime.Sub(previousTime)
			previousTime = currentTime

			listenEvent(s, win, c)

			win.Clear(colornames.Black)

			drawGame(g, win)
			g.Update(deltaTime.Seconds())
			win.Update()

			time.Sleep(time.Millisecond * 16) // 60 FPS
		}
	})
}
