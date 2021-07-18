package main

import (
	"fmt"
	"sort"
)

type Hand int

const (
	HighCard Hand = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

func (h Hand) String() string {
	if h == HighCard {
		return "High Card"
	}
	if h == OnePair {
		return "One Pair"
	}
	if h == TwoPair {
		return "Two Pair"
	}
	if h == ThreeOfAKind {
		return "Three Of A Kind"
	}
	if h == Straight {
		return "Straight"
	}
	if h == Flush {
		return "Flush"
	}
	if h == FullHouse {
		return "Full House"
	}
	if h == FourOfAKind {
		return "Four Of A Kind"
	}
	if h == StraightFlush {
		return "Straight Flush"
	}
	if h == RoyalFlush {
		return "Royal Flush"
	}

	return "Invalid"
}

func calculateHand(cards []Card) Hand {
	hand := HighCard

	isOnePair := false
	isTwoPair := false
	isThreeOfAKind := false
	isStraight := false
	isFlush := false
	isFullHouse := false
	isFourOfAKind := false
	isStraightFlush := false
	isRoyalFlush := false

	valueGroups := make(map[int]int)
	suitGroups := make(map[Suit]int)
	for _, card := range cards {
		valueGroups[card.Value]++
		suitGroups[card.Suit]++
	}

	// value groups
	valueGroupCount := 0
	for _, quantity := range valueGroups {
		if quantity > 1 {
			valueGroupCount++
			if quantity == 2 {
				if isOnePair {
					isTwoPair = true
				}
				if isThreeOfAKind {
					isFullHouse = true
				}
				isOnePair = true
			}
			if quantity == 3 {
				isThreeOfAKind = true
				if isOnePair {
					isFullHouse = true
				}
			}
			if quantity == 4 {
				isFourOfAKind = true
			}
		}
	}

	// flush
	for _, quantity := range suitGroups {
		if quantity >= 5 {
			isFlush = true
		}
	}

	// straight
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Value < cards[j].Value
	})

	fmt.Println(cards)

	for i := 0; i < 3; i++ {
		straightCount := 0
		lastValue := cards[i].Value - 1
		for _, card := range cards {
			if lastValue == card.Value-1 || (lastValue == 13 && card.Value == 1) {
				straightCount++
			}
			if lastValue == card.Value {
				continue
			}
			if straightCount == 5 {
				isStraight = true
			}
			lastValue = card.Value
		}
	}

	if isRoyalFlush {
		hand = RoyalFlush
	} else if isStraightFlush {
		hand = StraightFlush
	} else if isFourOfAKind {
		hand = FourOfAKind
	} else if isFullHouse {
		hand = FullHouse
	} else if isFlush {
		hand = Flush
	} else if isStraight {
		hand = Straight
	} else if isThreeOfAKind {
		hand = ThreeOfAKind
	} else if isTwoPair {
		hand = TwoPair
	} else if isOnePair {
		hand = OnePair
	}

	return hand
}
