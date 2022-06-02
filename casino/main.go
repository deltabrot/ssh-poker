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

func (casino *Casino) CreateNewGame() {
	casino.mu.Lock()
	defer casino.mu.Unlock()

	id := uuid.New().String()
	casino.Games[id] = game.New(time.Second * 20)
}

func (casino *Casino) GetGameById(id string) *game.Game {
	return casino.Games[id]
}

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
