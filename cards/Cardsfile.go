package cards

import (
	"fmt"
	"math/rand"
)

var Trash = map[string]Card{}

var (
	Cards = []string{
		"двойка",
		"тройка",
		"четверка",
		"пятерка",
		"шестерка",
		"семерка",
		"восьмерка",
		"девятка",
		"десятка",
		"валет",
		"дама",
		"король",
		"туз",
	}
)

type Suit rune

const (
	Clubs    Suit = '♧'
	Diamonds Suit = '♦'
	Hearts   Suit = '♥'
	Spades   Suit = '♤'
)

var Suits = [4]Suit{Clubs, Diamonds, Hearts, Spades}

// Ranks
type Rank rune

const (
	Ace   Rank = 'A'
	Two   Rank = '2'
	Three Rank = '3'
	Four  Rank = '4'
	Five  Rank = '5'
	Six   Rank = '6'
	Seven Rank = '7'
	Eight Rank = '8'
	Nine  Rank = '9'
	Ten   Rank = 'X'
	Jack  Rank = 'J'
	Queen Rank = 'Q'
	King  Rank = 'K'
)

var Ranks = [13]Rank{
	Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King,
}

// These are initial values of cards. Card.Values() should be used to get its
// possible values.
var RankValues = map[Rank]int{
	Ace:   1,
	Two:   2,
	Three: 3,
	Four:  4,
	Five:  5,
	Six:   6,
	Seven: 7,
	Eight: 8,
	Nine:  9,
	Ten:   10,
	Jack:  10,
	Queen: 10,
	King:  10,
}

// Card
type Card struct {
	rank Rank
	suit Suit
}

func (c *Card) String() string {
	return string(rune(c.rank) + rune(c.suit))
}

func TakeCard() int {
	for {
		card := Card{Ranks[rand.Int()%len(Ranks)], Suits[rand.Int()%len(Suits)]}
		if _, ok := Trash[card.String()]; ok {
			//fmt.Println("Найдена похожая карта", string(card.rank), string(card.suit))
			continue
		} else {
			fmt.Println("Вытянутая карта:", string(card.rank), string(card.suit))
			Trash[card.String()] = card
			return int(RankValues[card.rank])
		}
	}
}
