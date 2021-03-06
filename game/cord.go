package game

type Cord struct {
	X int
	Y int
}

func NewCord(x int, y int) (Cord) {
	return Cord{
		X: x,
		Y: y,
	}
}

func (c Cord) delta(end Cord) (Cord) {
	x := end.X - c.X
	y := end.Y - c.Y
	return Cord{
		X: x,
		Y: y,
	}
}

func (c Cord) move(delta Cord) (Cord) {
	x := c.X + delta.X
	y := c.Y + delta.Y
	return Cord{
		X: x,
		Y: y,
	}
}
