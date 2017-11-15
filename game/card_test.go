package game

import (
	"testing"
)

func TestGenerateRandom(t *testing.T) {
	cards := GetRandomCards(5, 1)
	if len(cards) != 5 {
		t.Errorf("Incorrect resulting number of cards: want:5 got: %v", len(cards))
		return
	}
	seen := make(map[string]bool)
	for _, v := range cards {
		if seen[v.Name] {
			t.Errorf("There are duplicates within the resulting cards")
			return
		}
		seen[v.Name] = true
	}
}

func TestIsValidMove(t *testing.T) {
	card := &Card {
		Moves: []*Cord{
			{ X: -1, Y: 0 },
			{ X:  0, Y: 1 },
			{ X:  1, Y: 0 },
		},
	}
	tests := []struct {
		input Cord
		res bool
	} {
		{ input: Cord{ X: -1, Y: 0 }, res: true },
		{ input: Cord{ X: 0, Y: 1 }, res: true },
		{ input: Cord{ X: 1, Y: 0 }, res: true },
		{ input: Cord{ X: 1, Y: 1 }, res: false },
		{ input: Cord{ X: 0, Y: 0 }, res: false },
	}

	for _, v := range tests {
		res := card.isValidMove(v.input)
		if res != v.res {
			t.Errorf("TestIsValidMove failed: got: %v, want: %v", res, v.res)
		}
	}
}
