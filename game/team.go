package game

type Team int

func (t Team) String() string {
	switch t {
	case Red:
		return "Read"
	case Blue:
		return "Blue"
	}
	return ""
}

const (
	Red Team = iota
	Blue
)
