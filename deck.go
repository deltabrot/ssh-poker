package main

import (
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

func getValue(value int) string {
	if value == 1 {
		return "A"
	}
	if value == 10 {
		return "\b10"
	}
	if value == 11 {
		return "J"
	}
	if value == 12 {
		return "Q"
	}
	if value == 13 {
		return "K"
	}
	return strconv.Itoa(value)
}
