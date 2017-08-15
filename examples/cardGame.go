package main

import (
	"fmt"

	"github.com/ibrdrahim/goRemi"
)

func main() {

	// initialize Deck
	cDeck := goRemi.InitDeck()

	fmt.Printf("\n\n------- Total Card In Deck %d-------\n", len(cDeck))

	// // show all card in deck
	cDeck.Show()

	// initialize deck, fill deck with card
	// // shuffle card in deck

	fmt.Printf("\n\n------- Shuffle Deck -------\n")

	cDeck.Shuffle()
	// // show all card in deck
	cDeck.Show()

	// initialze player

	players := goRemi.Players{}

	players.Register(0, "Ibrahim")

	players.Register(1, "Computer 1")

	players.Register(2, "Computer 2")

	players.Register(3, "Computer 3")

	// draw card from deck
	for index := range players {
		// draw card
		fmt.Println("\n\n-------Draw-------")
		players[index].DrawCards(&cDeck, 7)
		players[index].ShowHand()
		fmt.Println("\n\n-------DECK-------")
		cDeck.Show()
	}

	// return card to deck
	for index := range players {

		fmt.Printf("\n\n-------%s Return All Card-------\n", players[index].Name)
		players[index].ShowHand()

		for cardIdx := 0; cardIdx < 7; cardIdx++ {
			players[index].ThrowCards(&cDeck)
		}

		// draw card
		fmt.Printf("\n\n-------%s Hand-------\n", players[index].Name)
		players[index].ShowHand()
		fmt.Println("\n\n-------DECK-------")
		cDeck.Show()
	}

	fmt.Printf("\n\n------- Total Card In Deck %d-------\n", len(cDeck))
}
