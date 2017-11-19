package game

import (
	"testing"
)

func setupCardsInGame(game *Game) {
  c := GetRandomCards(5, 0)  // will get crane, eel, monkey, boar, dragon
  game.cards = make(map[Team]map[string]*Card)
  game.cards[Red] = make(map[string]*Card)
  game.cards[Blue] = make(map[string]*Card)
  game.cards[Red][c[0].Name] = c[0]
  game.cards[Red][c[1].Name] = c[1]
  game.cards[Blue][c[2].Name] = c[2]
  game.cards[Blue][c[3].Name] = c[3]
  game.neutralCard = c[4]
}

func TestValidateMoveMissingStarPiece(t *testing.T) {
	game := NewGame()
	setupCardsInGame(game)
	err := game.validateMove(Cord{1, 1}, Cord{0, 0}, "card")
	if err == nil {
		t.Errorf("Should be an error")
	}
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
