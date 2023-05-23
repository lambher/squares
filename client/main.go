package main

import (
	"errors"
	"fmt"
	"github.com/lambher/multiplayer/chore/entity"
	"github.com/lambher/multiplayer/chore/game"
	"github.com/lambher/multiplayer/chore/messages"
	"github.com/lambher/multiplayer/client/client_conn"
	"github.com/lambher/multiplayer/client/window"
	"strings"
)

func main() {
	//window.Start()
	g := game.NewGame()

	conn, err := client_conn.NewConnection("aerotoulousain.fr:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	square, err := conn.AskForConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	g.AddSquare(square)

	go func() {
		for {
			message, err := conn.Listen()
			if err != nil {
				fmt.Println("cannot listen", err)
				continue
			}
			err = handleMessage(square, g, message)
			if err != nil {
				fmt.Println("cannot handle message", err)
				continue
			}
		}
	}()

	c := make(chan messages.Message)
	listenMessageChannel(conn, c)
	window.Start(square, g, c)

	defer conn.Close()
}

func handleMessage(square *entity.Square, g *game.Game, message string) error {
	infos := strings.Split(message, " ")
	if len(infos) != 7 {
		return errors.New(fmt.Sprintf("not enough infos: %s", message))
	}
	if infos[0] == "square" {
		s, err := messages.ParseSquareInfos(infos)
		if err != nil {
			return err
		}
		square := g.GetSquare(s.ID)
		if square == nil {
			g.AddSquare(s)
			return nil
		}
		square.Set(s)
	}

	return nil
}

func listenMessageChannel(conn *client_conn.Connection, c chan messages.Message) {
	go func() {
		for message := range c {
			err := conn.Send(message)
			if err != nil {
				fmt.Println("cannot send message", err)
				continue
			}
		}
	}()

}
