package goRemi

import (
	"fmt"
)

// Player struct
type Player struct {
	Name     string
	Score    float64
	LastPlay Card
	Hand     []Card
	AI       bool // false player, true ai
}

// Less : sort player
func (p Players) Less(i, j int) bool {
	return p[i].Score < p[j].Score
}

// DrawCards : Draw card from deck
func (p *Player) DrawCards(deck *Deck, numberOfDraw int) {
	fmt.Printf("%s Draw %d card(s)\n", p.Name, numberOfDraw)
	// draw card(s) from deck
	draws := (*deck).Draw(numberOfDraw)
	sortCard(draws)
	p.Hand = append(p.Hand, draws...)
}

// ThrowCards : Return card to deck
func (p *Player) ThrowCards(cardIndex int, deck *Deck) {

	if cardIndex < 0 || cardIndex > (len(p.Hand)-1) {
		println("Invalid card index")
		return
	}

	deck.AddCard(p.Hand[cardIndex])
	p.LastPlay = p.Hand[cardIndex]
	p.Hand = append(p.Hand[:cardIndex], p.Hand[cardIndex+1:]...)

	var cName, cLevel = p.LastPlay.Info()
	fmt.Printf("\n\n%s Play %s %s", p.Name, cLevel, cName)
}

// ShowHand : show human readable info of cards in player hand
func (p Player) ShowHand() {
	for i := 0; i < len(p.Hand); i++ {

		var cName, cLevel = p.Hand[i].Info()
		fmt.Printf("Card Index %d \n", i+1)
		fmt.Print("┌───┐\n")
		fmt.Printf("|%v %v|\n", cLevel, cName)
		fmt.Print("└───┘\n")

		if i < len(p.Hand)-1 {
			//fmt.Printf(",")
		}

	}
}

// Players : collection of player
type Players []Player

// Register : Register Player
func Register(pl []string, numberOfAI int) Players {

	var numberOfPlayers = len(pl) + numberOfAI
	players := make(Players, numberOfPlayers)

	for index := 0; index < len(pl)+numberOfAI; index++ {
		if index <= len(pl)-1 {
			players[index] = Player{Name: pl[index], Score: 0, AI: false}
		} else {
			players[index] = Player{Name: fmt.Sprintf("Computer %d", index), Score: 0, AI: true}
		}
	}

	return players
}
