package game

import (
	"time"
)

type Game struct {
	board *Board
	turn Team
	redCards []*Card
	blueCards []*Card
	neutralCard *Card
}

func NewGame() (*Game) {
	board := NewBoard()
	board.SetNewPiece(Cord{X: 0, Y: 0}, NewPiece(false, Red))
	board.SetNewPiece(Cord{X: 1, Y: 0}, NewPiece(false, Red))
	board.SetNewPiece(Cord{X: 2, Y: 0}, NewPiece(true, Red))
	board.SetNewPiece(Cord{X: 3, Y: 0}, NewPiece(false, Red))
	board.SetNewPiece(Cord{X: 4, Y: 0}, NewPiece(false, Red))

	board.SetNewPiece(Cord{X: 0, Y: 4}, NewPiece(false, Blue))
	board.SetNewPiece(Cord{X: 1, Y: 4}, NewPiece(false, Blue))
	board.SetNewPiece(Cord{X: 2, Y: 4}, NewPiece(true, Blue))
	board.SetNewPiece(Cord{X: 3, Y: 4}, NewPiece(false, Blue))
	board.SetNewPiece(Cord{X: 4, Y: 4}, NewPiece(false, Blue))

	cards := GetRandomCards(5, time.Now().UnixNano())
	redCards := cards[0:2]
	blueCards := cards[2:4]
	neutralCard := cards[4]
	return &Game {
		board: board,
		turn: neutralCard.Team,
		redCards: redCards,
		blueCards: blueCards,
		neutralCard: neutralCard,
	}
}
