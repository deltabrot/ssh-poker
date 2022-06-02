// from main()
	/* term := terminal.NewTerminal(os.Stdin, "> ")
	for {
		line, _ := term.ReadLine()
		fmt.Println("LO: ", line)
	} */
	// sshd()

// from test()

	// printSimpleCards(hand)

	// hand = append(hand, deck.Draw(1)...)
	// printSimpleCards(hand)

	// hand = append(hand, deck.Draw(1)...)
	// printCards(hand)

	// buffer := "\033[1;1H\033[2J" + getCardsAscii(deck.Draw(2))

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
