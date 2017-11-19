package game

import (
	"testing"
)

func getTestGame() (*Game) {
	game := NewGame()
  c := GetRandomCards(5, 0)  // will get crane, eel, monkey, boar, dragon
  game.cards = make(map[Team]map[string]*Card)
  game.cards[Red] = make(map[string]*Card)
  game.cards[Blue] = make(map[string]*Card)
  game.cards[Red][c[0].Name] = c[0]
  game.cards[Red][c[1].Name] = c[1]
  game.cards[Blue][c[2].Name] = c[2]
  game.cards[Blue][c[3].Name] = c[3]
  game.neutralCard = c[4]
	game.turn = Red
	return game
}

func TestValidateMove(t *testing.T) {
	t.Run("Valid Move", func(t *testing.T) {
		game := getTestGame()
		err := game.validateMove(Cord{1, 0}, Cord{1, 1}, "crane")
		if err != nil {
			t.Errorf("Should not be an error")
		}
	})

	t.Run("Starting Cord Out of Bounds", func(t *testing.T) {
		game := getTestGame()
		err := game.validateMove(Cord{-1, 1}, Cord{0, 0}, "card")
		if err == nil {
			t.Errorf("Should be an error")
		}
	})

	t.Run("Missing Starting Piece", func(t *testing.T) {
		game := getTestGame()
		err := game.validateMove(Cord{1, 1}, Cord{0, 0}, "card")
		if err == nil {
			t.Errorf("Should be an error")
		}
	})

	t.Run("Trying to Move Opposing Team's piece", func(t *testing.T) {
		game := getTestGame()
		err := game.validateMove(Cord{1, 4}, Cord{0, 0}, "card")
		if err == nil {
			t.Errorf("Should be an error")
		}
	})

	t.Run("Trying to use incorrect card", func(t *testing.T) {
		game := getTestGame()
		err := game.validateMove(Cord{1, 0}, Cord{1, 1}, "monkey")
		if err == nil {
			t.Errorf("Should be an error")
		}
	})

	t.Run("Trying invalid move with card", func(t *testing.T) {
		game := getTestGame()
		err := game.validateMove(Cord{1, 0}, Cord{1, 4}, "crane")
		if err == nil {
			t.Errorf("Should be an error")
		}
	})

	t.Run("Trying to move off board", func(t *testing.T) {
		game := getTestGame()
		err := game.validateMove(Cord{1, 0}, Cord{0, -1}, "crane")
		if err == nil {
			t.Errorf("Should be an error")
		}
	})

	t.Run("Trying to move onto ally occupied space", func(t *testing.T) {
		game := getTestGame()
		err := game.validateMove(Cord{1, 0}, Cord{2, 0}, "eel")
		if err == nil {
			t.Errorf("Should be an error")
		}
	})

	t.Run("Take Oppoents Piece", func(t *testing.T) {
		game := getTestGame()
		err := game.board.setPiece(Cord{X:1, Y:1}, NewPiece(false, Blue))
		if err != nil {
			t.Fatalf("Setup Failed: %v", err)
		}
		err = game.validateMove(Cord{1, 0}, Cord{1, 1}, "crane")
		if err != nil {
			t.Errorf("Should not be an error")
		}
	})
}

func TestPerformNextMove(t *testing.T) {
	t.Run("Failed Validation", func(t *testing.T) {
		game := getTestGame()
		err := game.PerformNextMove(Cord{X:0, Y:0}, Cord{X:0, Y:0}, "crane")
		if err == nil {
			t.Errorf("Should be an error")
		}
	})

	t.Run("Successful Move", func(t *testing.T) {
		game := getTestGame()
		pieceToMove, _ := game.board.getPiece(Cord{X:0, Y:0})
		err := game.PerformNextMove(Cord{X:0, Y:0}, Cord{X:0, Y:1}, "crane")
		if err != nil {
			t.Errorf("Should not be an error")
		}
		if game.neutralCard.Name != "crane" {
			t.Errorf("Neutral Card failed: got: %v, want: %v", game.neutralCard.Name, "crane")
		}
		if game.cards[Red]["crane"] != nil {
			t.Errorf("Crane card should not still be in Red Hand")
		}
		piece, _ := game.board.getPiece(Cord{X:0, Y:0})
		if piece != nil {
			t.Errorf("Piece should be no longer exist at {0, 0}")
		}
		piece, _ = game.board.getPiece(Cord{X:0, Y:1})
		if piece != pieceToMove {
			t.Errorf("Piece @ {0, 1} should be the piece that was at {0, 0}")
		}
	})
}

func TestGetWinner(t *testing.T) {
	t.Run("New Board", func(t *testing.T) {
		game := NewGame()
		winner, err := game.GetWinner()
		if err != nil {
			t.Errorf("There shouldn't be an error")
		}
		if winner != None {
			t.Errorf("There shouldn't be a winner")
		}
	})

	t.Run("Temple Conditions", func(t *testing.T) {
		t.Run("Blue king @ Red temple", func(t *testing.T) {
			game := NewGame()
			game.board.setPiece(redTemple, NewPiece(true, Blue))
			game.board.setPiece(Cord{X:0, Y:4}, NewPiece(true, Red))
			winner, err := game.GetWinner()
			if err != nil {
				t.Errorf("There shouldn't be an error")
			}
			if winner != Blue {
				t.Errorf("There should be a winner: got: %v, want: %v", winner, Blue)
			}
		})

		t.Run("Red king @ Blue temple", func(t *testing.T) {
			game := NewGame()
			game.board.setPiece(blueTemple, NewPiece(true, Red))
			game.board.setPiece(Cord{X:0, Y:0}, NewPiece(true, Blue))
			winner, err := game.GetWinner()
			if err != nil {
				t.Errorf("There shouldn't be an error")
			}
			if winner != Red {
				t.Errorf("There should be a winner: got: %v, want: %v", winner, Red)
			}
		})

		t.Run("Red king @ Blue temple and Blue king @ Red Temple", func(t *testing.T) {
			game := NewGame()
			game.board.setPiece(blueTemple, NewPiece(true, Red))
			game.board.setPiece(redTemple, NewPiece(true, Blue))
			_, err := game.GetWinner()
			if err == nil {
				t.Errorf("There should be an error")
			}
		})

		t.Run("Kings not on any temple", func(t *testing.T) {
			game := NewGame()
			game.board.setPiece(blueTemple, NewPiece(false, Red))
			game.board.setPiece(redTemple, NewPiece(false, Red))
			game.board.setPiece(Cord{X:0, Y:0}, NewPiece(true, Red))
			game.board.setPiece(Cord{X:0, Y:4}, NewPiece(true, Blue))
			winner, err := game.GetWinner()
			if err != nil {
				t.Errorf("There should not be an error")
			}
			if winner != None {
				t.Errorf("Winner failed: got: %v, want: %v", winner, None)
			}
		})
	})

	t.Run("King Conditions", func(t *testing.T) {
		t.Run("Missing Red King", func(t *testing.T) {
			game := NewGame()
			game.board.setPiece(redTemple, nil)
			winner, err := game.GetWinner()
			if err != nil {
				t.Errorf("There shouldn't be an error")
			}
			if winner != Blue {
				t.Errorf("There should be a winner: got: %v, want: %v", winner, Blue)
			}
		})

		t.Run("Missing Blue King", func(t *testing.T) {
			game := NewGame()
			game.board.setPiece(blueTemple, nil)
			winner, err := game.GetWinner()
			if err != nil {
				t.Errorf("There shouldn't be an error")
			}
			if winner != Red {
				t.Errorf("There should be a winner: got: %v, want: %v", winner, Red)
			}
		})

		t.Run("Missing Red and Blue King", func(t *testing.T) {
			game := NewGame()
			game.board.setPiece(blueTemple, nil)
			game.board.setPiece(redTemple, nil)
			_, err := game.GetWinner()
			if err == nil {
				t.Errorf("There should be an error")
			}
		})
	})
}
