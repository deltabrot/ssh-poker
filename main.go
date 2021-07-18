package main

import (
	"fmt"

	"github.com/chehsunliu/poker"
)

var simple = true

func main() {
	players := []Player{
		{"deltabrot", []poker.Card{}},
		{"DryYourRice", []poker.Card{}},
		{"Fishy Dish", []poker.Card{}},
	}

	deck := poker.NewDeck()

	for i := 0; i < len(players); i++ {
		players[i].Cards = deck.Draw(2)
	}

	hand := deck.Draw(3)
	printSimpleCards(hand)

	hand = append(hand, deck.Draw(1)...)
	printSimpleCards(hand)

	hand = append(hand, deck.Draw(1)...)
	printSimpleCards(hand)

	for _, player := range players {
		rank := poker.Evaluate(append(hand, player.Cards...))
		rankString := poker.RankString(rank)

		fmt.Println(fmt.Sprintf(
			"%s: %d [%s]",
			player.Name,
			rank,
			rankString,
		))
	}
}
