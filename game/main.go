package game

import (
	"ssh-poker/chat"
	"time"

	"github.com/chehsunliu/poker"
)

type GameState int

const (
	Waiting GameState = iota
	PreFlop
	Flop
	Turn
	River
	Showdown
)

type Game struct {
	Players     []*Player
	Deck        *poker.Deck
	Cards       []poker.Card
	State       GameState
	TurnTimeout time.Duration
	Pot         int
	Chat        *chat.Chat
}

// New initialises a new game pointer.
func New(turnTimeout time.Duration) *Game {
	return &Game{
		Deck:        poker.NewDeck(),
		Pot:         0,
		State:       Waiting,
		TurnTimeout: turnTimeout,
		Chat:        chat.New(),
	}
}

// AddPlayer adds a new player to the game.
func (game *Game) AddPlayer(player *Player) {
	game.Players = append(game.Players, player)
}

// NextRound sets the state of the game to the next round.
func (game *Game) NextRound() {
	game.State++
}
