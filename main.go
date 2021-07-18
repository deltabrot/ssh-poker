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

	/* deck := generateDeck()
	deck = shuffleCards(deck)

	for i := 0; i < len(players); i++ {
		players[i].Deal(&deck)
	}

	if simple {
		printSimpleCards(deck[0:5])
		// printSimpleCards([]Card{deck[3]})
		// printSimpleCards([]Card{deck[4]})
	} else {
		printCards(deck[0:5])
		// printCards([]Card{deck[3]})
		// printCards([]Card{deck[4]})
	}

	for _, player := range players {
		hand := calculateHand(append(deck[0:5], player.Cards...))
		fmt.Println(player.Name + ": " + hand.String())
		if simple {
			printSimpleCards(player.Cards)
		} else {
			printCards(player.Cards)
		}
	} */
}
