package client_conn

import (
	"errors"
	"fmt"
	"github.com/lambher/multiplayer/chore/entity"
	"github.com/lambher/multiplayer/chore/messages"
	"net"
	"strings"
)

type Connection struct {
	addr *net.UDPAddr
	conn *net.UDPConn
}

func NewConnection(address string) (*Connection, error) {
	var err error
	c := &Connection{}

	c.addr, err = net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve address: %w", err)
	}
	c.conn, err = net.DialUDP("udp", nil, c.addr)
	if err != nil {
		return nil, fmt.Errorf("cannot open connection: %w", err)
	}

	return c, nil
}

func (c *Connection) Close() error {
	err := c.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *Connection) Send(message messages.Message) error {
	_, err := c.conn.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("cannot Send message: %w", err)
	}
	return nil
}

func (c *Connection) Listen() (string, error) {
	buffer := make([]byte, 1024)

	n, err := c.conn.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("cannot Listen: %w", err)
	}

	data := string(buffer[:n])
	fmt.Println("receive", data)
	return data, nil
}

func (c *Connection) AskForConnection() (*entity.Square, error) {
	err := c.Send(messages.AskForConnection)
	if err != nil {
		return nil, err
	}

	data, err := c.Listen()
	if err != nil {
		return nil, err
	}
	infos := strings.Split(data, " ")
	if len(infos) != 6 {
		return nil, errors.New(fmt.Sprintf("not enough infos: %s", data))
	}
	if infos[0] != "square" {
		return nil, errors.New(fmt.Sprintf("unexpected message: %s", data))
	}

	square, err := messages.ParseSquareInfos(infos)
	if err != nil {
		return nil, err
	}

	return square, nil
}
