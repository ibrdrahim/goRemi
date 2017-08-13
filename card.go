package goRemi

import ( 
    "fmt"
)

type Card struct {
    number int // ace , 2,3,4,5,6,7,8,9,10,jack,queen,king
    grade int // ♦ = 0 , ♣ = 1 , ♥ = 2 , ♠ = 3    
}

func (c Card) Info() (string,string){
    // translate card to human readable info
    var cardName,cardLevel string

    switch c.number {
    case 1:
        cardName = "Ace"
    case 11:
        cardName = "Jack"
    case 12:
        cardName = "Queen"
    case 13:
        cardName = "King"
    default:
        cardName = fmt.Sprint(c.number)
    }

    // translate grate to human readable info
    switch c.grade {
    case 0:
        cardLevel = "♦"
    case 1:
        cardLevel = "♣"
    case 2:
        cardLevel = "♥"
    case 3:
        cardLevel = "♠"
    default:
    }

    return cardName,cardLevel
}