package messages

import (
	"github.com/lambher/multiplayer/chore/entity"
	"strconv"
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

	return &square, nil
}
