package server

import (
	"testing"

	o "github.com/noxoin/onitama/game"
)

func TestNewGameRoom(t *testing.T) {
	gameRoom := NewGameRoom()
	if gameRoom.ready {
		t.Errorf("GameRoom should not be ready")
	}
	if gameRoom.clients == nil {
		t.Errorf("recv is not initialized")
	}
	if gameRoom.game == nil {
		t.Errorf("game was not initialized")
	}
}

func TestBroadcast(t *testing.T) {
	count := 3
	c := make(chan GameRoomResponse, count)
	defer close(c)

	gameRoom := NewGameRoom()
	for i := 0; i < count; i = i + 1 {
		gameRoom.clients = append(gameRoom.clients, &Client{
			recv: c,
		})
	}

	move := Move{
		from: o.Cord{X: 2, Y: 3},
		to:   o.Cord{X: 1, Y: 4},
		cardName: "card",
	}

	gameRoom.Broadcast(move)

	for i := 0; i < count;  i = i + 1 {
		m := <-c
		if *m.move != move {
			t.Errorf("Broadcast failed: got: %v, want: %v", m.move, move)
		}
	}
}

func TestMonitorClientError(t *testing.T) {
	c := make(chan GameRoomResponse, 1)
	defer close(c)
	s := make(chan GameRoomRequest, 1)
	defer close(s)
	gameRoom := NewGameRoom()
	client := &Client{
		recv: c,
		send: s,
	}
  gameRoom.clients = append(gameRoom.clients, client)

	s <- GameRoomRequest{
		move: Move{
			from: o.Cord{X: 2, Y: 2},
			to: o.Cord{X: 2, Y: 2},
			cardName: "card",
		},
	}
  go gameRoom.MonitorClient(client)
	resp := <-c
	if resp.err == nil {
		t.Errorf("Expecting error response")
	}
	if resp.move != nil {
		t.Errorf("Should not have move in response")
	}
	if resp.opponent != "" {
		t.Errorf("Should not have an opponent specified")
	}
}

/*
func TestMonitorClientSuccess(t *testing.T) {
	c := make(chan GameRoomResponse, 1)
	defer close(c)
	s := make(chan GameRoomRequest, 1)
	defer close(s)
	gameRoom := NewGameRoom()
	client := &Client{
		recv: c,
		send: s,
	}
  gameRoom.clients = append(gameRoom.clients, client)
  game := o.NewGame()
  c := o.GetRandomCards(5, 0)  // will get crane, eel, monkey, boar, dragon
  game.cards = make(map[o.Team]map[string]*o.Card)
  game.cards[o.Red] = make(map[string]*o.Card)
  game.cards[o.Blue] = make(map[string]*o.Card)
  game.cards[o.Red][c[0].Name] = c[0]
  game.cards[o.Red][c[1].Name] = c[1]
  game.cards[o.Blue][c[2].Name] = c[2]
  game.cards[o.Blue][c[3].Name] = c[3]
  game.neutralCard = c[4]
  game.turn = o.Red
	gameRoom.game = game

	m := Move{
		from: o.Cord{X: 0, Y: 0},
		to: o.Cord{X: 0, Y: 1},
		cardName: "crane",
	}

	s <- GameRoomRequest{
		move: m,
	}
  go gameRoom.MonitorClient(client)
	resp := <-c
	if resp.err != nil {
		t.Errorf("Expecting error response")
	}
	if resp.move != nil {
		t.Errorf("Should not have move in response")
	}
	if *resp.move != m {
		t.Errorf("Move failed: got: %v, want: %v", *resp.move, m)
	}
	if resp.opponent != "" {
		t.Errorf("Should not have an opponent specified")
	}
}
*/
