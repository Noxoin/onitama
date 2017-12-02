package server

import (
	"net"
	"strconv"

	o "github.com/noxoin/onitama/game"
)

type Move struct {
	from     o.Cord
	to       o.Cord
	cardName string
}

type GameRoomRequest struct {
	move Move
	player struct {
		name string
		team o.Team
	}
}

type GameRoomResponse struct {
	move     *Move
	opponent string
	err      error
}

type GameRoom struct {
	game  *o.Game
	ready bool
	clients  []*Client
}

func NewGameRoom() *GameRoom {
	game := o.NewGame()
	clients := make([]*Client, 0)
	return &GameRoom{
		game:  game,
		ready: false,
		clients:  clients,
	}
}

func (g *GameRoom) MonitorClient(client *Client) {
	for req := range client.send {
		move := req.move
		if err := g.game.PerformNextMove(move.from, move.to, move.cardName); err != nil {
			client.recv <- GameRoomResponse{
				err: err,
			}
		} else {
			g.Broadcast(move)
		}
	}
}

func (g *GameRoom) Join(conn net.Conn) {
	client := NewClient(conn, strconv.Itoa(len(g.clients)), o.None)
	g.clients = append(g.clients, client)
	go g.MonitorClient(client)
}

func (g *GameRoom) Broadcast(move Move) {
	resp := GameRoomResponse{
		move: &move,
	}
	for _, client := range g.clients {
		client.recv <- resp
	}
}
