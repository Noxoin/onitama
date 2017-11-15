package game

type Piece struct {
	Position Cord
	Team int
	king bool
}

func NewPiece(position Cord, king bool, team int) (*Piece) {
	return &Piece{
		Position: position,
		king: king,
		Team: team,
	}
}

func (p *Piece) isKing() (bool) {
	return p.king;
}
