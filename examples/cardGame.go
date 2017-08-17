package main

import (
	"fmt"

	"github.com/ibrdrahim/goRemi"
)

func main() {

	// initialize Deck
	var newGame goRemi.Cangkulan
	var playerName string

	fmt.Print("Enter Player Name (without space) : ")
	fmt.Scan(&playerName)
	newGame.InitGame(playerName, 2)

	newGame.StartGame()

}
