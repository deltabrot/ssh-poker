package render

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

// getPrintBuffer uses the passed list of cards to generate a string such that
// they can be rendered in order, using multiple lines.
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

// getSimplePrintBuffer uses the passed list of cards to generate a string such
// that they can be rendered in order, using a single line.
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

// PrintCards prints full size playing cards in the terminal complete with
// color.
func PrintCards(cards []poker.Card) {
	for i := 0; i < 7; i++ {
		for _, card := range cards {
			fmt.Print(getPrintBuffer(card)[i] + " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

// PrintSimpleCards prints small playing cards in the terminal complete with
// color.
func PrintSimpleCards(cards []poker.Card) {
	for _, card := range cards {
		fmt.Print(getSimplePrintBuffer(card) + " ")
	}
	fmt.Println()
	fmt.Println()
}

// GetCardsAscii retrieves a string which when output in a terminal would
// display full size playing cards in order complete with color.
func GetCardsAscii(cards []poker.Card) string {
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

// GetSimpleCardsAscii retrieves a string which when output in a terminal would
// display small playing cards in order complete with color.
func GetSimpleCardsAscii(cards []poker.Card) string {
	ascii := ""
	for _, card := range cards {
		ascii += getSimplePrintBuffer(card) + " "
	}
	ascii += "\n\n"
	return ascii
}
