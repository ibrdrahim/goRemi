package goRemi

import (
	"fmt"
	"sort"
)

// Card : Remi Card struct
type Card struct {
	number int // 2,3,4,5,6,7,8,9,10,jack,queen,king ,ace
	symbol int // ♦ = 0 , ♣ = 1 , ♥ = 2 , ♠ = 3
}

// SortSymbolNumber : sort card by symbol and number
type SortSymbolNumber []Card

func (a SortSymbolNumber) Len() int      { return len(a) }
func (a SortSymbolNumber) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortSymbolNumber) Less(i, j int) bool {
	if a[i].symbol < a[j].symbol {
		return true
	}
	if a[i].symbol > a[j].symbol {
		return false
	}
	return a[i].number < a[j].number
}

func sortCard(cards []Card) {
	sort.Sort(SortSymbolNumber(cards))
}

// Info : Show human readable card info
func (c Card) Info() (string, string) {
	// translate card to human readable info
	var cardName, cardLevel string

	switch c.number {
	case 11:
		cardName = "J"
	case 12:
		cardName = "Q"
	case 13:
		cardName = "K"
	case 14:
		cardName = "A"
	default:
		cardName = fmt.Sprint(c.number)
	}

	// translate grate to human readable info
	switch c.symbol {
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

	return cardName, cardLevel
}

// IsCardHigher : Compare 2 card, return true if card 1 number and symbol is higher
func IsCardHigher(c1 Card, c2 Card) bool {
	if c1.symbol >= c2.symbol && c1.number > c2.number {
		return true
	}
	return false
}
