package logic

import (
	"fmt"
	"game/GameRefactor/cards"
	"math/rand"
	"os"
)

var Player int

func PlayerGame() int {
	for {
		var input string
		fmt.Scanln(&input)
		if input == "Hit" {
			n := rand.Int() % len(cards.Cards)
			if cards.Deck[cards.Cards[n]] > 0 {
				Player += cards.CardValue("Игрок", cards.Cards[n], Player)
				cards.Deck[cards.Cards[n]] -= 1
				fmt.Println("Карта", cards.Cards[n])
				fmt.Println("Счет игрока", Player)
				switch Score(Player) {
				case 1:
					fmt.Println("Вы победили у Вас 21 очко!")
					os.Exit(1)
				case 2:
					fmt.Println("Вы проиграли у Вас перебор!")
					os.Exit(1)
				}
			} else {
				delete(cards.Deck, cards.Cards[n])
				n := rand.Int() % len(cards.Cards)
				Player += cards.CardValue("Игрок", cards.Cards[n], Player)
				fmt.Println("Карта", cards.Cards[n])
				fmt.Println("Счет игрока", Player)
				switch Score(Player) {
				case 1:
					fmt.Println("Вы победили у Вас 21 очко!")
					os.Exit(1)
				case 2:
					fmt.Println("Вы проиграли у Вас перебор!")
					os.Exit(1)
				}
			}

		} else if input == "Stand" {
			return Player
		} else {
			fmt.Println("Вы ввели неправильную команду, введите заного")
			continue
		}
	}
}
