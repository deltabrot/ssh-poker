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
		startColor = "\u001b[31m\u001b[1;47m"
	} else {
		startColor = "\u001b[30m\u001b[1;47m"
	}
	resetColor := "\u001b[0m"

	rank := string(strRanks[card.Rank()])
	suit := string(prettySuits[card.Suit()])

	if rank == "T" {
		rank = "\b" + "10"
	}

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
		startColor = "\u001b[31m\u001b[1;47m"
	} else {
		startColor = "\u001b[30m\u001b[1;47m"
	}
	resetColor := "\u001b[0m"

	rank := string(strRanks[card.Rank()])
	suit := string(prettySuits[card.Suit()])

	printBuffer := startColor + fmt.Sprintf(" %s %s ", rank, suit) + resetColor
	if rank == "T" {
		rank = "10"
		printBuffer = startColor + fmt.Sprintf("%s %s ", rank, suit) + resetColor
	}

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

func getCardsAscii(cards []poker.Card) string {
	ascii := ""
	for i := 0; i < 7; i++ {
		for j, card := range cards {
			ascii += getPrintBuffer(card)[i]
			if j != len(cards)-1 {
				ascii += " "
			}
		}
		ascii += "\n"
	}
	ascii += "\n"
	return ascii
}

func getSimpleCardsAscii(cards []poker.Card) string {
	ascii := ""
	for _, card := range cards {
		ascii += getSimplePrintBuffer(card) + " "
	}
	ascii += "\n\n"
	return ascii
}
