package main

import (
	"carl/cardgames/deck"
	"fmt"
)

func main() {
	dk := deck.New()
	i := 0
	for dk.CardsRemaining() > 0 {
		c := dk.Draw()
		fmt.Printf("%d: %s", i, c)
		c.TurnOver()
		fmt.Printf(" -flip- %s\n", c)
		i++
	}
	dk.Shuffle()
	i = 0
	for dk.CardsRemaining() > 0 {
		c := dk.Draw()
		fmt.Printf("%d: %s", i, c)
		c.TurnOver()
		fmt.Printf(" -flip- %s\n", c)
		i++
	}
}
