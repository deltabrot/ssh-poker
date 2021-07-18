package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Suit int

const (
	Spades Suit = iota
	Hearts
	Clubs
	Diamonds
)

type Card struct {
	Value int
	Suit  Suit
}

func (c *Card) getValue() string {
	if c.Value == 1 {
		return "A"
	}
	if c.Value == 10 {
		return "\b10"
	}
	if c.Value == 11 {
		return "J"
	}
	if c.Value == 12 {
		return "Q"
	}
	if c.Value == 13 {
		return "K"
	}
	return strconv.Itoa(c.Value)
}

func (c *Card) getSuit() string {
	if c.Suit == Spades {
		return "♠"
	}
	if c.Suit == Hearts {
		return "♥"
	}
	if c.Suit == Clubs {
		return "♣"
	}
	if c.Suit == Diamonds {
		return "♦"
	}
	return " "
}

func (c *Card) getPrintBuffer() []string {
	var startColor string
	if c.Suit == Diamonds || c.Suit == Hearts {
		startColor = "\u001b[31m\u001b[47m"
	} else {
		startColor = "\u001b[30m\u001b[47m"
	}
	resetColor := "\u001b[0m"

	printBuffer := []string{
		startColor + fmt.Sprintf(" %s       ", c.getValue()) + resetColor,
		startColor + fmt.Sprintf(" %s       ", c.getSuit()) + resetColor,
		startColor + "         " + resetColor,
		startColor + fmt.Sprintf("    %s    ", c.getSuit()) + resetColor,
		startColor + "         " + resetColor,
		startColor + fmt.Sprintf("       %s ", c.getSuit()) + resetColor,
		startColor + fmt.Sprintf("       %s ", c.getValue()) + resetColor,
	}

	return printBuffer
}

func (c *Card) printCard() {
	var startColor string
	if c.Suit == Diamonds || c.Suit == Hearts {
		startColor = "\u001b[31m\u001b[47m"
	} else {
		startColor = "\u001b[30m\u001b[47m"
	}
	resetColor := "\u001b[0m"

	fmt.Println(
		fmt.Sprintf(
			`%[1]s %[4]s       %[2]s
`+
				`%[1]s %[3]s       %[2]s
`+
				`%[1]s         %[2]s
`+
				`%[1]s    %[3]s    %[2]s
`+
				`%[1]s         %[2]s
`+
				`%[1]s       %[3]s %[2]s
`+
				`%[1]s       %[4]s %[2]s`,
			startColor,
			resetColor,
			c.getSuit(),
			c.getValue(),
		),
	)
}

func printCards(cards []Card) {
	for i := 0; i < 7; i++ {
		for _, card := range cards {
			fmt.Print(card.getPrintBuffer()[i] + " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func generateDeck() []Card {
	deck := []Card{}
	suits := []Suit{Spades, Hearts, Clubs, Diamonds}
	for _, suit := range suits {
		for value := 1; value <= 13; value++ {
			card := Card{value, suit}
			deck = append(deck, card)
		}
	}

	return deck
}

func shuffleCards(cards []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(
		len(cards),
		func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		},
	)
	return cards
}

func main() {
	deck := generateDeck()
	deck = shuffleCards(deck)
	printCards(deck[0:5])
}
