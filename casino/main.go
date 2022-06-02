package casino

import (
	"math/rand"
	"ssh-poker/game"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Casino struct {
	mu    sync.Mutex
	Games map[string]*game.Game
}

func New() *Casino {
	return &Casino{
		Games: map[string]*game.Game{},
	}
}

// CreateNewGame creates a new game and adds it to the Casino's map of Games
// using a randomly generated UUID as the key.
func (casino *Casino) CreateNewGame() {
	casino.mu.Lock()
	defer casino.mu.Unlock()

	id := uuid.New().String()
	casino.Games[id] = game.New(time.Second * 20)
}

// GetGameById retrieves a pointer to a Game using the passed id.
func (casino *Casino) GetGameById(id string) *game.Game {
	return casino.Games[id]
}

// GetRandomGameId selects a random Game id from all of the Casino's Games.
func (casino *Casino) GetRandomGameId() string {
	casino.mu.Lock()
	defer casino.mu.Unlock()

	randomIndex := rand.Intn(len(casino.Games))

	var id string

	counter := 0
	for key, _ := range casino.Games {
		if counter == randomIndex {
			id = key
			break
		}
		counter++
	}

	return id
}
