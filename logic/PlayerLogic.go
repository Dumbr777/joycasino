package logic

import (
	"fmt"
	"game/Game2.0/cards"
	"os"
)

var Player int

func PlayerGame() int {
	for {
		var input string
		fmt.Scanln(&input)
		if input == "Hit" {
			n := cards.TakeCard()

			if n == 1 {
				fmt.Printf("Выберете 1 или 11, Ваш счет: %v\n", Player)
				fmt.Scanln(&input)
				if input == "1" {
					Player += 1
				} else {
					Player += 11
				}
				fmt.Println("Ваш счет:", Player)
				switch Score(Player) {
				case 1:
					fmt.Println("Вы победили у Вас 21 очко!")
					os.Exit(1)
				case 2:
					fmt.Println("Вы проиграли у Вас перебор!")
					os.Exit(1)
				}
			} else {
				Player += n
				fmt.Println("Ваш счет:", Player)
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
