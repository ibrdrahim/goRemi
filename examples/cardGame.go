package main
import ( 
    "fmt"
    "github.com/ibrdrahim/goRemi"
)

func main() {
    cDeck := goRemi.Init()

    // initialize deck, fill deck with card
    // // shuffle card in deck
    cDeck.Shuffle()
    // // show all card in deck
    cDeck.Show()

    // initialze player

    players := goRemi.Players{}

    players.Register(0,"Ibrahim")

    players.Register(1,"Computer 1")
    
    players.Register(2,"Computer 2")
    
    players.Register(3,"Computer 3")

    for _,player := range players {
        // draw card
        fmt.Println("\n\n-------Draw-------\n")
        player.DrawCards(&cDeck,7)
        player.ShowHand()
        fmt.Println("\n\n-------DECK-------\n")
        cDeck.Show()
    }

}