package server

import (
	"bufio"
	"net"

	o "github.com/noxoin/onitama/game"
)

type Client struct {
	id     string
	team   o.Team
	recv   chan GameRoomResponse // recv-ing from the GameRoom
	send   chan GameRoomRequest  // send-ing to the GameRoom
	reader *bufio.Reader         // reading from socket
	writer *bufio.Writer         // writer to socket
}

func NewClient(conn net.Conn, id string, team o.Team) *Client {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	client := &Client{
		id:     id,
		team:   team,
		recv:   make(chan GameRoomResponse),
		send:   make(chan GameRoomRequest),
		reader: reader,
		writer: writer,
	}

	return client
}

func (c *Client) Listen() {
	go c.Read()
	go c.Write()
}

func (c *Client) Read() {
	for {
		c.reader.ReadString('\n')
		// TODO(noxoin): Handle the move somehow
	}
}

func (c *Client) Write() {
	for response := range c.recv {
		if response.err != nil {
			c.writer.WriteString(response.err.Error())
		} else {
			c.writer.WriteString(response.move.cardName)
		}
		c.writer.Flush()
	}
}
