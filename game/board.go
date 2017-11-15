package game

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

func (b *Board) GetPiece(c Cord) (*Piece) {
	return b.board[c.X][c.Y]
}

func (b *Board) SetNewPiece(c Cord, p *Piece) {
	b.board[c.X][c.Y] = p
}

