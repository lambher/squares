package server_conn

import (
	"fmt"
	"github.com/lambher/multiplayer/chore/entity"
	"github.com/lambher/multiplayer/chore/messages"
	"net"
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
	c.conn, err = net.ListenUDP("udp", c.addr)
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

func (c *Connection) Send(message messages.Message, addr *net.UDPAddr) error {
	_, err := c.conn.WriteToUDP([]byte(message), addr)
	if err != nil {
		return fmt.Errorf("cannot send message: %w", err)
	}
	return nil
}

func (c *Connection) SendUpdateSquare(square *entity.Square, addr *net.UDPAddr) error {
	pX, pY, vX, vY := square.Position.X, square.Position.Y, square.Velocity.X, square.Velocity.Y
	_, err := c.conn.WriteToUDP([]byte(fmt.Sprintf("square %s %f %f %f %f", square.ID, pX, pY, vX, vY)), addr)
	if err != nil {
		return fmt.Errorf("cannot send message: %w", err)
	}
	return nil
}

func (c *Connection) Listen() (*net.UDPAddr, string, error) {
	buffer := make([]byte, 1024)

	n, addr, err := c.conn.ReadFromUDP(buffer)
	if err != nil {
		return nil, "", fmt.Errorf("cannot listen: %w", err)
	}
	message := string(buffer[:n])
	fmt.Println("receive", addr.String(), message)

	return addr, message, nil
}