package game

import (
	"fmt"
	"errors"
)

type Board struct {
	board [][]*Piece
}

func NewBoard() (*Board) {
	board := make([][]*Piece, 5)
	for i := range board {
		board[i] = make([]*Piece, 5)
	}
	return &Board {
		board: board,

	}
}

func (b *Board) getPiece(c Cord) (*Piece, error) {
	if c.X < 0 || c.Y < 0 || c.X >= 5 || c.Y >= 5 {
		return nil, errors.New(fmt.Sprintf("Invalid Cord: %v", c))
	}
	return b.board[c.X][c.Y], nil
}

func (b *Board) setPiece(c Cord, p *Piece) (error) {
	if c.X < 0 || c.Y < 0 || c.X >= 5 || c.Y >= 5 {
		return errors.New(fmt.Sprintf("Invalid Cord: %v", c))
	}
	b.board[c.X][c.Y] = p
	return nil
}

func (b *Board) holdsKing(t Team) (bool) {
	for _, row := range b.board {
		for _, val := range row {
			if val != nil && val.isKing() && val.Team == t {
				return true
			}
		}
	}
	return false
}

