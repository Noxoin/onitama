package game

type Piece struct {
	Team Team
	king bool
}

func NewPiece(king bool, team Team) (*Piece) {
	return &Piece{
		king: king,
		Team: team,
	}
}

func (p *Piece) isKing() (bool) {
	return p.king;
}
