package game

import (
	"testing"
)

func TestGetWinner(t *testing.T) {
	game := NewGame()
	winner, _ := game.GetWinner()
	if winner == Red || winner == Blue{
		t.Errorf("There shouldn't be a winner")
	}
}
