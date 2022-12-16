package blackjack

import (
	"fmt"

	"Game2.0.0/card"
	"Game2.0.0/deck"
)

type Blackjack struct {
	deck        *deck.Deck
	playerHand  []card.Card
	playerscore int
	dealerHand  []card.Card
	dealerscore int
}

const (
	Win = iota + 1
	Lose
	Draw
	gameContinue
)

func (b *Blackjack) gameScore() int {
	if b.playerscore == 21 || b.dealerscore > 21 {
		return Win
	} else if b.playerscore > 21 || b.dealerscore == 21 {
		return Lose
	} else if b.playerscore == b.dealerscore {
		return Draw
	}

	return gameContinue
}

func New(d *deck.Deck) *Blackjack {
	return &Blackjack{deck: d}
}

func (b *Blackjack) Play() int {

	for {
		var input string
		fmt.Scanln(&input)
		if input == "Hit" {
			c := b.deck.TakeCard()
			if c.Rank() == card.Ace {
				fmt.Printf("Вам выпал туз, выберете 1 или 11, Ваш счет: %v\n", b.playerscore)
				fmt.Scanln(&input)
				if input == "1" {
					b.playerscore += 1
				} else {
					b.playerscore += 11
				}
				fmt.Println("Ваш счет:", b.playerscore)
				switch b.gameScore() {
				case Win:
					fmt.Println("Вы победили у Вас 21 очко!")
					return Win
				case Lose:
					fmt.Println("Вы проиграли у Вас перебор!")
					return Lose
				}
			} else {
				b.playerscore += c.Par()
				fmt.Println("Ваш счет:", b.playerscore)
				switch b.gameScore() {
				case Win:
					fmt.Println("Вы победили у Вас 21 очко!")
					return Win
				case Lose:
					fmt.Println("Вы проиграли у Вас перебор!")
					return Lose
				}
			}

		} else if input == "Stand" {
			for b.dealerscore < 17 {
				c := b.deck.TakeCard()
				if c.Rank() == card.Ace && b.dealerscore < 11 {
					b.dealerscore += 11
					fmt.Println("Счет дилера:", b.dealerscore)
					switch b.gameScore() {
					case Lose:
						fmt.Println("Дилер победил у него 21 очко!")
						return Lose
					case Win:
						fmt.Println("Вы выиграли у Дилера перебор!")
						return Win
					}
					fmt.Println("Счет дилера:", b.dealerscore)
				} else if c.Rank() == card.Ace && b.dealerscore >= 11 {
					b.dealerscore += 1
					fmt.Println("Счет дилера:", b.dealerscore)
					switch b.gameScore() {
					case Lose:
						fmt.Println("Дилер победил у него 21 очко!")
						return Lose
					case Win:
						fmt.Println("Вы выиграли у Дилера перебор!")
						return Win
					}
				} else {
					b.dealerscore += c.Par()
					fmt.Println("Счет дилера:", b.dealerscore)
					switch b.gameScore() {
					case Lose:
						fmt.Println("Дилер победил у него 21 очко!")
						return Lose
					case Win:
						fmt.Println("Вы выиграли у Дилера перебор!")
						return Win
					}
				}
			}

			return b.gameScore()
		} else {
			fmt.Println("Вы ввели неправильную команду, введите заного")
			continue
		}
	}
}
