package game

type Team int

func (t Team) String() string {
	switch t {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	}
	return "None"
}

const (
	None Team = iota
	Red
	Blue
)
