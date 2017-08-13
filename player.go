package goRemi

import ( 
    "fmt"
)

type Player struct {
	name string
	score float64
	hand []Card 
}

func (p *Player) DrawCards(deck *Deck, numberOfDraw int){
	fmt.Printf("%s Draw %d card(s)\n",p.name,numberOfDraw)
	// draw card(s) from deck
	draws := deck.Draw(numberOfDraw)
	p.hand = draws
}

func (p Player) ShowHand() {
	for i := 0; i < len(p.hand); i++ {
		
        var cName,cLevel = p.hand[i].Info()
        fmt.Printf("%v %v",cLevel,cName)

        if( i < len(p.hand) - 1){
		    fmt.Printf(",")
        }

	}
}

type Players [4]Player

func (p *Players) Register(index int, name string) {
	p[index].name = name
}
