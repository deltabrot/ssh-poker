package main

import (
	"fmt"
	"io"
	"log"
	"ssh-poker/casino"
	"ssh-poker/game"
	"ssh-poker/render"
	"time"

	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

var simple = true

func main() {
	// create casino
	casino := casino.New()

	// create new game
	casino.CreateNewGame()

	handleSsh(casino)

	log.Println("starting ssh server on port 2222...")
	log.Fatal(
		ssh.ListenAndServe(
			":2222",
			nil,
			ssh.HostKeyFile("./keys/server_key"),
		),
	)
}

func handleSsh(casino *casino.Casino) {
	ssh.Handle(func(s ssh.Session) {
		defer s.Close()

		// alert server a new connection has been made
		fmt.Printf("\033[0;33m'%s' has connected\033[0m\n", s.User())

		// greet new player
		// io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))

		// find random game
		currentGame := casino.GetGameById(casino.GetRandomGameId())

		// create player entity
		player := game.NewPlayer(s.User())

		// add player to currentGame
		currentGame.AddPlayer(player)

		term := terminal.NewTerminal(s, "")
		go func() {
			for {
				line, err := term.ReadLine()
				if err != nil {
					fmt.Printf("\033[0;33m'%s' has disconnected\033[0m\n", s.User())
					s.Exit(1)
					break
				}

				// check
				if line == "/c" {
					fmt.Printf("%s: \"%s\"\n", s.User(), line)
					currentGame.Chat.AddMessage("", fmt.Sprintf("'%s' checked", s.User()))
					continue
				}

				// draw card
				if line == "/d" {
					currentGame.Cards = append(currentGame.Cards, currentGame.Deck.Draw(1)...)
					handAscii := render.GetCardsAscii(currentGame.Cards)
					cardsAscii := render.GetSimpleCardsAscii(currentGame.Deck.Draw(2))

					io.WriteString(s, "\033[1;1H\033[2J")
					io.WriteString(s, handAscii)
					io.WriteString(s, cardsAscii)
					io.WriteString(s, "\033[11;1H")
					io.WriteString(s, currentGame.Chat.GetMessages(5))
					io.WriteString(s, fmt.Sprintf("[%s] ", s.User()))
					continue
				}

				// fmt.Printf("hehehe%s: \"%s\"\n", s.User(), line)
				currentGame.Chat.AddMessage(s.User(), line)

				io.WriteString(s, "\033[11;1H")
				io.WriteString(s, currentGame.Chat.GetMessages(5))
				io.WriteString(s, fmt.Sprintf("[%s] ", s.User()))
				time.Sleep(100 * time.Millisecond)
			}
		}()

		handAscii := render.GetCardsAscii(currentGame.Cards)
		cardsAscii := render.GetSimpleCardsAscii(currentGame.Deck.Draw(2))

		io.WriteString(s, "\033[1;1H\033[2J")
		io.WriteString(s, handAscii)
		io.WriteString(s, cardsAscii)
		for {
			// io.WriteString(s, "\033[2J")
			io.WriteString(s, "\033[11;1H")
			io.WriteString(s, currentGame.Chat.GetMessages(5))
			io.WriteString(s, fmt.Sprintf("[%s] ", s.User()))
			time.Sleep(100 * time.Millisecond)
		}
	})

}
