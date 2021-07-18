package main

import (
	"github.com/chehsunliu/poker"
)

type Player struct {
	Name  string
	Cards []poker.Card
}

/* func (p *Player) Deal(cards *[]Card) {
	p.Cards, *cards = (*cards)[0:2], (*cards)[2:]
}

func (p *Player) Print() {
	fmt.Println(p.Name)
	if simple {
		printSimpleCards(p.Cards)
	} else {
		printCards(p.Cards)
	}
} */
