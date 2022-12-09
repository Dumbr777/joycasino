package logic

import (
	"fmt"
	"game/Game2.0/cards"
)

type PlayerNew struct {
	Name  string
	Score int
}

func ScoreNew(i int) string {
	if i < 21 {
		return "continue"
	} else if i == 21 {
		return "win"
	} else {
		return "lose"
	}
}
func WhoWin(p, d int) string {
	if p > d {
		return "player win"
	} else if p == d {
		return "draw"
	} else {
		return "player lose"
	}
}

func Game() string {
	player := PlayerNew{"Player", 0}
	dealer := PlayerNew{"Dealer", 0}

	for {
		var input string
		fmt.Scanln(&input)
		if input == "Hit" {
			n := cards.TakeCard()
			if n == 1 {
				fmt.Printf("Выберете 1 или 11, Ваш счет: %v\n", player.Score)
				fmt.Scanln(&input)
				if input == "1" {
					player.Score += 1
				} else {
					player.Score += 11
				}
				fmt.Println("Ваш счет:", player.Score)
				switch ScoreNew(player.Score) {
				case "win":
					fmt.Println("Вы победили у Вас 21 очко!")
					return "player win"
				case "lose":
					fmt.Println("Вы проиграли у Вас перебор!")
					return "player lose"
				}
			} else {
				player.Score += n
				fmt.Println("Ваш счет:", player.Score)
				switch ScoreNew(player.Score) {
				case "win":
					fmt.Println("Вы победили у Вас 21 очко!")
					return "player win"
				case "lose":
					fmt.Println("Вы проиграли у Вас перебор!")
					return "player lose"
				}
			}

		} else if input == "Stand" {
			for dealer.Score < 17 {
				n := cards.TakeCard()
				if n == 1 && dealer.Score < 11 {
					dealer.Score += 11
					fmt.Println("Счет дилера:", dealer.Score)
					switch ScoreNew(dealer.Score) {
					case "win":
						fmt.Println("Дилер победил у него 21 очко!")
						return "dealer win"
					case "lose":
						fmt.Println("Вы выиграли у Дилера перебор!")
						return "dealer lose"
					}
					fmt.Println("Счет дилера:", dealer.Score)
				} else if n == 1 && dealer.Score >= 11 {
					dealer.Score += 1
					fmt.Println("Счет дилера:", dealer.Score)
					switch ScoreNew(dealer.Score) {
					case "win":
						fmt.Println("Дилер победил у него 21 очко!")
						return "dealer win"
					case "lose":
						fmt.Println("Вы выиграли у Дилера перебор!")
						return "dealer lose"
					}
				} else {
					dealer.Score += n
					fmt.Println("Счет дилера:", dealer.Score)
					switch ScoreNew(dealer.Score) {
					case "win":
						fmt.Println("Дилер победил у него 21 очко!")
						return "dealer win"
					case "lose":
						fmt.Println("Вы выиграли у Дилера перебор!")
						return "dealer lose"
					}
				}
			}

			return WhoWin(player.Score, dealer.Score)
		} else {
			fmt.Println("Вы ввели неправильную команду, введите заного")
			continue
		}
	}
}
