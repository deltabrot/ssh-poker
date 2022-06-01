package main

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/chehsunliu/poker"
	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

var simple = true

type Chat struct {
	Name     string
	Message  string
	IsServer bool
}

type ChatLog struct {
	Chats []Chat
}

func (cl *ChatLog) addChat(name string, message string) {
	cl.Chats = append(cl.Chats, Chat{name, message, false})
}

func (cl *ChatLog) addMessage(message string) {
	cl.Chats = append(cl.Chats, Chat{"", message, true})
}

func (cl *ChatLog) getChats(quantity int) string {
	chats := "── Chat log ──\n"
	if len(cl.Chats) < quantity {
		for i := 0; i < quantity-len(cl.Chats); i++ {
			chats += "\n"
		}
		quantity = len(cl.Chats)
	}
	for _, chat := range cl.Chats[len(cl.Chats)-quantity:] {
		if chat.IsServer {
			chats += fmt.Sprintf("\033[2K\033[0;33m%s\033[0m\n", chat.Message)
		} else {
			chats += fmt.Sprintf("\033[2K%s: %s\n", chat.Name, chat.Message)
		}
	}
	chats += "──────────────\n\n"
	return chats
}

func main() {
	chatLog := ChatLog{}

	hand, deck := game()
	ssh.Handle(func(s ssh.Session) {
		defer s.Close()
		fmt.Printf("\033[0;33m'%s' has connected\033[0m\n", s.User())
		// io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
		term := terminal.NewTerminal(s, "")
		go func() {
			for {
				line, err := term.ReadLine()
				if err != nil {
					fmt.Printf("\033[0;33m'%s' has disconnected\033[0m\n", s.User())
					s.Exit(1)
					break
				}

				if line == "/c" {
					fmt.Printf("%s: \"%s\"\n", s.User(), line)
					chatLog.addMessage(fmt.Sprintf("'%s' checked", s.User()))
					continue
				}

				if line == "/d" {
					hand = append(hand, deck.Draw(1)...)
					handAscii := getCardsAscii(hand)
					cardsAscii := getSimpleCardsAscii(deck.Draw(2))

					io.WriteString(s, "\033[1;1H\033[2J")
					io.WriteString(s, handAscii)
					io.WriteString(s, cardsAscii)
					io.WriteString(s, "\033[11;1H")
					io.WriteString(s, chatLog.getChats(5))
					io.WriteString(s, fmt.Sprintf("[%s] ", s.User()))
					continue
				}

				fmt.Printf("%s: \"%s\"\n", s.User(), line)
				chatLog.addChat(s.User(), line)

				io.WriteString(s, "\033[11;1H")
				io.WriteString(s, chatLog.getChats(5))
				io.WriteString(s, fmt.Sprintf("[%s] ", s.User()))
				time.Sleep(100 * time.Millisecond)
			}
		}()

		handAscii := getCardsAscii(hand)
		cardsAscii := getSimpleCardsAscii(deck.Draw(2))

		io.WriteString(s, "\033[1;1H\033[2J")
		io.WriteString(s, handAscii)
		io.WriteString(s, cardsAscii)
		for {
			// io.WriteString(s, "\033[2J")
			io.WriteString(s, "\033[11;1H")
			io.WriteString(s, chatLog.getChats(5))
			io.WriteString(s, fmt.Sprintf("[%s] ", s.User()))
			time.Sleep(100 * time.Millisecond)
		}
	})

	log.Println("starting ssh server on port 2222...")
	log.Fatal(
		ssh.ListenAndServe(
			":2222",
			nil,
			ssh.HostKeyFile("./server_key"),
		),
	)

	/* term := terminal.NewTerminal(os.Stdin, "> ")
	for {
		line, _ := term.ReadLine()
		fmt.Println("LO: ", line)
	} */
	// sshd()
}

func game() ([]poker.Card, *poker.Deck) {
	players := []Player{
		{1, "deltabrot", []poker.Card{}},
		{2, "DryYourRice", []poker.Card{}},
		{3, "Fishy Dish", []poker.Card{}},
	}

	deck := poker.NewDeck()

	for i := 0; i < len(players); i++ {
		players[i].Cards = deck.Draw(2)
	}

	hand := deck.Draw(3)
	// printSimpleCards(hand)

	// hand = append(hand, deck.Draw(1)...)
	// printSimpleCards(hand)

	// hand = append(hand, deck.Draw(1)...)
	// printCards(hand)

	// buffer := "\033[1;1H\033[2J" + getCardsAscii(deck.Draw(2))
	return hand, deck
}

/* func game(channel ssh.Channel) {
		players := []Player{
			{1, "deltabrot", []poker.Card{}},
			{2, "DryYourRice", []poker.Card{}},
			{3, "Fishy Dish", []poker.Card{}},
		}

		deck := poker.NewDeck()

		for i := 0; i < len(players); i++ {
			players[i].Cards = deck.Draw(2)
		}

		hand := deck.Draw(3)
		// printSimpleCards(hand)

		hand = append(hand, deck.Draw(1)...)
		// printSimpleCards(hand)

		hand = append(hand, deck.Draw(1)...)
		// printCards(hand)

		// smallest := int32(999999)
		/* winner := ""
		winnerString := ""
		for _, player := range players {
			rank := poker.Evaluate(append(hand, player.Cards...))
			rankString := poker.RankString(rank)

			if smallest > rank {
				smallest = rank
				winner = player.Name
				winnerString = rankString
			}

			fmt.Println(
				player.Name + ": " + rankString + " [" + strconv.Itoa(int(rank)) + "]",
			)
			printSimpleCards(player.Cards)
		}

		// fmt.Println("Winner: " + winner + " (" + winnerString + ")")
	}
} */
