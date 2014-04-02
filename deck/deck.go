package deck

import (
	"carl/cardgames/card"
	"math/rand"
	"time"
)

type deck interface {
	Shuffle()
	Draw() *card.Card
	CardsRemaining() uint
}

type Deck struct {
	cards [52]*card.Card
	top uint
}

func New() *Deck {
	var deck[52]*card.Card
	for s := card.CLUBS; s <= card.SPADES; s++ {
		for r  := card.ACE; r <= card.KING; r++ {
			deck[int(r)+int(s)*13] = card.New(r, s)
		}
	}
	return &Deck{deck,0}
}

func (d *Deck) Shuffle() {
	d.top = 0
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	for i := 0; i < 52; i++ {
		rIndex := r.Intn(52)
		d.cards[i], d.cards[rIndex] = d.cards[rIndex], d.cards[i]
		if d.cards[i].IsFaceUp() {
			d.cards[i].TurnOver()
		}
		if d.cards[rIndex].IsFaceUp() {
			d.cards[rIndex].TurnOver()
		}
	}
}

func (d *Deck) Draw() *card.Card {
	if d.CardsRemaining() > 0 {
		tmp := d.cards[d.top]
		d.top++
		return tmp
	} else {
		return nil
	}
}

func (d *Deck) CardsRemaining() uint {
	return 52 - d.top
}
