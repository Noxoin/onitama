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
	game := NewGame()
	winner, _ := game.GetWinner()
	if winner == Red || winner == Blue{
		t.Errorf("There shouldn't be a winner")
	}
}
