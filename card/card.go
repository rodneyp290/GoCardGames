package card

import (
	"fmt"
)

type Rank uint

const (
	ACE Rank = iota
	TWO Rank = iota
	THREE Rank = iota
	FOUR Rank = iota
	FIVE Rank = iota
	SIX Rank = iota
	SEVEN Rank = iota
	EIGHT Rank = iota
	NINE Rank = iota
	TEN Rank = iota
	JACK Rank = iota
	QUEEN Rank = iota
	KING Rank = iota
)

func (r *Rank) String() string {
	switch *r {
	case ACE: return "A"
	case JACK: return "J"
	case QUEEN: return "Q"
	case KING: return "K"
	default: return fmt.Sprintf("%d",*r)
	}
}

type Suit uint

const (
	CLUBS Suit = iota
	DIAMONDS Suit = iota
	HEARTS Suit = iota
	SPADES Suit = iota
)

func (s *Suit) String() string {
	switch *s {
	case CLUBS: return "C"
	case DIAMONDS: return "D"
	case HEARTS: return "H"
	case SPADES: return "S"
	default: return fmt.Sprintf("%d", *s)
	}
}

type Carder interface {
	TurnOver()
	String() string
	IsFaceUp() bool
	GetSuit() Suit
	GetRank() Rank
}

type Card struct {
	rank Rank
	suit Suit
	faceUp bool
}

func New(r Rank, s Suit) *Card {
	return &Card{r, s, false}
}

func (c *Card) TurnOver() {
	c.faceUp = !c.faceUp
}

func (c *Card) String() string {
	if c.faceUp {
		return fmt.Sprintf("%s%s", c.suit.String(),
			c.rank.String())
	} else {
		return "**"
	}
}

func (c *Card) IsFaceUp() bool {
	return c.faceUp
}

func (c *Card) GetSuit() Suit {
	return c.suit
}

func (c *Card) GetRank() Rank {
	return c.rank
}
