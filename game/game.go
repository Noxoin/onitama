package game

import (
	"fmt"
	"errors"
	"time"
)

type Game struct {
	board *Board
	turn Team
	cards map[Team]map[string]*Card
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
	cards := make(map[Team]map[string]*Card)
	cards[Red] = make(map[string]*Card)
	cards[Blue] = make(map[string]*Card)
	cards[Red][c[0].Name] = c[0]
	cards[Red][c[1].Name] = c[1]
	cards[Blue][c[2].Name] = c[2]
	cards[Blue][c[3].Name] = c[3]
	neutralCard := c[4]
	return &Game {
		board: board,
		turn: neutralCard.Team,
		cards: cards,
		neutralCard: neutralCard,
	}
}

func (g *Game) validateMove(from Cord, to Cord, c string) (error) {
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
	card := g.cards[g.turn][c]
	if card == nil {
		return errors.New("Unable to perform that move with the cards at hand")
	}
	if !card.isValidMove(move) {
		return errors.New("The specified move is not valid for the card stated")
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

func (g *Game) PerformNextMove(from Cord, to Cord, c string) (error) {
	if err := g.validateMove(from, to, c); err != nil {
		return err
	}
	piece, _ := g.board.GetPiece(from)
	g.board.SetPiece(to, piece)
	g.board.SetPiece(from, nil)
	g.cards[g.turn][g.neutralCard.Name] = g.neutralCard
	g.neutralCard = g.cards[g.turn][c]
	delete(g.cards[g.turn], c)
	return nil
}

func (g *Game) GetWinner() (Team, error) {
	return None, nil
}

