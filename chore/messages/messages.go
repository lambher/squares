package messages

import (
	"github.com/lambher/multiplayer/chore/entity"
	"golang.org/x/image/colornames"
	"image/color"
	"strconv"
	"strings"
)

type Message string

const (
	AskForConnection Message = "1"
	Up               Message = "up"
	Down             Message = "down"
	Left             Message = "left"
	Right            Message = "right"
)

const (
	ConnectionAccepted Message = "2"
)

func getColor(data string) color.RGBA {
	c := colornames.Darkgray

	colors := strings.Split(data, ",")
	if len(colors) != 4 {
		return c
	}
	r, err := strconv.Atoi(colors[0])
	if err != nil {
		return c
	}
	g, err := strconv.Atoi(colors[1])
	if err != nil {
		return c
	}
	b, err := strconv.Atoi(colors[2])
	if err != nil {
		return c
	}
	a, err := strconv.Atoi(colors[3])
	if err != nil {
		return c
	}

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}
}

func ParseSquareInfos(infos []string) (*entity.Square, error) {
	pX, err := strconv.ParseFloat(infos[2], 64)
	if err != nil {
		return nil, err
	}
	pY, err := strconv.ParseFloat(infos[3], 64)
	if err != nil {
		return nil, err
	}
	vX, err := strconv.ParseFloat(infos[4], 64)
	if err != nil {
		return nil, err
	}
	vY, err := strconv.ParseFloat(infos[5], 64)
	if err != nil {
		return nil, err
	}
	square := entity.NewSquare(infos[1])
	square.Position.X = pX
	square.Position.Y = pY
	square.Velocity.X = vX
	square.Velocity.Y = vY

	square.Color = getColor(infos[6])

	return &square, nil
}
