package game

type Cord struct {
	X int
	Y int
}

func NewCord(x int, y int) (*Cord) {
	return &Cord{
		X: x,
		Y: y,
	}
}

func (c Cord) Move(delta Cord) (*Cord) {
	x := c.X + delta.X
	y := c.Y + delta.Y
	return &Cord{
		X: x,
		Y: y,
	}
}
