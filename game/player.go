package game

import (
	"fmt"
	"ssh-poker/render"

	"github.com/chehsunliu/poker"
	"github.com/google/uuid"
)

type Player struct {
	Id    string
	Name  string
	Cards []poker.Card
	Chips int
}

func NewPlayer(name string) *Player {
	return &Player{
		Id:   uuid.New().String(),
		Name: name,
	}
}

// AddCards adds multiple cards to a players Cards.
func (player *Player) AddCards(cards []poker.Card) {
	player.Cards = append(player.Cards, cards...)
}

// Print outputs the player name and their cards to the terminal.
func (player *Player) Print() {
	fmt.Println(player.Name)
	render.PrintCards(player.Cards)
}

// Print outputs the player name and their cards to the terminal in a simple
// form.
func (player *Player) PrintSimple() {
	fmt.Println(player.Name)
	render.PrintSimpleCards(player.Cards)
}
