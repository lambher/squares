package game

import "github.com/lambher/multiplayer/chore/entity"

type EventType string

const (
	AppleCollision EventType = "apple_collision"
)

type Event struct {
	Type   EventType
	Square *entity.Square
	Apple  *entity.Apple
}
