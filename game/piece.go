package game

type Piece struct {
	Position Cord
	Team Team
	king bool
}

func NewPiece(position Cord, king bool, team Team) (*Piece) {
	return &Piece{
		Position: position,
		king: king,
		Team: team,
	}
}

func (p *Piece) isKing() (bool) {
	return p.king;
}
