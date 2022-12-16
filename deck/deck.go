package deck

import (
	"fmt"
	"math/rand"
	"time"

	"joycasino/card"
)

type Deck struct {
	deck  []card.Card
	index int
}

func New() *Deck {
	deck := make([]card.Card, 0, len(card.Ranks)*len(card.Suits))
	for _, r := range card.Ranks {
		for _, s := range card.Suits {
			deck = append(deck, card.New(r, s))
		}
	}
	return &Deck{deck: deck}
}

func (d *Deck) Shuffle() *Deck {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(d.deck), func(i, j int) {
		d.deck[i], d.deck[j] = d.deck[j], d.deck[i]
	})
	return d

}

func (d *Deck) TakeCard() *card.Card {
	c := d.deck[d.index]
	fmt.Println(string(c.Rank()) + string(c.Suit()))
	d.index++
	return &c
}
