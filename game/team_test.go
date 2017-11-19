package game

import (
	"testing"
)

func TestTeamStringNone(t *testing.T) {
	n := None.String()
	if n != "None" {
		t.Errorf("None failed: got: %v want: %v", n, "None")
	}
}

func TestTeamStringRed(t *testing.T) {
	n := Red.String()
	if n != "Red" {
		t.Errorf("None failed: got: %v want: %v", n, "Red")
	}
}

func TestTeamStringBlue(t *testing.T) {
	n := Blue.String()
	if n != "Blue" {
		t.Errorf("None failed: got: %v want: %v", n, "Blue")
	}
}
