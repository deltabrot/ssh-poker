package main

import (
	"fmt"

	"github.com/chehsunliu/poker"
)

var (
	strRanks = "23456789TJQKA"
)

var (
	prettySuits = map[int32]string{
		1: "♠", // spades
		2: "♥", // hearts
		4: "♦", // diamonds
		8: "♣", // clubs
	}
)

func getPrintBuffer(card poker.Card) []string {
	var startColor string
	if card.Suit() == 2 || card.Suit() == 4 {
		startColor = "\u001b[31m\u001b[47m"
	} else {
		startColor = "\u001b[30m\u001b[47m"
	}
	resetColor := "\u001b[0m"

	rank := string(strRanks[card.Rank()])
	suit := string(prettySuits[card.Suit()])

	printBuffer := []string{
		startColor + fmt.Sprintf(" %s       ", rank) + resetColor,
		startColor + fmt.Sprintf(" %s       ", suit) + resetColor,
		startColor + "         " + resetColor,
		startColor + fmt.Sprintf("    %s    ", suit) + resetColor,
		startColor + "         " + resetColor,
		startColor + fmt.Sprintf("       %s ", suit) + resetColor,
		startColor + fmt.Sprintf("       %s ", rank) + resetColor,
	}

	return printBuffer
}

func getSimplePrintBuffer(card poker.Card) string {
	var startColor string
	if card.Suit() == 2 || card.Suit() == 4 {
		startColor = "\u001b[31m\u001b[47m"
	} else {
		startColor = "\u001b[30m\u001b[47m"
	}
	resetColor := "\u001b[0m"

	rank := string(strRanks[card.Rank()])
	suit := string(prettySuits[card.Suit()])

	printBuffer := startColor + fmt.Sprintf(" %s %s ", rank, suit) + resetColor

	return printBuffer
}

func printCards(cards []poker.Card) {
	for i := 0; i < 7; i++ {
		for _, card := range cards {
			fmt.Print(getPrintBuffer(card)[i] + " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func printSimpleCards(cards []poker.Card) {
	for _, card := range cards {
		fmt.Print(getSimplePrintBuffer(card) + " ")
	}
	fmt.Println()
	fmt.Println()
}
