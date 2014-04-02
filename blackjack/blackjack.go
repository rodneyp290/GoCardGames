package blackjack

import (
	"carl/cardsgames/card"
	"carl/cardsgames/deck"
)

type Game struct {
	dck deck.Deck
	player Hand
	dealer Hand
}


