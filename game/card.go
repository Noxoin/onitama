package game

import (
	"math/rand"
)

var (
	cards = []*Card {
		{
			Name: "boar",
			Team: 0,
			Moves: []*Cord{
				{ X: -1, Y: 0 },
				{ X:  0, Y: 1 },
				{ X:  1, Y: 0 },
			},
		},
		{
			Name: "dragon",
			Team: 0,
			Moves: []*Cord{
				{ X: -2, Y:  1 },
				{ X: -1, Y: -1 },
				{ X:  1, Y: -1 },
				{ X:  2, Y:  1 },
			},
		},
		{
			Name: "eel",
			Team: 1,
			Moves: []*Cord{
				{ X: -1, Y: -1 },
				{ X: -1, Y:  1 },
				{ X:  1, Y:  0 },
			},
		},
		{
			Name: "monkey",
			Team: 1,
			Moves: []*Cord{
				{ X: -1, Y: -1 },
				{ X: -1, Y:  1 },
				{ X:  1, Y: -1 },
				{ X:  1, Y:  1 },
			},
		},
		{
			Name: "crane",
			Team: 1,
			Moves: []*Cord{
				{ X: -1, Y: -1 },
				{ X:  0, Y:  1 },
				{ X:  1, Y: -1 },
			},
		},
		{
			Name: "ox",
			Team: 1,
			Moves: []*Cord{
				{ X:  0, Y: -1 },
				{ X:  0, Y:  1 },
				{ X:  1, Y:  0 },
			},
		},
		{
			Name: "frog",
			Team: 0,
			Moves: []*Cord{
				{ X: -2, Y:  0 },
				{ X: -1, Y:  1 },
				{ X:  1, Y: -1 },
			},
		},
		{
			Name: "tiger",
			Team: 1,
			Moves: []*Cord{
				{ X:  0, Y: -1 },
				{ X:  0, Y:  2 },
			},
		},
		{
			Name: "goose",
			Team: 1,
			Moves: []*Cord{
				{ X: -1, Y:  1 },
				{ X: -1, Y:  0 },
				{ X:  1, Y:  0 },
				{ X:  1, Y: -1 },
			},
		},
		{
			Name: "crab",
			Team: 1,
			Moves: []*Cord{
				{ X: -2, Y:  0 },
				{ X:  0, Y:  1 },
				{ X:  2, Y:  0 },
			},
		},
		{
			Name: "cobra",
			Team: 0,
			Moves: []*Cord{
				{ X: -1, Y:  0 },
				{ X:  1, Y: -1 },
				{ X:  1, Y:  1 },
			},
		},
		{
			Name: "rooser",
			Team: 0,
			Moves: []*Cord{
				{ X: -1, Y: -1 },
				{ X: -1, Y:  0 },
				{ X:  1, Y:  0 },
				{ X:  1, Y:  1 },
			},
		},
		{
			Name: "mantis",
			Team: 0,
			Moves: []*Cord{
				{ X: -1, Y:  1 },
				{ X:  0, Y: -1 },
				{ X:  1, Y:  1 },
			},
		},
		{
			Name: "rabbit",
			Team: 1,
			Moves: []*Cord{
				{ X: -1, Y: -1 },
				{ X:  1, Y:  1 },
				{ X:  2, Y:  0 },
			},
		},
		{
			Name: "horse",
			Team: 0,
			Moves: []*Cord{
				{ X: -1, Y:  0 },
				{ X:  0, Y: -1 },
				{ X:  0, Y:  1 },
			},
		},
		{
			Name: "elephant",
			Team: 0,
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
	Team int  // Red = 0, Blue = 1
}

func (c *Card) isValidMove(cord Cord) (bool) {
	for _, delta := range c.Moves {
		if *delta == cord {
			return true
		}
	}
	return false
}

