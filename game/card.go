package game

import (
	"math/rand"
)

var (
	cards = []*Card {
		{
			Name: "boar",
			Team: Red,
			Moves: []*Cord{
				{ X: -1, Y: 0 },
				{ X:  0, Y: 1 },
				{ X:  1, Y: 0 },
			},
		},
		{
			Name: "dragon",
			Team: Red,
			Moves: []*Cord{
				{ X: -2, Y:  1 },
				{ X: -1, Y: -1 },
				{ X:  1, Y: -1 },
				{ X:  2, Y:  1 },
			},
		},
		{
			Name: "eel",
			Team: Blue,
			Moves: []*Cord{
				{ X: -1, Y: -1 },
				{ X: -1, Y:  1 },
				{ X:  1, Y:  0 },
			},
		},
		{
			Name: "monkey",
			Team: Blue,
			Moves: []*Cord{
				{ X: -1, Y: -1 },
				{ X: -1, Y:  1 },
				{ X:  1, Y: -1 },
				{ X:  1, Y:  1 },
			},
		},
		{
			Name: "crane",
			Team: Blue,
			Moves: []*Cord{
				{ X: -1, Y: -1 },
				{ X:  0, Y:  1 },
				{ X:  1, Y: -1 },
			},
		},
		{
			Name: "ox",
			Team: Blue,
			Moves: []*Cord{
				{ X:  0, Y: -1 },
				{ X:  0, Y:  1 },
				{ X:  1, Y:  0 },
			},
		},
		{
			Name: "frog",
			Team: Red,
			Moves: []*Cord{
				{ X: -2, Y:  0 },
				{ X: -1, Y:  1 },
				{ X:  1, Y: -1 },
			},
		},
		{
			Name: "tiger",
			Team: Blue,
			Moves: []*Cord{
				{ X:  0, Y: -1 },
				{ X:  0, Y:  2 },
			},
		},
		{
			Name: "goose",
			Team: Blue,
			Moves: []*Cord{
				{ X: -1, Y:  1 },
				{ X: -1, Y:  0 },
				{ X:  1, Y:  0 },
				{ X:  1, Y: -1 },
			},
		},
		{
			Name: "crab",
			Team: Blue,
			Moves: []*Cord{
				{ X: -2, Y:  0 },
				{ X:  0, Y:  1 },
				{ X:  2, Y:  0 },
			},
		},
		{
			Name: "cobra",
			Team: Red,
			Moves: []*Cord{
				{ X: -1, Y:  0 },
				{ X:  1, Y: -1 },
				{ X:  1, Y:  1 },
			},
		},
		{
			Name: "rooser",
			Team: Red,
			Moves: []*Cord{
				{ X: -1, Y: -1 },
				{ X: -1, Y:  0 },
				{ X:  1, Y:  0 },
				{ X:  1, Y:  1 },
			},
		},
		{
			Name: "mantis",
			Team: Red,
			Moves: []*Cord{
				{ X: -1, Y:  1 },
				{ X:  0, Y: -1 },
				{ X:  1, Y:  1 },
			},
		},
		{
			Name: "rabbit",
			Team: Blue,
			Moves: []*Cord{
				{ X: -1, Y: -1 },
				{ X:  1, Y:  1 },
				{ X:  2, Y:  0 },
			},
		},
		{
			Name: "horse",
			Team: Red,
			Moves: []*Cord{
				{ X: -1, Y:  0 },
				{ X:  0, Y: -1 },
				{ X:  0, Y:  1 },
			},
		},
		{
			Name: "elephant",
			Team: Red,
			Moves: []*Cord{
				{ X: -1, Y:  0 },
				{ X: -1, Y:  1 },
				{ X:  1, Y:  0 },
				{ X:  1, Y:  1 },
			},
		},
	}
)

func GetRandomCards(n int, seed int64) ([]*Card) {
	r := rand.New(rand.NewSource(seed))
	var selected []*Card
	idxs := r.Perm(n)
	for _, idx := range idxs[:5] {
		selected = append(selected, cards[idx])
	}
	return selected
}

type Card struct {
	Name string
	Moves []*Cord
	Team Team
}

func (c *Card) isValidMove(cord Cord) (bool) {
	for _, delta := range c.Moves {
		if *delta == cord {
			return true
		}
	}
	return false
}

