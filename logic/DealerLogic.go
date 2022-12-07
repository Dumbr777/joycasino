package logic

import (
	"fmt"
	"game/Game2.0/cards"
	"os"
)

var Dealer int

func DealerGame() int {
	for {
		if Dealer < 17 {
			n := cards.TakeCard()
			if n == 1 && Dealer < 11 {
				Dealer += 11
				fmt.Println("Счет дилера:", Dealer)
				switch Score(Dealer) {
				case 1:
					fmt.Println("Дилер победил у него 21 очко!")
					os.Exit(1)
				case 2:
					fmt.Println("Вы выиграли у Дилера перебор!")
					os.Exit(1)
				}
				fmt.Println("Счет дилера:", Dealer)
			} else if n == 1 && Dealer >= 11 {
				Dealer += 1
				fmt.Println("Счет дилера:", Dealer)
				switch Score(Dealer) {
				case 1:
					fmt.Println("Дилер победил у него 21 очко!")
					os.Exit(1)
				case 2:
					fmt.Println("Вы выиграли у Дилера перебор!")
					os.Exit(1)
				}
			} else {
				Dealer += n
				fmt.Println("Счет дилера:", Dealer)
			}

		} else {
			return Dealer
		}
	}
}
