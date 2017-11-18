package game

type Team int

func (t Team) String() string {
	switch t {
	case Red:
		return "Read"
	case Blue:
		return "Blue"
	case None:
		return "None"
	}
	return ""
}

const (
	None Team = iota
	Red
	Blue
)
