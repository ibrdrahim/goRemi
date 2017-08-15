package goRemi

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
	deck    Deck
	field   Deck
	players []Player
}

func (c *Cangkulan) initGame(numberOfPlayers int) {
	c.deck = InitDeck()
	c.field = InitDeck()
	c.players = make([]Player, numberOfPlayers)
}
