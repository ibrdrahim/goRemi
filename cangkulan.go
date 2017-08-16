package goRemi

import (
	"fmt"
)

// Cangkulan : Cangkulan The Game
// 1. shuffle deck
// 2. put 1 card from top of deck to table
// 3. give each player 7 cards
// 4. start playing from player 1 (clockwise)
// 5. each player put one card with the same flag as the card on the table
// 6. if player doesnt have card with the same flag draw card from deck until player get the card with the same flag
//  7.a. if deck doesnt have any card, draw card from the top of the table
// 8. after all player throw card to the table, compare players card, player with higher number win the round and get the score
// 9. repeat 5-8 until deck is empty
type Cangkulan struct {
	Deck    Deck
	Field   Deck
	Players []Player
}

// InitGame : create new game instance
func (c *Cangkulan) InitGame(numberOfPlayers int) {
	fmt.Printf("\n\n------- Register Player -------\n")
	c.Players = Register([]string{"ibrahim"}, 3)
	// init Deck , fill deck with cards
	fmt.Printf("\n\n------- Preparing Deck -------\n")
	c.Deck = InitDeck()
	// field for playing card
	c.Field = make([]Card, 0)
}

// StartGame : start game instance
func (c *Cangkulan) StartGame() {

	var isPlaying = true

	fmt.Printf("\n\n------- Shuffle Deck -------\n")
	c.Deck.Shuffle()

	fmt.Println("\n\n-------FIELD-------")
	c.Field.AddCard(c.Deck.Draw(1)[0])

	c.Field.Show()
	// all player draw card from deck

	// draw card from deck
	for index := range c.Players {
		// draw card
		fmt.Println("\n\n-------Draw-------")
		c.Players[index].DrawCards(&c.Deck, 7)
		c.Players[index].ShowHand()
	}

	var round = 0

	for isPlaying {
		// return card to deck
		for index := range c.Players {

			fmt.Println("\n\n-------FIELD-------")
			c.Field.Show()
			fmt.Printf("\n\n-------%s Play Card-------\n", c.Players[index].Name)
			c.Players[index].ShowHand()

			var playerTurn = true
			for playerTurn {

				// select first card for last round winner
				if round > 0 && index == 0 {
					// throw card with smalest symbol and number
					c.Players[index].ThrowCards(0, &c.Field)
					//playerTurn = false
					break
				} else {
					for cIdx := range c.Players[index].Hand {
						if c.Players[index].Hand[cIdx].symbol == c.Field[len(c.Field)-1].symbol {
							c.Players[index].ThrowCards(cIdx, &c.Field)
							playerTurn = false
							break
						}
					}

					// Draw a new card until player have playable card
					for playerTurn {

						if len(c.Deck) > 0 {

							fmt.Println("\n\n-------Draw-------")
							c.Players[index].DrawCards(&c.Deck, 1)
							fmt.Printf("\n\n-------%s Hand-------\n", c.Players[index].Name)
							c.Players[index].ShowHand()

							if c.Players[index].Hand[len(c.Players[index].Hand)-1].symbol == c.Field[len(c.Field)-1].symbol {
								c.Players[index].ThrowCards(len(c.Players[index].Hand)-1, &c.Field)
								playerTurn = false
								break
							}

						} else {

							// draw card from field because deck is empty
							fmt.Println("\n\n-------Penalty Draw From FIELD-------")
							c.Players[index].DrawCards(&c.Field, 1)
							fmt.Printf("\n\n-------%s Hand-------\n", c.Players[index].Name)
							c.Players[index].ShowHand()

							playerTurn = false
							break
						}
					}
				}

			}

			if len(c.Players[index].Hand) < 1 {
				fmt.Printf("\n\n-------%s WIN-------\n", c.Players[index].Name)
				isPlaying = false
				break
			}

		}

		//fmt.Printf("\n\n-------ROUND %d OVER-------\n", round)
		round = round + 1
		// re order player based on last round winner
		var biggestHand = c.Players[0].LastPlay
		var roundWinnerIdx = 0
		for index := range c.Players {
			if IsCardHigher(biggestHand, c.Players[index].LastPlay) {
				biggestHand = c.Players[index].LastPlay
				roundWinnerIdx = index
			}
		}

		if roundWinnerIdx > 0 {

			fmt.Printf(" Winner : %d\n", roundWinnerIdx)
			var tempPlayer = make(Players, len(c.Players))
			copy(tempPlayer, c.Players)
			var newCounter = 0
			for index := range c.Players {

				if (roundWinnerIdx + index) <= (len(c.Players) - 1) {
					c.Players[index] = tempPlayer[roundWinnerIdx+index]
				} else {
					c.Players[index] = tempPlayer[newCounter]
					newCounter = newCounter + 1
				}
			}

			fmt.Println(c.Players)

		}
	}

	fmt.Println("\n\n-------GAME OVER-------")

}
