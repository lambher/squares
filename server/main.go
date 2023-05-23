package main

import (
	"fmt"
	"github.com/lambher/multiplayer/chore/entity"
	"github.com/lambher/multiplayer/chore/game"
	"github.com/lambher/multiplayer/chore/messages"
	"github.com/lambher/multiplayer/server/server_conn"
	"image/color"
	"math/rand"
	"net"
	"time"
)

type Client struct {
	addr   *net.UDPAddr
	square *entity.Square
}

func NewClient(addr *net.UDPAddr) *Client {
	return &Client{
		addr: addr,
	}
}

func (c *Client) SetSquare(square *entity.Square) {
	c.square = square
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := game.NewGame()
	clients := make(map[string]*Client)

	// créer une adresse pour écouter sur un port spécifique
	conn, err := server_conn.NewConnection(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	go func() {
		fmt.Println("Listening...")
		for {
			addr, message, err := conn.Listen()
			if err != nil {
				fmt.Println(err)
				continue
			}
			client, ok := clients[addr.String()]
			if !ok {
				fmt.Println(addr.String(), "new client")
				client = NewClient(addr)
				clients[addr.String()] = client
			}
			err = handleMessage(g, conn, client, message)
			if err != nil {
				fmt.Println("cannot handle message", err)
				continue
			}
		}
	}()

	var previousTime = time.Now()

	for {
		currentTime := time.Now()
		deltaTime := currentTime.Sub(previousTime)
		previousTime = currentTime

		g.Update(deltaTime.Seconds())
		broadCastGameState(g, clients, conn)
		time.Sleep(time.Millisecond * 16) // 60 FPS
	}
}

func broadCastGameState(g *game.Game, clients map[string]*Client, conn *server_conn.Connection) {
	for _, square := range g.Squares {
		for _, client := range clients {
			err := conn.SendUpdateSquare(square, client.addr)
			if err != nil {
				fmt.Println("cannot send update square", square.ID, client.addr.String(), err)
				continue
			}
		}
	}
}

func randomColor() color.RGBA {
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
}

func newSquare(client *Client) *entity.Square {
	square := entity.NewSquare(client.addr.String())
	square.Color = randomColor()
	square.Position.X = game.ScreenWidth / 2
	square.Position.Y = game.ScreenHeight / 2
	return &square
}

func handleMessage(g *game.Game, conn *server_conn.Connection, client *Client, message string) error {
	if messages.Message(message) == messages.AskForConnection {
		square := newSquare(client)
		client.SetSquare(square)
		g.AddSquare(square)

		err := conn.SendUpdateSquare(square, client.addr)
		if err != nil {
			return err
		}
		return nil
	}
	if messages.Message(message) == messages.Up {
		client.square.MoveUp()
	}
	if messages.Message(message) == messages.Down {
		client.square.MoveDown()
	}
	if messages.Message(message) == messages.Left {
		client.square.MoveLeft()
	}
	if messages.Message(message) == messages.Right {
		client.square.MoveRight()
	}

	return nil
}
