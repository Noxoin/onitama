package game

import (
	"fmt"
	"errors"
	"time"
)

type Game struct {
	board *Board
	turn Team
	cards map[Team][]*Card
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

func (g *Game) getCardFromMove(move Cord) (*Card) {
	for _, card := range g.cards[g.turn] {
		if card.isValidMove(move) {
			return card
		}
	}
	return nil
}

func (g *Game) validateMove(from Cord, to Cord) (error) {
	startPiece, err := g.board.GetPiece(from)
	if err != nil {
		return err
	}
	if startPiece == nil {
		return errors.New(fmt.Sprintf("There is no piece on starting Cord: %v", from))
	}

	if g.turn != startPiece.Team {
		return errors.New("Trying to move a piece from the other team")
	}

	move := from.Delta(to)
	card := g.getCardFromMove(move)
	if card == nil {
		return errors.New("Unable to perform that move with the cards at hand")
	}

	endPiece, err := g.board.GetPiece(to)
	if err != nil {
		return err
	}
	if endPiece != nil && startPiece.Team == endPiece.Team {
		return errors.New(fmt.Sprintf("Ending Cord %v is already occupied by a piece on the same team", to))
	}

	return nil
}

func (g *Game) PerformNextMove(from Cord, to Cord) (error) {
	if err := g.validateMove(from, to); err != nil {
		return err
	}
	piece, _ := g.board.GetPiece(from)
	g.board.SetPiece(to, piece)
	g.board.SetPiece(from, nil)
	return nil
}

func (g *Game) GetWinner() (Team, error) {
	return None, nil
}

