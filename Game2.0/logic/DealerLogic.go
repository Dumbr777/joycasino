package logic

import (
	"fmt"
	"game/GameRefactor/cards"
	"math/rand"
	"os"
)

var Dealer int

func DealerGame() int {
	for {
		if Dealer < 17 {
			n := rand.Int() % len(cards.Cards)
			if cards.Deck[cards.Cards[n]] > 0 {
				Dealer += cards.CardValue("Дилер", cards.Cards[n], Dealer)
				cards.Deck[cards.Cards[n]] -= 1
				fmt.Println("Карта", cards.Cards[n])
				fmt.Println("Счет Дилера", Dealer)
				switch Score(Dealer) {
				case 1:
					fmt.Println("Дилер победил у него 21 очко!")
					os.Exit(1)
				case 2:
					fmt.Println("Вы выиграли у Дилера перебор!")
					os.Exit(1)
				}
			} else {
				delete(cards.Deck, cards.Cards[n])
				n := rand.Int() % len(cards.Cards)
				Dealer += cards.CardValue("Дилер", cards.Cards[n], Dealer)
				fmt.Println("Карта", cards.Cards[n])
				fmt.Println("Счет Дилера", Dealer)
				switch Score(Dealer) {
				case 1:
					fmt.Println("Дилер победил у него 21 очко!")
					os.Exit(1)
				case 2:
					fmt.Println("Вы выиграли у Дилера перебор!")
					os.Exit(1)
				}
			}
		} else {
			return Dealer
		}
	}
}
