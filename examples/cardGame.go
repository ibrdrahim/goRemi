package main

import (
	"github.com/ibrdrahim/goRemi"
)

func main() {

	// initialize Deck
	var newGame goRemi.Cangkulan

	newGame.InitGame(4)

	newGame.StartGame()

}
