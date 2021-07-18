package main

import (
	"fmt"

	"github.com/chehsunliu/poker"
)

type Player struct {
	Name  string
	Cards []poker.Card
}

func (p *Player) Print() {
	fmt.Println(p.Name)
	if simple {
		printSimpleCards(p.Cards)
	} else {
		printCards(p.Cards)
	}
}
