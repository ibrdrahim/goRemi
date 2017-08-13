package goRemi

import ( 
    "fmt"
    "math/rand"
    "time"
)

type Deck []Card

func Init() Deck{

	var deck = make([]Card,52)
    for i := 0; i <= len(deck)-1; i++ {
        deck[i].number = (i%13)+1
        deck[i].grade  = (i/13)
    }

    return deck
}

func (deck Deck) Shuffle(){
    rand.Seed(time.Now().UTC().UnixNano())
    for i := len(deck) - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        deck[i], deck[j] = deck[j], deck[i]
    }
}

func (deck Deck) Show(){
    for i := 0; i <= len(deck)-1; i++ {
        var cName,cLevel = deck[i].Info()
        fmt.Printf("%v %v",cLevel,cName)

        if( i < len(deck) - 1){
		    fmt.Printf(",")
        }
    }
    fmt.Println()
}

func (deck *Deck) Draw(numberOfCards int) (Deck){
	
	var draw = (*deck)[len(*deck)-numberOfCards:len(*deck)]
	*deck = (*deck)[:len(*deck)-numberOfCards]

	return draw
}

