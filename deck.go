package goRemi

import (
	"fmt"
	"math/rand"
	"time"
)

// Deck : Array collection of card
type Deck []Card

// InitDeck : initialize dynamic array Deck
func InitDeck() Deck {

	var deck = make([]Card, 52)
	for i := 0; i <= len(deck)-1; i++ {
		deck[i].number = (i % 13) + 2
		deck[i].symbol = (i / 13)
	}

	return deck
}

// Shuffle Deck
func (deck Deck) Shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
}

// Show Deck Content
func (deck Deck) Show() {
	for i := 0; i <= len(deck)-1; i++ {
		var cName, cLevel = deck[i].Info()
		fmt.Printf("%s %s", cLevel, cName)

		if i < len(deck)-1 {
			fmt.Printf(",")
		}
	}
	fmt.Println()
}

// Draw card from deck
func (deck *Deck) Draw(numberOfCards int) Deck {

	var draw = (*deck)[len(*deck)-numberOfCards : len(*deck)]
	*deck = (*deck)[:len(*deck)-numberOfCards]

	return draw
}

// AddCard : add card to top of deck
func (deck *Deck) AddCard(card Card) {
	*deck = append(*deck, card)
}
