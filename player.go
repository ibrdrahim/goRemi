package goRemi

import (
	"fmt"
)

// Player struct
type Player struct {
	Name  string
	Score float64
	Hand  []Card
}

// DrawCards : Draw card from deck
func (p *Player) DrawCards(deck *Deck, numberOfDraw int) {
	fmt.Printf("%s Draw %d card(s)\n", p.Name, numberOfDraw)
	// draw card(s) from deck
	draws := deck.Draw(numberOfDraw)
	sortCard(draws)
	p.Hand = draws
}

// ThrowCards : Return card to deck
func (p *Player) ThrowCards(deck *Deck) {
	deck.AddCard(p.Hand[len(p.Hand)-1])
	p.Hand = p.Hand[:len(p.Hand)-1]
}

// ShowHand : show human readable info of cards in player hand
func (p Player) ShowHand() {
	for i := 0; i < len(p.Hand); i++ {

		var cName, cLevel = p.Hand[i].Info()
		fmt.Printf("%v %v", cLevel, cName)

		if i < len(p.Hand)-1 {
			fmt.Printf(",")
		}

	}
}

// Players : collection of player
type Players []Player

// Register : Register Player
func Register(pl []string, numberOfAI int) Players {

	players := make(Players, len(pl))

	for index := range pl {
		players = append(players, Player{Name: pl[index], Score: 0})
	}

	for idx := 0; idx < numberOfAI; idx++ {
		players = append(players, Player{Name: fmt.Sprintf("Computer %d", idx+1), Score: 0})
	}

	return players
}
