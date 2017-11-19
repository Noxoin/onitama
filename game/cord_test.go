package game

import (
	"fmt"
	"testing"
)

func TestNewCord(t *testing.T) {
	cord := NewCord(3, 5)
	if cord.X != 3 {
		t.Errorf("Cord.X failed: got: %v, want: %v", cord.X, 3)
	}
	if cord.Y != 5 {
		t.Errorf("Cord.Y failed: got: %v, want: %v", cord.Y, 5)
	}
}

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
		t.Run(fmt.Sprintf("%v + %v = %v", test.input, test.delta, test.res), func(t *testing.T) {
			res := test.input.Move(test.delta)
			if res != test.res {
				t.Errorf("Incorrect Resulting Cord: got: %v, want: %v", res, test.res)
			}
		})
	}
}

func TestCordDelta(t *testing.T) {
	tests := []struct {
		input Cord
		end Cord
		res Cord
	} {
		{
			input: Cord{ X: 2, Y: 4 },
			end:   Cord{ X: 3, Y: 3 },
			res:   Cord{ X: 1, Y: -1 },
		},
		{
			input: Cord{ X: 2, Y: 4 },
			end:   Cord{ X: 2, Y: 4 },
			res:   Cord{ X: 0, Y: 0 },
		},
		{
			input: Cord{ X: 2, Y: 4 },
			end:   Cord{ X: 1, Y: 1 },
			res:   Cord{ X: -1, Y: -3 },
		},
		{
			input: Cord{ X: 2, Y: 4 },
			end:   Cord{ X: 4, Y: 4 },
			res:   Cord{ X: 2, Y: 0 },
		},
		{
			input: Cord{ X: 2, Y: 4 },
			end:   Cord{ X: 3, Y: 5 },
			res:   Cord{ X: 1, Y: 1 },
		},
		{
			input: Cord{ X: 2, Y: 4 },
			end:   Cord{ X: 1, Y: 5 },
			res:   Cord{ X: -1, Y: 1 },
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v - %v = %v", test.end, test.input, test.res), func(t *testing.T) {
			res := test.input.Delta(test.end)
			if res != test.res {
				t.Errorf("Incorrect Resulting Cord: got: %v, want: %v", res, test.res)
			}
		})
	}
}
