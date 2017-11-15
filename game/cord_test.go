package game

import (
	"testing"
)

func TestCordMove(t *testing.T) {
	tests := []struct {
		input Cord
		delta Cord
		res Cord
	} {
		{
			input: Cord{ X: 2, Y: 4 },
			delta: Cord{ X: 0, Y: 1 },
			res:   Cord{ X: 2, Y: 5 },
		},
		{
			input: Cord{ X: 2, Y: 4 },
			delta: Cord{ X: 3, Y: 1 },
			res:   Cord{ X: 5, Y: 5 },
		},
		{
			input: Cord{ X: 2, Y: 4 },
			delta: Cord{ X: -3, Y: -1 },
			res:   Cord{ X: -1, Y: 3 },
		},
		{
			input: Cord{ X: 2, Y: 4 },
			delta: Cord{ X: 2, Y: 4 },
			res:   Cord{ X: 4, Y: 8 },
		},
		{
			input: Cord{ X: 2, Y: 4 },
			delta: Cord{ X: -2, Y: -4 },
			res:   Cord{ X: 0, Y: 0 },
		},
		{
			input: Cord{ X: 2, Y: 4 },
			delta: Cord{ X: -1, Y: 3 },
			res:   Cord{ X: 1, Y: 7 },
		},
		{
			input: Cord{ X: 2, Y: 4 },
			delta: Cord{ X: 3, Y: -9 },
			res:   Cord{ X: 5, Y: -5 },
		},
	}

	for _, test := range tests {
		res := test.input.Move(test.delta)
		if *res != test.res {
			t.Errorf("Incorrect Resulting Cord: got: %v, want: %v", *res, test.res)
		}
	}
}
