package game

import (
	"time"
)

type Game struct {
	board *Board
	turn Team
	cards map[Team][]*Card
	redCards []*Card
	blueCards []*Card
	neutralCard *Card
}

func NewGame() (*Game) {
	board := NewBoard()
	board.SetPiece(Cord{X: 0, Y: 0}, NewPiece(false, Red))
	board.SetPiece(Cord{X: 1, Y: 0}, NewPiece(false, Red))
	board.SetPiece(Cord{X: 2, Y: 0}, NewPiece(true, Red))
	board.SetPiece(Cord{X: 3, Y: 0}, NewPiece(false, Red))
	board.SetPiece(Cord{X: 4, Y: 0}, NewPiece(false, Red))

	board.SetPiece(Cord{X: 0, Y: 4}, NewPiece(false, Blue))
	board.SetPiece(Cord{X: 1, Y: 4}, NewPiece(false, Blue))
	board.SetPiece(Cord{X: 2, Y: 4}, NewPiece(true, Blue))
	board.SetPiece(Cord{X: 3, Y: 4}, NewPiece(false, Blue))
	board.SetPiece(Cord{X: 4, Y: 4}, NewPiece(false, Blue))

	c := GetRandomCards(5, time.Now().UnixNano())
	cards := make(map[Team][]*Card)
	cards[Red] = c[0:2]
	cards[Blue] = c[2:4]
	neutralCard := c[4]
	return &Game {
		board: board,
		turn: neutralCard.Team,
		cards: cards,
		neutralCard: neutralCard,
	}
}

func (g *Game) PerformNextMove(from Cord, to Cord) (error) {
	// TODO(noxoin): Implement
	return nil
}

