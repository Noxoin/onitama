package game

type Board struct {
	turn int
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

func (b *Board) getPiece(c Cord) (*Piece) {
	return b.board[c.X][c.Y]
}

