package game

import (
	"testing"
)

func TestGetSetPiece(t *testing.T) {
	board := NewBoard()
	piece, err := board.GetPiece(Cord{X:0, Y:3})
	if err != nil {
		t.Errorf("TestGetSetPiece failed: received err: %v", err)
	}
	if piece != nil {
		t.Errorf("TestGetSetPiece failed: not expecting piece: got: %v", piece)
	}
	p := NewPiece(true, Red)
	err = board.SetPiece(Cord{X:0, Y:2}, p)
	if err != nil {
		t.Errorf("TestGetSetPiece failed: received err: %v", err)
	}
	piece, err = board.GetPiece(Cord{X:0, Y:2})
	if err != nil {
		t.Errorf("TestGetSetPiece failed: received err: %v", err)
	}
	if piece == nil {
		t.Errorf("TestGetSetPiece failed: didn't get a piece; want: %v", *p)
	} else if piece != p {
		t.Errorf("TestGetSetPiece failed: got: %v; want: %v", *piece, *p)
	}
}

func TestGetPieceError(t *testing.T) {
	tests := []struct {
		input Cord
	}{
		{ input: Cord{X: -3, Y:  3} },
		{ input: Cord{X: -3, Y: -3} },
		{ input: Cord{X:  3, Y:  5} },
		{ input: Cord{X:  9, Y:  3} },
	}

	board := NewBoard()
	for _, test := range tests {
		_, err := board.GetPiece(test.input)
		if err == nil {
			t.Errorf("TestGetPieceError failed: expect error on cord %v", test.input)
		}
	}
}

func TestSetPieceError(t *testing.T) {
	tests := []struct {
		input Cord
	}{
		{ input: Cord{X: -3, Y:  3} },
		{ input: Cord{X: -3, Y: -3} },
		{ input: Cord{X:  3, Y:  5} },
		{ input: Cord{X:  9, Y:  3} },
	}

	board := NewBoard()
	piece := NewPiece(true, Red)
	for _, test := range tests {
		err := board.SetPiece(test.input, piece)
		if err == nil {
			t.Errorf("TestSetPieceError failed: expect error on cord %v", test.input)
		}
	}
}

func TestHoldsKing(t *testing.T) {
	tests := []struct {
		piece *Piece
		cord Cord
		resRed bool
		resBlue bool
	} {
		{
			piece: NewPiece(false, Red),
			cord: Cord{X:0, Y:0},
			resRed: false,
			resBlue: false,
		},
		{
			piece: NewPiece(true, Red),
			cord: Cord{X:1, Y:0},
			resRed: true,
			resBlue: false,
		},
		{
			piece: NewPiece(true, Blue),
			cord: Cord{X:3, Y:0},
			resRed: true,
			resBlue: true,
		},
		{
			piece: NewPiece(false, Blue),
			cord: Cord{X:1, Y:0},
			resRed: false,
			resBlue: true,
		},
		{
			piece: NewPiece(false, Red),
			cord: Cord{X:3, Y:0},
			resRed: false,
			resBlue: false,
		},
	}

	board := NewBoard()
	if board.HoldsKing(Red) {
		t.Errorf("TestHoldsKing failed: was not expecting to have a red king")
	}
	if board.HoldsKing(Blue) {
		t.Errorf("TestHoldsKing failed: was not expecting to have a blue king")
	}
	for idx, test := range tests {
		board.SetPiece(test.cord, test.piece)
		if board.HoldsKing(Red) != test.resRed {
			t.Errorf("TestHoldsKing failed: %v: got: %v, want: %v",
			         idx, !test.resRed, test.resRed)
		}
		if board.HoldsKing(Blue) != test.resBlue {
			t.Errorf("TestHoldsKing failed: %v: got: %v, want: %v",
			         idx, !test.resBlue, test.resBlue)
		}
	}
}
